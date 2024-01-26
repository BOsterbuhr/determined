//go:build integration
// +build integration

package internal

import (
	"context"
	"crypto/rand"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/determined-ai/determined/master/internal/db"
	"github.com/determined-ai/determined/master/pkg/model"
)

func setRetentionTime(timestamp string) error {
	_, err := db.Bun().NewRaw(fmt.Sprintf(`
	CREATE FUNCTION retention_timestamp() RETURNS TIMESTAMPTZ AS $$ 
    BEGIN
        RETURN %s;
    END 
    $$ LANGUAGE PLPGSQL;
	`, timestamp)).Exec(context.Background())
	return err
}

func resetRetentionTime() error {
	return setRetentionTime("transaction_timestamp()")
}

func TestDeleteExpiredTaskLogs(t *testing.T) {
	api, curUser, ctx := setupAPITest(t, nil)
	trial, task0 := createTestTrial(t, api, curUser)

	task1 := &model.Task{
		TaskType:   model.TaskTypeTrial,
		LogVersion: model.TaskLogVersion1,
		StartTime:  task0.StartTime.Add(time.Second),
		TaskID:     trialTaskID(trial.ExperimentID, model.NewRequestID(rand.Reader)),
	}
	require.NoError(t, api.m.db.AddTask(task1))

	task2 := &model.Task{
		TaskType:   model.TaskTypeTrial,
		LogVersion: model.TaskLogVersion1,
		StartTime:  task1.StartTime.Add(time.Second),
		TaskID:     trialTaskID(trial.ExperimentID, model.NewRequestID(rand.Reader)),
	}
	require.NoError(t, api.m.db.AddTask(task2))

	_, err := db.Bun().NewInsert().Model(&[]model.TrialTaskID{
		{TrialID: trial.ID, TaskID: task1.TaskID},
		{TrialID: trial.ID, TaskID: task2.TaskID},
	}).Exec(ctx)
	require.NoError(t, err)

	require.NoError(t, resetRetentionTime())
}
