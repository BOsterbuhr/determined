// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/determined-ai/determined/master/pkg/model"

	tensorboardv1 "github.com/determined-ai/determined/proto/pkg/tensorboardv1"
)

// NSCAuthZ is an autogenerated mock type for the NSCAuthZ type
type NSCAuthZ struct {
	mock.Mock
}

// AccessibleScopes provides a mock function with given fields: ctx, curUser, requestedScope
func (_m *NSCAuthZ) AccessibleScopes(ctx context.Context, curUser model.User, requestedScope model.AccessScopeID) (map[model.AccessScopeID]bool, error) {
	ret := _m.Called(ctx, curUser, requestedScope)

	var r0 map[model.AccessScopeID]bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, model.AccessScopeID) (map[model.AccessScopeID]bool, error)); ok {
		return rf(ctx, curUser, requestedScope)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.User, model.AccessScopeID) map[model.AccessScopeID]bool); ok {
		r0 = rf(ctx, curUser, requestedScope)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[model.AccessScopeID]bool)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.User, model.AccessScopeID) error); ok {
		r1 = rf(ctx, curUser, requestedScope)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CanCreateNSC provides a mock function with given fields: ctx, curUser, workspaceID
func (_m *NSCAuthZ) CanCreateNSC(ctx context.Context, curUser model.User, workspaceID model.AccessScopeID) error {
	ret := _m.Called(ctx, curUser, workspaceID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, model.AccessScopeID) error); ok {
		r0 = rf(ctx, curUser, workspaceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanGetActiveTasksCount provides a mock function with given fields: ctx, curUser
func (_m *NSCAuthZ) CanGetActiveTasksCount(ctx context.Context, curUser model.User) error {
	ret := _m.Called(ctx, curUser)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User) error); ok {
		r0 = rf(ctx, curUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanGetNSC provides a mock function with given fields: ctx, curUser, workspaceID
func (_m *NSCAuthZ) CanGetNSC(ctx context.Context, curUser model.User, workspaceID model.AccessScopeID) (bool, error) {
	ret := _m.Called(ctx, curUser, workspaceID)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, model.AccessScopeID) (bool, error)); ok {
		return rf(ctx, curUser, workspaceID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.User, model.AccessScopeID) bool); ok {
		r0 = rf(ctx, curUser, workspaceID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.User, model.AccessScopeID) error); ok {
		r1 = rf(ctx, curUser, workspaceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CanGetTensorboard provides a mock function with given fields: ctx, curUser, workspaceID, experimentIDs, trialIDs
func (_m *NSCAuthZ) CanGetTensorboard(ctx context.Context, curUser model.User, workspaceID model.AccessScopeID, experimentIDs []int32, trialIDs []int32) (bool, error) {
	ret := _m.Called(ctx, curUser, workspaceID, experimentIDs, trialIDs)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, model.AccessScopeID, []int32, []int32) (bool, error)); ok {
		return rf(ctx, curUser, workspaceID, experimentIDs, trialIDs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.User, model.AccessScopeID, []int32, []int32) bool); ok {
		r0 = rf(ctx, curUser, workspaceID, experimentIDs, trialIDs)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.User, model.AccessScopeID, []int32, []int32) error); ok {
		r1 = rf(ctx, curUser, workspaceID, experimentIDs, trialIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CanSetNSCsPriority provides a mock function with given fields: ctx, curUser, workspaceID, priority
func (_m *NSCAuthZ) CanSetNSCsPriority(ctx context.Context, curUser model.User, workspaceID model.AccessScopeID, priority int) error {
	ret := _m.Called(ctx, curUser, workspaceID, priority)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, model.AccessScopeID, int) error); ok {
		r0 = rf(ctx, curUser, workspaceID, priority)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanTerminateNSC provides a mock function with given fields: ctx, curUser, workspaceID
func (_m *NSCAuthZ) CanTerminateNSC(ctx context.Context, curUser model.User, workspaceID model.AccessScopeID) error {
	ret := _m.Called(ctx, curUser, workspaceID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, model.AccessScopeID) error); ok {
		r0 = rf(ctx, curUser, workspaceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanTerminateTensorboard provides a mock function with given fields: ctx, curUser, workspaceID
func (_m *NSCAuthZ) CanTerminateTensorboard(ctx context.Context, curUser model.User, workspaceID model.AccessScopeID) error {
	ret := _m.Called(ctx, curUser, workspaceID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, model.AccessScopeID) error); ok {
		r0 = rf(ctx, curUser, workspaceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilterTensorboards provides a mock function with given fields: ctx, curUser, requestedScope, tensorboards
func (_m *NSCAuthZ) FilterTensorboards(ctx context.Context, curUser model.User, requestedScope model.AccessScopeID, tensorboards []*tensorboardv1.Tensorboard) ([]*tensorboardv1.Tensorboard, error) {
	ret := _m.Called(ctx, curUser, requestedScope, tensorboards)

	var r0 []*tensorboardv1.Tensorboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, model.AccessScopeID, []*tensorboardv1.Tensorboard) ([]*tensorboardv1.Tensorboard, error)); ok {
		return rf(ctx, curUser, requestedScope, tensorboards)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.User, model.AccessScopeID, []*tensorboardv1.Tensorboard) []*tensorboardv1.Tensorboard); ok {
		r0 = rf(ctx, curUser, requestedScope, tensorboards)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*tensorboardv1.Tensorboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.User, model.AccessScopeID, []*tensorboardv1.Tensorboard) error); ok {
		r1 = rf(ctx, curUser, requestedScope, tensorboards)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewNSCAuthZ interface {
	mock.TestingT
	Cleanup(func())
}

// NewNSCAuthZ creates a new instance of NSCAuthZ. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNSCAuthZ(t mockConstructorTestingTNewNSCAuthZ) *NSCAuthZ {
	mock := &NSCAuthZ{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
