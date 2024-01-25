package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/o1egl/paseto"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"

	"github.com/determined-ai/determined/master/pkg/model"
)

// initAllocationSessions purges sessions of all closed allocations.
func initAllocationSessions(ctx context.Context) error {
	subq := Bun().NewSelect().Table("allocations").
		Column("allocation_id").Where("start_time IS NOT NULL AND end_time IS NOT NULL")
	_, err := Bun().NewDelete().Table("allocation_sessions").Where("allocation_id in (?)", subq).Exec(ctx)
	return err
}

// AddTask UPSERT's the existence of a task.
func AddTask(ctx context.Context, t *model.Task) error {
	return AddTaskTx(ctx, Bun(), t)
}

// AddTaskTx UPSERT's the existence of a task in a tx.
func AddTaskTx(ctx context.Context, idb bun.IDB, t *model.Task) error {
	_, err := idb.NewInsert().Model(t).
		Column("task_id", "task_type", "start_time", "job_id", "log_version").
		On("CONFLICT (task_id) DO UPDATE").
		Set("task_type=EXCLUDED.task_type").
		Set("start_time=EXCLUDED.start_time").
		Set("job_id=EXCLUDED.job_id").
		Set("log_version=EXCLUDED.log_version").
		Exec(ctx)
	return MatchSentinelError(err)
}

// TaskByID returns a task by its ID.
func TaskByID(ctx context.Context, tID model.TaskID) (*model.Task, error) {
	var t model.Task
	if err := Bun().NewSelect().Model(&t).Where("task_id = ?", tID).Scan(ctx, &t); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrNotFound
		}
		return nil, fmt.Errorf("querying task ID %s: %w", tID, err)
	}

	return &t, nil
}

// AddNonExperimentTasksContextDirectory adds a context directory for a non experiment task.
func AddNonExperimentTasksContextDirectory(ctx context.Context, tID model.TaskID, bytes []byte) error {
	if bytes == nil {
		bytes = []byte{}
	}

	if _, err := Bun().NewInsert().Model(&model.TaskContextDirectory{
		TaskID:           tID,
		ContextDirectory: bytes,
	}).Exec(ctx); err != nil {
		return fmt.Errorf("persisting context directory files for task %s: %w", tID, err)
	}

	return nil
}

// NonExperimentTasksContextDirectory returns a non experiment's context directory.
func NonExperimentTasksContextDirectory(ctx context.Context, tID model.TaskID) ([]byte, error) {
	res := &model.TaskContextDirectory{}
	if err := Bun().NewSelect().Model(res).Where("task_id = ?", tID).Scan(ctx, res); err != nil {
		return nil, fmt.Errorf("querying task ID %s context directory files: %w", tID, err)
	}

	return res.ContextDirectory, nil
}

// TaskCompleted checks if the end time exists for a task, if so, the task has completed.
func TaskCompleted(ctx context.Context, tID model.TaskID) (bool, error) {
	return Bun().NewSelect().Table("tasks").
		Where("task_id = ?", tID).Where("end_time IS NOT NULL").Exists(ctx)
}

// CompleteTask persists the completion of a task.
func CompleteTask(ctx context.Context, tID model.TaskID, endTime time.Time) error {
	if _, err := Bun().NewUpdate().Table("tasks").Set("end_time = ?", endTime).
		Where("task_id = ?", tID).Exec(ctx); err != nil {
		return fmt.Errorf("completing task: %w", err)
	}
	return nil
}

// AddAllocation upserts the existence of an allocation. Allocation IDs may conflict in the event
// the master restarts and the trial run ID increment is not persisted, but it is the same
// allocation so this is OK.
func AddAllocation(ctx context.Context, a *model.Allocation) error {
	_, err := Bun().NewInsert().Table("allocations").Model(a).On("CONFLICT (allocation_id) DO UPDATE").
		Set("task_id=EXCLUDED.task_id, slots=EXCLUDED.slots").
		Set("resource_pool=EXCLUDED.resource_pool,start_time=EXCLUDED.start_time").
		Set("state=EXCLUDED.state, ports=EXCLUDED.ports").Exec(ctx)
	return err
}

