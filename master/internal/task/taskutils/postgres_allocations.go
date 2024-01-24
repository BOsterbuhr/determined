package taskutils

import (
	"context"
	"fmt"
	"strings"

	"github.com/determined-ai/determined/master/internal/db"
	"github.com/determined-ai/determined/master/pkg/model"
	"github.com/pkg/errors"
)

// AddAllocation upserts the existence of an allocation. Allocation IDs may conflict in the event
// the master restarts and the trial run ID increment is not persisted, but it is the same
// allocation so this is OK.
func AddAllocation(a *model.Allocation) error {
	_, err := db.Bun().NewInsert().Model(a).
		On("CONFLICT (allocation_id) DO UPDATE").
		Set("task_id=EXCLUDED.task_id, slots=EXCLUDED.slots, resource_pool=EXCLUDED.resource_pool").
		Set("start_time=EXCLUDED.start_time, state=EXCLUDED.state, ports=EXCLUDED.ports").
		Exec(context.Background()) // TODO CAROLINA, pass in ctx
	return err
}

// AddAllocationExitStatus adds the allocation exit status to the allocations table.
func AddAllocationExitStatus(ctx context.Context, a *model.Allocation) error {
	_, err := db.Bun().NewUpdate().
		Model(a).
		Column("exit_reason", "exit_error", "status_code").
		Where("allocation_id = ?", a.AllocationID).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("adding allocation exit status to db: %w", err)
	}
	return nil
}

// CompleteAllocation persists the end of an allocation lifetime.
func CompleteAllocation(a *model.Allocation) error {
	if a.StartTime == nil {
		a.StartTime = a.EndTime
	}

	_, err := db.Bun().NewUpdate().Table("allocations").
		Set("start_time = ?, end_time = ?", a.StartTime, a.EndTime).
		Where("allocation_id = ?", a.AllocationID).Exec(context.Background()) // TODO CAROLINA, pass in the ctx

	return err
}

// CompleteAllocationTelemetry returns the analytics of an allocation for the telemetry.
func CompleteAllocationTelemetry(aID model.AllocationID) ([]byte, error) {
	var bytes []byte
	err := db.Bun().NewRaw(`
	SELECT json_build_object(
		'allocation_id', a.allocation_id,
		'job_id', t.job_id,
		'task_type', t.task_type,
		'duration_sec', COALESCE(EXTRACT(EPOCH FROM (a.end_time - a.start_time)), 0)
	)
	FROM allocations as a JOIN tasks as t
	ON a.task_id = t.task_id
	WHERE a.allocation_id = ?;
	`, aID).Scan(context.Background(), &bytes) // TODO CAROLINA, pass in the context
	return bytes, err
}

// AddAllocationAcceleratorData stores acceleration data for an allocation.
func AddAllocationAcceleratorData(ctx context.Context, accData model.AcceleratorData,
) error {
	_, err := db.Bun().NewInsert().Model(&accData).Exec(ctx)
	if err != nil {
		return fmt.Errorf("adding allocation acceleration data: %w", err)
	}
	return nil
}

// GetAllocation stores acceleration data for an allocation.
func GetAllocation(ctx context.Context, allocationID string,
) (*model.Allocation, error) {
	var allocation model.Allocation
	err := db.Bun().NewRaw(`
SELECT allocation_id, task_id, state, slots, is_ready, start_time, 
end_time, exit_reason, exit_error, status_code
FROM allocations
WHERE allocation_id = ?
	`, allocationID).Scan(ctx, &allocation)
	if err != nil {
		return nil, fmt.Errorf("querying allocation %s: %w", allocationID, err)
	}

	return &allocation, nil
}

// CloseOpenAllocations finds all allocations that were open when the master crashed
// and adds an end time.
func CloseOpenAllocations(exclude []model.AllocationID) error {
	if _, err := db.Bun().NewRaw(`
	UPDATE allocations
	SET start_time = cluster_heartbeat FROM cluster_id
	WHERE start_time is NULL`).Exec(context.Background()); err != nil {
		return errors.Wrap(err,
			"setting start time to cluster heartbeat when it's assigned to zero value")
	}

	excludedFilter := ""
	if len(exclude) > 0 {
		excludeStr := make([]string, 0, len(exclude))
		for _, v := range exclude {
			excludeStr = append(excludeStr, v.String())
		}

		excludedFilter = strings.Join(excludeStr, ",")
	}

	if _, err := db.Bun().NewRaw(`
	UPDATE allocations
	SET end_time = greatest(cluster_heartbeat, start_time), state = 'TERMINATED'
	FROM cluster_id
	WHERE end_time IS NULL AND
	(? = '' OR allocation_id NOT IN (
		SELECT unnest(string_to_array(?, ','))))`, excludedFilter, excludedFilter).Exec(context.Background()); err != nil {
		return errors.Wrap(err, "closing old allocations")
	}
	return nil
}
