package taskutils

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"

	"github.com/determined-ai/determined/master/internal/db"
	"github.com/determined-ai/determined/master/pkg/model"
)

// AddTask UPSERT's the existence of a task.
func AddTask(ctx context.Context, t *model.Task) error {
	// Since AddTaskTx is a single query, RunInTx is an overkill.
	return AddTaskTx(ctx, db.Bun(), t)
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
	return db.MatchSentinelError(err)
}

// TaskByID returns a task by its ID.
func TaskByID(ctx context.Context, tID model.TaskID) (*model.Task, error) {
	var t model.Task
	if err := db.Bun().NewSelect().Model(&t).Where("task_id = ?", tID).Scan(ctx, &t); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = db.ErrNotFound
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

	if _, err := db.Bun().NewInsert().Model(&model.TaskContextDirectory{
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
	if err := db.Bun().NewSelect().Model(res).Where("task_id = ?", tID).Scan(ctx, res); err != nil {
		return nil, fmt.Errorf("querying task ID %s context directory files: %w", tID, err)
	}

	return res.ContextDirectory, nil
}

// TaskCompleted checks if the end time exists for a task, if so, the task has completed.
func TaskCompleted(ctx context.Context, tID model.TaskID) (bool, error) {
	return db.Bun().NewSelect().Table("tasks").
		Where("task_id = ?", tID).Where("end_time IS NOT NULL").Exists(ctx)
}

// CompleteTask persists the completion of a task.
func CompleteTask(tID model.TaskID, endTime time.Time) error {
	if err := db.Bun().NewUpdate().Table("tasks").Set("end_time = ?", endTime).
		Where("task_id = ?", tID).Scan(context.Background()); err != nil {
		return errors.Wrap(err, "completing task")
	} // TODO CAROLINA -- pass in the ctx

	return nil
}

// AllocationByID retrieves an allocation by its ID.
func AllocationByID(ctx context.Context, aID model.AllocationID) (*model.Allocation, error) {
	var a model.Allocation
	if err := db.Bun().NewSelect().Model(&a).Where("allocation_id = ?", aID).
		Scan(ctx); err != nil {
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

	if _, err := db.Bun().NewInsert().Model(&taskSession).
		Table("allocation_sessions").ColumnExpr("allocation_id, owner_id").
		Returning("id").Exec(ctx); err != nil {
		return "", err
	}

	v2 := paseto.NewV2()
	token, err := v2.Sign(db.GetTokenKeys().PrivateKey, taskSession, nil)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate task authentication token")
	}
	return token, nil
}

// DeleteAllocationSession deletes the task session with the given AllocationID.
func DeleteAllocationSession(ctx context.Context, allocationID model.AllocationID) error {
	_, err := db.Bun().NewDelete().Table("allocation_sessions").Where("allocation_id = ?", allocationID).Exec(ctx)
	return err
}

// UpdateAllocationState stores the latest task state and readiness.
func UpdateAllocationState(ctx context.Context, a model.Allocation) error {
	_, err := db.Bun().NewUpdate().Table("allocations").Set("state = ?", a.State).Set("is_ready = ?", a.IsReady).
		Where("allocation_id = ?", a.AllocationID).Exec(ctx)
	return err
}

// UpdateAllocationPorts stores the latest task state and readiness.
func UpdateAllocationPorts(ctx context.Context, a model.Allocation) error {
	_, err := db.Bun().NewUpdate().Table("allocations").
		Set("ports = ?", a.Ports).
		Where("allocation_id = ?", a.AllocationID).
		Exec(context.TODO())
	return err
}

// UpdateAllocationStartTime stores the latest start time.
func UpdateAllocationStartTime(ctx context.Context, a model.Allocation) error {
	_, err := db.Bun().NewUpdate().Table("allocations").Set("start_time = ?", a.StartTime).
		Where("allocation_id = ?", a.AllocationID).Exec(ctx)
	return err
}

// UpdateAllocationProxyAddress stores the proxy address.
func UpdateAllocationProxyAddress(ctx context.Context, a model.Allocation) error {
	_, err := db.Bun().NewUpdate().Table("allocations").
		Set("proxy_address = ?", a.ProxyAddress).
		Where("allocation_id = ?", a.AllocationID).Exec(ctx)
	return err
}