// AddAllocationExitStatus adds the allocation exit status to the allocations table.
func AddAllocationExitStatus(ctx context.Context, a *model.Allocation) error {
	if _, err := Bun().NewUpdate().Model(a).
		Column("exit_reason", "exit_error", "status_code").
		Where("allocation_id = ?", a.AllocationID).Exec(ctx); err != nil {
		return fmt.Errorf("adding allocation exit status to db: %w", err)
	}
	return nil
}

// CompleteAllocation persists the end of an allocation lifetime.
func CompleteAllocation(ctx context.Context, a *model.Allocation) error {
	if a.StartTime == nil {
		a.StartTime = a.EndTime
	}

	_, err := Bun().NewUpdate().Table("allocations").Model(a).
		Set("start_time = ?, end_time = ?", a.StartTime, a.EndTime).
		Where("allocation_id = ?", a.AllocationID).Exec(ctx)

	return err
}

// CompleteAllocationTelemetry returns the analytics of an allocation for the telemetry.
func CompleteAllocationTelemetry(ctx context.Context, aID model.AllocationID) ([]byte, error) {
	var res []byte
	err := Bun().NewRaw(`
	SELECT json_build_object(
		'allocation_id', a.allocation_id,
		'job_id', t.job_id,
		'task_type', t.task_type,
		'duration_sec', COALESCE(EXTRACT(EPOCH FROM (a.end_time - a.start_time)), 0)
	)
	FROM allocations as a JOIN tasks as t
	ON a.task_id = t.task_id
	WHERE a.allocation_id = ?`, aID).Scan(ctx, &res)
	return res, err
}

// AllocationByID retrieves an allocation by its ID.
func AllocationByID(ctx context.Context, aID model.AllocationID) (*model.Allocation, error) {
	var a model.Allocation
	if err := Bun().NewSelect().Table("allocations").
		Where("allocation_id = ?", aID).Scan(ctx, &a); err != nil {
		return nil, err
	}
	return &a, nil
}

// StartAllocationSession creates a row in the allocation_sessions table.
func StartAllocationSession(
	ctx context.Context,
	allocationID model.AllocationID,
	owner *model.User,
) (string, error) {
	if owner == nil {
		return "", errors.New("owner cannot be nil for allocation session")
	}

	taskSession := &model.AllocationSession{
		AllocationID: allocationID,
		OwnerID:      &owner.ID,
	}

	if _, err := Bun().NewInsert().Table("allocation_sessions").
		Model(taskSession).Returning("id").Exec(ctx, &taskSession.ID); err != nil {
		return "", err
	}

	v2 := paseto.NewV2()
	token, err := v2.Sign(GetTokenKeys().PrivateKey, taskSession, nil)
	if err != nil {
		return "", fmt.Errorf("failed to generate task authentication token: %w", err)
	}
	return token, nil
}

// DeleteAllocationSession deletes the task session with the given AllocationID.
func DeleteAllocationSession(ctx context.Context, allocationID model.AllocationID) error {
	_, err := Bun().NewDelete().Table("allocation_sessions").Where("allocation_id = ?", allocationID).Exec(ctx)
	return err
}

// UpdateAllocationState stores the latest task state and readiness.
func UpdateAllocationState(ctx context.Context, a model.Allocation) error {
	_, err := Bun().NewUpdate().Table("allocations").
		Set("state = ?, is_ready = ?", a.State, a.IsReady).
		Where("allocation_id = ?", a.AllocationID).Exec(ctx)

	return err
}

// UpdateAllocationPorts stores the latest task state and readiness.
func UpdateAllocationPorts(ctx context.Context, a model.Allocation) error {
	_, err := Bun().NewUpdate().Table("allocations").
		Set("ports = ?", a.Ports).
		Where("allocation_id = ?", a.AllocationID).
		Exec(ctx)
	return err
}

