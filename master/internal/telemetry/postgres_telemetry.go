package telemetry

import (
	"context"

	"github.com/determined-ai/determined/master/internal/db"
	"github.com/determined-ai/determined/master/pkg/model"
)

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
