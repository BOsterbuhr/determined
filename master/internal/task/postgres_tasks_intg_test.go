package task

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slices"

	"github.com/determined-ai/determined/master/internal/db"
	"github.com/determined-ai/determined/master/pkg/etc"
	"github.com/determined-ai/determined/master/pkg/model"
	"github.com/determined-ai/determined/master/pkg/ptrs"
	"github.com/determined-ai/determined/proto/pkg/taskv1"
)

// RootFromDB returns the relative path from db to root.
const RootFromDB = "../../static/srv"

func TestAllocationState(t *testing.T) {
	require.NoError(t, etc.SetRootPath(RootFromDB))
	pgDB := db.MustResolveTestPostgres(t)
	db.MustMigrateTestPostgres(t, pgDB, db.MigrationsFromDB)

	// Add an allocation of every possible state.
	states := []model.AllocationState{
		model.AllocationStatePending,
		model.AllocationStateAssigned,
		model.AllocationStatePulling,
		model.AllocationStateStarting,
		model.AllocationStateRunning,
		model.AllocationStateTerminating,
		model.AllocationStateTerminated,
	}
	for _, state := range states {
		tID := model.NewTaskID()
		task := &model.Task{
			TaskID:    tID,
			TaskType:  model.TaskTypeTrial,
			StartTime: time.Now().UTC().Truncate(time.Millisecond),
		}
		require.NoError(t, db.AddTask(context.Background(), task), "failed to add task")

		s := state
		a := &model.Allocation{
			TaskID:       tID,
			AllocationID: model.AllocationID(tID + "allocationID"),
			ResourcePool: "default",
			State:        &s,
		}
		require.NoError(t, pgDB.AddAllocation(a), "failed to add allocation")

		// Update allocation to every possible state.
		testNoUpdate := true
		for j := 0; j < len(states); j++ {
			if testNoUpdate {
				testNoUpdate = false
				j-- // Go to first iteration of loop after this.
			} else {
				a.State = &states[j]
				require.NoError(t, UpdateAllocationState(context.Background(), *a),
					"failed to update allocation state")
			}

			// Get task back as a proto struct.
			tOut := &taskv1.Task{}
			require.NoError(t, pgDB.QueryProto("get_task", tOut, tID), "failed to get task")

			// Ensure our state is the same as allocation.
			require.Equal(t, len(tOut.Allocations), 1, "failed to get exactly 1 allocation")
			aOut := tOut.Allocations[0]

			if slices.Contains([]model.AllocationState{
				model.AllocationStatePending,
				model.AllocationStateAssigned,
			}, *a.State) {
				require.Equal(t, "STATE_QUEUED", aOut.State.String(),
					"allocation states not converted to queued")
			} else {
				require.Equal(t, a.State.Proto(), aOut.State, "proto state not equal")
				require.Equal(t, fmt.Sprintf("STATE_%s", *a.State), aOut.State.String(),
					"proto state to strings not equal")
			}
		}
	}
}

// TestJobTaskAndAllocationAPI, in lieu of an ORM, ensures that the mappings into and out of the
// database are total. We should look into an ORM in the near to medium term future.
func TestJobTaskAndAllocationAPI(t *testing.T) {
	ctx := context.Background()
	require.NoError(t, etc.SetRootPath(RootFromDB))
	pgDB := db.MustResolveTestPostgres(t)
	db.MustMigrateTestPostgres(t, pgDB, db.MigrationsFromDB)

	// Add a mock user.
	user := db.RequireMockUser(t, pgDB)

	// Add a job.
	jID := model.NewJobID()
	jIn := &model.Job{
		JobID:   jID,
		JobType: model.JobTypeExperiment,
		OwnerID: &user.ID,
		QPos:    decimal.New(0, 0),
	}
	err := pgDB.AddJob(jIn)
	require.NoError(t, err, "failed to add job")

	// Retrieve it back and make sure the mapping is exhaustive.
	jOut, err := pgDB.JobByID(jID)
	require.NoError(t, err, "failed to retrieve job")
	require.True(t, reflect.DeepEqual(jIn, jOut), pprintedExpect(jIn, jOut))

	// Add a task.
	tID := model.NewTaskID()
	tIn := &model.Task{
		TaskID:    tID,
		JobID:     &jID,
		TaskType:  model.TaskTypeTrial,
		StartTime: time.Now().UTC().Truncate(time.Millisecond),
	}
	err = db.AddTask(context.Background(), tIn)
	require.NoError(t, err, "failed to add task")

	// Retrieve it back and make sure the mapping is exhaustive.
	tOut, err := db.TaskByID(ctx, tID)
	require.NoError(t, err, "failed to retrieve task")
	require.True(t, reflect.DeepEqual(tIn, tOut), pprintedExpect(tIn, tOut))

	// Complete it.
	tIn.EndTime = ptrs.Ptr(time.Now().UTC().Truncate(time.Millisecond))
	err = pgDB.CompleteTask(tID, *tIn.EndTime)
	require.NoError(t, err, "failed to mark task completed")

	// Re-retrieve it back and make sure the mapping is still exhaustive.
	tOut, err = db.TaskByID(ctx, tID)
	require.NoError(t, err, "failed to re-retrieve task")
	require.True(t, reflect.DeepEqual(tIn, tOut), pprintedExpect(tIn, tOut))

	// And an allocation.
	ports := map[string]int{}
	ports["dtrain_port"] = 0
	ports["inter_train_process_comm_port1"] = 0
	ports["inter_train_process_comm_port2"] = 0
	ports["c10d_port"] = 0

	aID := model.AllocationID(string(tID) + "-1")
	aIn := &model.Allocation{
		AllocationID: aID,
		TaskID:       tID,
		Slots:        8,
		ResourcePool: "somethingelse",
		StartTime:    ptrs.Ptr(time.Now().UTC().Truncate(time.Millisecond)),
		Ports:        ports,
	}
	err = pgDB.AddAllocation(aIn)
	require.NoError(t, err, "failed to add allocation")

	// Update ports
	ports["dtrain_port"] = 0
	ports["inter_train_process_comm_port1"] = 0
	ports["inter_train_process_comm_port2"] = 0
	ports["c10d_port"] = 0
	aIn.Ports = ports
	err = UpdateAllocationPorts(context.Background(), *aIn)
	require.NoError(t, err, "failed to update port offset")

	// Retrieve it back and make sure the mapping is exhaustive.
	aOut, err := AllocationByID(context.Background(), aIn.AllocationID)
	require.NoError(t, err, "failed to retrieve allocation")
	require.True(t, reflect.DeepEqual(aIn, aOut), pprintedExpect(aIn, aOut))

	// Complete it.
	aIn.EndTime = ptrs.Ptr(time.Now().UTC().Truncate(time.Millisecond))
	err = pgDB.CompleteAllocation(aIn)
	require.NoError(t, err, "failed to mark allocation completed")

	// Re-retrieve it back and make sure the mapping is still exhaustive.
	aOut, err = AllocationByID(context.Background(), aIn.AllocationID)
	require.NoError(t, err, "failed to re-retrieve allocation")
	require.True(t, reflect.DeepEqual(aIn, aOut), pprintedExpect(aIn, aOut))
}

