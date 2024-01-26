package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/determined-ai/determined/master/internal/api"
	"github.com/determined-ai/determined/master/pkg/model"
	"github.com/determined-ai/determined/proto/pkg/apiv1"
)

func TestTaskLogsFlow(t *testing.T) {
	db := MustResolveTestPostgres(t)
	t1In := RequireMockTask(t, db, nil)
	t2In := RequireMockTask(t, db, nil)

	// Test AddTaskLogs & TaskLogCounts
	taskLog1 := RequireMockTaskLog(t, db, t1In.TaskID, "1")
	taskLog2 := RequireMockTaskLog(t, db, t1In.TaskID, "2")
	taskLog3 := RequireMockTaskLog(t, db, t2In.TaskID, "3")

	// Try adding only taskLog1, and count only 1 log.
	err := db.AddTaskLogs([]*model.TaskLog{taskLog1})
	require.NoError(t, err)

	count, err := db.TaskLogsCount(t1In.TaskID, []api.Filter{})
	require.NoError(t, err)
	require.Equal(t, count, 1)

	// Try adding the rest of the Task logs, and count 2 for t1In.TaskID, and 1 for t2In.TaskID
	err = db.AddTaskLogs([]*model.TaskLog{taskLog2, taskLog3})
	require.NoError(t, err)

	count, err = db.TaskLogsCount(t1In.TaskID, []api.Filter{})
	require.NoError(t, err)
	require.Equal(t, count, 2)

	count, err = db.TaskLogsCount(t2In.TaskID, []api.Filter{})
	require.NoError(t, err)
	require.Equal(t, count, 1)

	// Test TaskLogsFields.
	resp, err := db.TaskLogsFields(t1In.TaskID)
	require.NoError(t, err)
	require.ElementsMatch(t, resp.AgentIds, []string{"testing-agent-1", "testing-agent-2"})
	require.ElementsMatch(t, resp.ContainerIds, []string{"1", "2"})

	// Test TaskLogs.
	// Get 1 task log matching t1In task ID.
	logs, _, err := db.TaskLogs(t1In.TaskID, 1, []api.Filter{}, apiv1.OrderBy_ORDER_BY_UNSPECIFIED, nil)
	require.NoError(t, err)
	require.Equal(t, 1, len(logs))
	require.Equal(t, logs[0].TaskID, string(t1In.TaskID))
	require.Contains(t, []string{"1", "2"}, *logs[0].ContainerID)

	// Get up to 5 tasks matching t2In task ID -- receive only 2.
	logs, _, err = db.TaskLogs(t1In.TaskID, 5, []api.Filter{}, apiv1.OrderBy_ORDER_BY_UNSPECIFIED, nil)
	require.NoError(t, err)
	require.Equal(t, 2, len(logs))

	// Test DeleteTaskLogs.
	err = db.DeleteTaskLogs([]model.TaskID{t2In.TaskID})
	require.NoError(t, err)

	count, err = db.TaskLogsCount(t2In.TaskID, []api.Filter{})
	require.NoError(t, err)
	require.Equal(t, 0, count)
}

func RequireMockTaskLog(t *testing.T, db *PgDB, tID model.TaskID, suffix string) *model.TaskLog {
	mockA := RequireMockAllocation(t, db, tID)
	agentID := fmt.Sprintf("testing-agent-%s", suffix)
	containerID := suffix
	log := &model.TaskLog{
		TaskID:       string(tID),
		AllocationID: (*string)(&mockA.AllocationID),
		Log:          fmt.Sprintf("this is a log for task %s-%s", tID, suffix),
		AgentID:      &agentID,
		ContainerID:  &containerID,
	}
	return log
}
