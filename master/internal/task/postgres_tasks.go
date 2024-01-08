package task

import (
	"context"

	"github.com/o1egl/paseto"
	"github.com/pkg/errors"

	"github.com/determined-ai/determined/master/internal/db"
	"github.com/determined-ai/determined/master/pkg/model"
)

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