func pprintedExpect(expected, got interface{}) string {
	return fmt.Sprintf("expected \n\t%s\ngot\n\t%s", spew.Sdump(expected), spew.Sdump(got))
}

func TestClusterAPI(t *testing.T) {
	require.NoError(t, etc.SetRootPath(db.RootFromDB))

	pgDB := db.MustResolveTestPostgres(t)
	db.MustMigrateTestPostgres(t, pgDB, db.MigrationsFromDB)

	_, err := pgDB.GetOrCreateClusterID("")
	require.NoError(t, err, "failed to get or create cluster id")

	// Add a mock user
	user := db.RequireMockUser(t, pgDB)

	// Add a job
	jID := model.NewJobID()
	jIn := &model.Job{
		JobID:   jID,
		JobType: model.JobTypeExperiment,
		OwnerID: &user.ID,
	}

	err = pgDB.AddJob(jIn)
	require.NoError(t, err, "failed to add job")

	// Add a task
	tID := model.NewTaskID()
	tIn := &model.Task{
		TaskID:    tID,
		JobID:     &jID,
		TaskType:  model.TaskTypeTrial,
		StartTime: time.Now().UTC().Truncate(time.Millisecond),
	}

	err = db.AddTask(context.Background(), tIn)
	require.NoError(t, err, "failed to add task")

	// Add an allocation
	aID := model.AllocationID(string(tID) + "-1")
	aIn := &model.Allocation{
		AllocationID: aID,
		TaskID:       tID,
		Slots:        8,
		ResourcePool: "somethingelse",
		StartTime:    ptrs.Ptr(time.Now().UTC().Truncate(time.Millisecond)),
	}

	err = pgDB.AddAllocation(aIn)
	require.NoError(t, err, "failed to add allocation")

	// Add a cluster heartbeat after allocation, so it is as if the master died with it open.
	currentTime := time.Now().UTC().Truncate(time.Millisecond)
	require.NoError(t, pgDB.UpdateClusterHeartBeat(currentTime))

	var clusterHeartbeat time.Time
	err = db.Bun().NewSelect().Table("cluster_id").Column("cluster_heartbeat").
		Scan(context.Background(), &clusterHeartbeat)
	require.NoError(t, err, "error reading cluster_heartbeat from cluster_id table")

	require.Equal(t, currentTime, clusterHeartbeat,
		"Retrieved cluster heartbeat doesn't match the correct time")

	// Don't complete the above allocation and call CloseOpenAllocations
	require.NoError(t, pgDB.CloseOpenAllocations(nil))

	// Retrieve the open allocation and check if end time is set to cluster_heartbeat
	aOut, err := AllocationByID(context.Background(), aIn.AllocationID)
	require.NoError(t, err)
	require.NotNil(t, aOut, "aOut is Nil")
	require.NotNil(t, aOut.EndTime, "aOut.EndTime is Nil")
	require.Equal(t, *aOut.EndTime, clusterHeartbeat,
		"Expected end time of open allocation is = %q but it is = %q instead",
		clusterHeartbeat.String(), aOut.EndTime.String())
}
