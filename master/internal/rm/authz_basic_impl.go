package rm

import (
	"context"

	"github.com/determined-ai/determined/master/pkg/generatedproto/resourcepoolv1"
	"github.com/determined-ai/determined/master/pkg/model"
)

// ResourceManagerAuthZBasic is classic OSS Determined authentication for resource managers.
type ResourceManagerAuthZBasic struct{}

// FilterResourcePools always returns provided list and a nil error.
func (a *ResourceManagerAuthZBasic) FilterResourcePools(
	ctx context.Context, curUser model.User, resourcePools []*resourcepoolv1.ResourcePool,
	_ []int32,
) ([]*resourcepoolv1.ResourcePool, error) {
	return resourcePools, nil
}

func init() {
	AuthZProvider.Register("basic", &ResourceManagerAuthZBasic{})
}