// UpdateAllocationStartTime stores the latest start time.
func UpdateAllocationStartTime(ctx context.Context, a model.Allocation) error {
	_, err := Bun().NewUpdate().Table("allocations").
		Set("start_time = ?", a.StartTime).Where("allocation_id = ?", a.AllocationID).Exec(ctx)
	return err
}

// UpdateAllocationProxyAddress stores the proxy address.
func UpdateAllocationProxyAddress(ctx context.Context, a model.Allocation) error {
	_, err := Bun().NewUpdate().Table("allocations").Set("proxy_address = ?", a.ProxyAddress).
		Where("allocation_id = ?", a.AllocationID).Exec(ctx)
	return err
}

// CloseOpenAllocations finds all allocations that were open when the master crashed
// and adds an end time.
func CloseOpenAllocations(ctx context.Context, exclude []model.AllocationID) error {
	if _, err := Bun().NewUpdate().Table("allocations").Set("start_time = cluster_heartbeat FROM cluster_id").
		Where("start_time is NULL").Exec(ctx); err != nil {
		return fmt.Errorf("setting start time to cluster heartbeat when it's assigned to zero value: %w", err)
	}

	excludedFilter := ""
	if len(exclude) > 0 {
		excludeStr := make([]string, 0, len(exclude))
		for _, v := range exclude {
			excludeStr = append(excludeStr, v.String())
		}

		excludedFilter = strings.Join(excludeStr, ",")
	}

	if _, err := Bun().NewUpdate().Table("allocations, cluster_id").
		Set("end_time = greatest(cluster_heartbeat, start_time), state = 'TERMINATED'").
		Where("end_time IS NULL AND (? = '' OR allocation_id NOT IN (SELECT unnest(string_to_array(?, ','))))",
			excludedFilter).Exec(ctx); err != nil {
		return fmt.Errorf("closing old allocations: %w", err)
	}
	return nil
}

// RecordTaskStats record stats for tasks.
func RecordTaskStats(ctx context.Context, stats *model.TaskStats) error {
	return RecordTaskStatsBun(ctx, stats)
}

// RecordTaskStatsBun record stats for tasks with bun.
func RecordTaskStatsBun(ctx context.Context, stats *model.TaskStats) error {
	_, err := Bun().NewInsert().Model(stats).Exec(context.TODO())
	return err
}

// RecordTaskEndStats record end stats for tasks.
func RecordTaskEndStats(ctx context.Context, stats *model.TaskStats) error {
	return RecordTaskEndStatsBun(ctx, stats)
}

// RecordTaskEndStatsBun record end stats for tasks with bun.
func RecordTaskEndStatsBun(ctx context.Context, stats *model.TaskStats) error {
	query := Bun().NewUpdate().Model(stats).Column("end_time").
		Where("allocation_id = ?", stats.AllocationID).
		Where("event_type = ?", stats.EventType).
		Where("end_time IS NULL")
	if stats.ContainerID == nil {
		// Just doing Where("container_id = ?", stats.ContainerID) in the null case
		// generates WHERE container_id = NULL which doesn't seem to match on null rows.
		// We don't use this case anywhere currently but this feels like an easy bug to write
		// without this.
		query = query.Where("container_id IS NULL")
	} else {
		query = query.Where("container_id = ?", stats.ContainerID)
	}

	if _, err := query.Exec(ctx); err != nil {
		return fmt.Errorf("recording task end stats %+v: %w", stats, err)
	}

	return nil
}

// EndAllTaskStats called at master starts, in case master previously crashed.
func EndAllTaskStats(ctx context.Context) error {
	if _, err := Bun().NewUpdate().Table("task_stats", "cluster_id, allocations").
		Set("end_time = greatest(cluster_heartbeat, task_stats.start_time)").
		Where("allocations.allocation_id = task_stats.allocation_id").
		Where("allocations_end_time IS NOT NULL AND task_stats.end_time IS NULL").
		Exec(ctx); err != nil {
		return fmt.Errorf("ending all task stats: %w", err)
	}

	return nil
}
