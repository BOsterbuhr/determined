//go:build integration
// +build integration

package internal

import (
	"context"
	"crypto/rand"
	"fmt"
	"testing"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
	"github.com/jonboulle/clockwork"
	"github.com/stretchr/testify/require"

	"github.com/determined-ai/determined/master/internal/db"
	"github.com/determined-ai/determined/master/internal/logretention"
	"github.com/determined-ai/determined/master/pkg/model"
	"github.com/determined-ai/determined/master/pkg/ptrs"
	"github.com/determined-ai/determined/master/pkg/schemas"
	"github.com/determined-ai/determined/master/pkg/schemas/expconf"
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

// nolint: exhaustruct
func createTestRetentionExpWithProjectID(
	t *testing.T, api *apiServer, curUser model.User, retentionDays *int16, projectID int, labels ...string,
) *model.Experiment {
	labelMap := make(map[string]bool)
	for _, l := range labels {
		labelMap[l] = true
	}

	activeConfig := schemas.Merge(minExpConfig, expconf.ExperimentConfig{
		RawLabels:           labelMap,
		RawDescription:      ptrs.Ptr("desc"),
		RawName:             expconf.Name{RawString: ptrs.Ptr("name")},
		RawLogRetentionDays: retentionDays,
	})
	activeConfig = schemas.WithDefaults(activeConfig)
	exp := &model.Experiment{
		JobID:                model.JobID(uuid.New().String()),
		State:                model.PausedState,
		OwnerID:              &curUser.ID,
		ProjectID:            projectID,
		StartTime:            time.Now(),
		ModelDefinitionBytes: []byte{10, 11, 12},
		Config:               activeConfig.AsLegacy(),
	}
	require.NoError(t, api.m.db.AddExperiment(exp, activeConfig))

	// Get experiment as our API mostly will to make it easier to mock.
	exp, err := db.ExperimentByID(context.TODO(), exp.ID)
	require.NoError(t, err)
	return exp
}

func createTestRetentionTrial(
	t *testing.T, api *apiServer, curUser model.User, retentionDays *int16,
) (*model.Trial, *model.Task) {
	exp := createTestRetentionExpWithProjectID(t, api, curUser, retentionDays, 1)

	task := &model.Task{
		TaskType:         model.TaskTypeTrial,
		LogVersion:       model.TaskLogVersion1,
		LogRetentionDays: retentionDays,
		StartTime:        time.Now(),
		TaskID:           trialTaskID(exp.ID, model.NewRequestID(rand.Reader)),
	}
	require.NoError(t, api.m.db.AddTask(task))

	trial := &model.Trial{
		StartTime:    time.Now(),
		State:        model.PausedState,
		ExperimentID: exp.ID,
	}
	require.NoError(t, db.AddTrial(context.TODO(), trial, task.TaskID))

	// Return trial exactly the way the API will generally get it.
	outTrial, err := db.TrialByID(context.TODO(), trial.ID)
	require.NoError(t, err)
	return outTrial, task
}

func TestDeleteExpiredTaskLogs(t *testing.T) {
	api, curUser, ctx := setupAPITest(t, nil)
	logRetentionDays := int16(30)
	trial, task0 := createTestRetentionTrial(t, api, curUser, &logRetentionDays)

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

func TestScheduleRetention(t *testing.T) {
	fakeClock := clockwork.NewFakeClock()
	logretention.SetupScheduler(gocron.WithClock(fakeClock))
}
