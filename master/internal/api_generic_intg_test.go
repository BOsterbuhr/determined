//go:build integration
// +build integration

package internal

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/determined-ai/determined/master/pkg/model"
	"github.com/determined-ai/determined/proto/pkg/apiv1"
)

func TestPropagateTaskState(t *testing.T) {
	api, _, ctx := setupAPITest(t, nil)

	parentID := model.NewTaskID()
	child1ID := model.NewTaskID()
	child2ID := model.NewTaskID()

	parentModel := &model.Task{TaskType: model.TaskTypeGeneric, TaskID: parentID}
	child1Model := &model.Task{TaskType: model.TaskTypeGeneric, TaskID: child1ID, ParentID: &parentID}
	child2Model := &model.Task{TaskType: model.TaskTypeGeneric, TaskID: child2ID, ParentID: &parentID}
	require.NoError(t, api.m.db.AddTask(parentModel))
	require.NoError(t, api.m.db.AddTask(child1Model))
	require.NoError(t, api.m.db.AddTask(child2Model))

	require.NoError(t, api.PropagateTaskState(ctx, parentID, model.TaskStateStoppingCanceled))

	parent, err := api.GetTask(ctx, &apiv1.GetTaskRequest{TaskId: parentID.String()})
	require.NoError(t, err)
	child1, err := api.GetTask(ctx, &apiv1.GetTaskRequest{TaskId: child1ID.String()})
	require.NoError(t, err)
	child2, err := api.GetTask(ctx, &apiv1.GetTaskRequest{TaskId: child2ID.String()})
	require.NoError(t, err)
	require.Equal(t, string(model.TaskStateStoppingCanceled), *parent.Task.TaskState)
	require.Equal(t, string(model.TaskStateStoppingCanceled), *child1.Task.TaskState)
	require.Equal(t, string(model.TaskStateStoppingCanceled), *child2.Task.TaskState)
}

func TestFindRoot(t *testing.T) {
	api, _, ctx := setupAPITest(t, nil)

	parentID := model.NewTaskID()
	child1ID := model.NewTaskID()
	child2ID := model.NewTaskID()

	parent := &model.Task{TaskType: model.TaskTypeGeneric, TaskID: parentID}
	child1 := &model.Task{TaskType: model.TaskTypeGeneric, TaskID: child1ID, ParentID: &parentID}
	child2 := &model.Task{TaskType: model.TaskTypeGeneric, TaskID: child2ID, ParentID: &parentID}
	require.NoError(t, api.m.db.AddTask(parent))
	require.NoError(t, api.m.db.AddTask(child1))
	require.NoError(t, api.m.db.AddTask(child2))

	taskID, err := api.FindRoot(ctx, child1ID)
	require.NoError(t, err)
	require.Equal(t, parentID, taskID)
}

func TestGetTaskChildren(t *testing.T) {
	api, _, ctx := setupAPITest(t, nil)

	parentID := model.NewTaskID()
	child1ID := model.NewTaskID()
	child2ID := model.NewTaskID()

	parent := &model.Task{TaskType: model.TaskTypeGeneric, TaskID: parentID}
	child1 := &model.Task{TaskType: model.TaskTypeGeneric, TaskID: child1ID, ParentID: &parentID}
	child2 := &model.Task{TaskType: model.TaskTypeGeneric, TaskID: child2ID, ParentID: &parentID}
	require.NoError(t, api.m.db.AddTask(parent))
	require.NoError(t, api.m.db.AddTask(child1))
	require.NoError(t, api.m.db.AddTask(child2))

	taskSet := map[model.TaskID]bool{parentID: true, child1ID: true, child2ID: true}

	tasks, err := api.GetTaskChildren(ctx, parentID)
	require.NoError(t, err)
	for _, e := range tasks {
		_, ok := taskSet[e.TaskID]
		require.Equal(t, true, ok)
	}
}
