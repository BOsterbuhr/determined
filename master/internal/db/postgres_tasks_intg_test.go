//go:build integration
// +build integration

package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/determined-ai/determined/master/internal/api"
	"github.com/determined-ai/determined/master/pkg/cproto"
	"github.com/determined-ai/determined/master/pkg/etc"
	"github.com/determined-ai/determined/master/pkg/model"
	"github.com/determined-ai/determined/master/pkg/ptrs"
	"github.com/determined-ai/determined/proto/pkg/apiv1"
	"github.com/determined-ai/determined/proto/pkg/taskv1"
)

// TestJobTaskAndAllocationAPI, in lieu of an ORM, ensures that the mappings into and out of the
// database are total. We should look into an ORM in the near to medium term future.
func TestJobTaskAndAllocationAPI(t *testing.T) {
	ctx := context.Background()
	require.NoError(t, etc.SetRootPath(RootFromDB))
	db := MustResolveTestPostgres(t)
	MustMigrateTestPostgres(t, db, MigrationsFromDB)

	// Add a mock user.
	user := RequireMockUser(t, db)

	// Add a job.
	jID := model.NewJobID()
	jIn := &model.Job{
		JobID:   jID,
		JobType: model.JobTypeExperiment,
		OwnerID: &user.ID,
		QPos:    decimal.New(0, 0),
	}
	err := db.AddJob(jIn)
	require.NoError(t, err, "failed to add job")

	// Retrieve it back and make sure the mapping is exhaustive.
	jOut, err := db.JobByID(jID)
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
	err = db.AddTask(tIn)
	require.NoError(t, err, "failed to add task")

	// Retrieve it back and make sure the mapping is exhaustive.
	tOut, err := TaskByID(ctx, tID)
	require.NoError(t, err, "failed to retrieve task")
	require.True(t, reflect.DeepEqual(tIn, tOut), pprintedExpect(tIn, tOut))

	// Complete it.
	tIn.EndTime = ptrs.Ptr(time.Now().UTC().Truncate(time.Millisecond))
	err = db.CompleteTask(tID, *tIn.EndTime)
	require.NoError(t, err, "failed to mark task completed")

	// Re-retrieve it back and make sure the mapping is still exhaustive.
	tOut, err = TaskByID(ctx, tID)
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
	err = db.AddAllocation(aIn)
	require.NoError(t, err, "failed to add allocation")

	// Update ports
	ports["dtrain_port"] = 0
	ports["inter_train_process_comm_port1"] = 0
	ports["inter_train_process_comm_port2"] = 0
	ports["c10d_port"] = 0
	aIn.Ports = ports
	err = UpdateAllocationPorts(*aIn)
	require.NoError(t, err, "failed to update port offset")

	// Retrieve it back and make sure the mapping is exhaustive.
	aOut, err := db.AllocationByID(aIn.AllocationID)
	require.NoError(t, err, "failed to retrieve allocation")
	require.True(t, reflect.DeepEqual(aIn, aOut), pprintedExpect(aIn, aOut))

	// Complete it.
	aIn.EndTime = ptrs.Ptr(time.Now().UTC().Truncate(time.Millisecond))
	err = db.CompleteAllocation(aIn)
	require.NoError(t, err, "failed to mark allocation completed")

	// Re-retrieve it back and make sure the mapping is still exhaustive.
	aOut, err = db.AllocationByID(aIn.AllocationID)
	require.NoError(t, err, "failed to re-retrieve allocation")
	require.True(t, reflect.DeepEqual(aIn, aOut), pprintedExpect(aIn, aOut))
}

func TestRecordAndEndTaskStats(t *testing.T) {
	require.NoError(t, etc.SetRootPath(RootFromDB))
	db := MustResolveTestPostgres(t)
	MustMigrateTestPostgres(t, db, MigrationsFromDB)

	tID := model.NewTaskID()
	require.NoError(t, db.AddTask(&model.Task{
		TaskID:    tID,
		TaskType:  model.TaskTypeTrial,
		StartTime: time.Now().UTC().Truncate(time.Millisecond),
	}), "failed to add task")

	allocationID := model.AllocationID(tID + "allocationID")
	require.NoError(t, db.AddAllocation(&model.Allocation{
		TaskID:       tID,
		AllocationID: allocationID,
	}), "failed to add allocation")

	var expected []*model.TaskStats
	for i := 0; i < 3; i++ {
		taskStats := &model.TaskStats{
			AllocationID: allocationID,
			EventType:    "IMAGEPULL",
			ContainerID:  ptrs.Ptr(cproto.NewID()),
			StartTime:    ptrs.Ptr(time.Now().Truncate(time.Millisecond)),
		}
		if i == 0 {
			taskStats.ContainerID = nil
		}
		require.NoError(t, RecordTaskStatsBun(taskStats))

		taskStats.EndTime = ptrs.Ptr(time.Now().Truncate(time.Millisecond))
		require.NoError(t, RecordTaskEndStatsBun(taskStats))
		expected = append(expected, taskStats)
	}

	var actual []*model.TaskStats
	err := Bun().NewSelect().
		Model(&actual).
		Where("allocation_id = ?", allocationID).
		Scan(context.TODO(), &actual)
	require.NoError(t, err)

	require.ElementsMatch(t, expected, actual)
}

func TestNonExperimentTasksContextDirectory(t *testing.T) {
	ctx := context.Background()
	require.NoError(t, etc.SetRootPath(RootFromDB))
	db := MustResolveTestPostgres(t)
	MustMigrateTestPostgres(t, db, MigrationsFromDB)

	// Task doesn't exist.
	_, err := NonExperimentTasksContextDirectory(ctx, model.TaskID(uuid.New().String()))
	require.ErrorIs(t, err, sql.ErrNoRows)

	// Nil context directory.
	tID := model.NewTaskID()
	require.NoError(t, db.AddTask(&model.Task{
		TaskID:    tID,
		TaskType:  model.TaskTypeNotebook,
		StartTime: time.Now().UTC().Truncate(time.Millisecond),
	}), "failed to add task")

	require.NoError(t, AddNonExperimentTasksContextDirectory(ctx, tID, nil))

	dir, err := NonExperimentTasksContextDirectory(ctx, tID)
	require.NoError(t, err)
	require.Len(t, dir, 0)

	// Non nil context directory.
	tID = model.NewTaskID()
	require.NoError(t, db.AddTask(&model.Task{
		TaskID:    tID,
		TaskType:  model.TaskTypeNotebook,
		StartTime: time.Now().UTC().Truncate(time.Millisecond),
	}), "failed to add task")

	expectedDir := []byte{3, 2, 1}
	require.NoError(t, AddNonExperimentTasksContextDirectory(ctx, tID, expectedDir))

	dir, err = NonExperimentTasksContextDirectory(ctx, tID)
	require.NoError(t, err)
	require.Equal(t, expectedDir, dir)
}

func TestAllocationState(t *testing.T) {
	require.NoError(t, etc.SetRootPath(RootFromDB))
	db := MustResolveTestPostgres(t)
	MustMigrateTestPostgres(t, db, MigrationsFromDB)

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
		require.NoError(t, db.AddTask(task), "failed to add task")

		s := state
		a := &model.Allocation{
			TaskID:       tID,
			AllocationID: model.AllocationID(tID + "allocationID"),
			ResourcePool: "default",
			State:        &s,
		}
		require.NoError(t, db.AddAllocation(a), "failed to add allocation")

		// Update allocation to every possible state.
		testNoUpdate := true
		for j := 0; j < len(states); j++ {
			if testNoUpdate {
				testNoUpdate = false
				j-- // Go to first iteration of loop after this.
			} else {
				a.State = &states[j]
				require.NoError(t, db.UpdateAllocationState(*a),
					"failed to update allocation state")
			}

			// Get task back as a proto struct.
			tOut := &taskv1.Task{}
			require.NoError(t, db.QueryProto("get_task", tOut, tID), "failed to get task")

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

func TestExhaustiveEnums(t *testing.T) {
	require.NoError(t, etc.SetRootPath(RootFromDB))
	db := MustResolveTestPostgres(t)
	MustMigrateTestPostgres(t, db, MigrationsFromDB)

	type check struct {
		goType          string
		goMembers       map[string]bool
		postgresType    string
		postgresMembers map[string]bool
		ignore          map[string]bool
	}
	checks := map[string]*check{}
	addCheck := func(goType, postgresType string, ignore map[string]bool) {
		checks[goType] = &check{
			goType:          goType,
			goMembers:       map[string]bool{},
			postgresType:    postgresType,
			postgresMembers: map[string]bool{},
			ignore:          ignore,
		}
	}
	addCheck("JobType", "public.job_type", map[string]bool{})
	addCheck("TaskType", "public.task_type", map[string]bool{})
	addCheck("State", "public.experiment_state", map[string]bool{
		"PARTIALLY_DELETED": true,
		"DELETED":           true,
	})
	addCheck("AllocationState", "public.allocation_state", map[string]bool{})

	// Populate postgres types.
	for _, c := range checks {
		q := fmt.Sprintf("SELECT unnest(enum_range(NULL::%s))::text", c.postgresType)
		rows, err := db.sql.Queryx(q)
		require.NoError(t, err, "querying postgres enum members")
		defer rows.Close()
		for rows.Next() {
			var text string
			require.NoError(t, rows.Scan(&text), "scanning enum value")
			c.postgresMembers[text] = true
		}
	}

	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, "../../pkg/model", nil, parser.ParseComments)
	require.NoError(t, err)
	for _, p := range pkgs {
		for _, f := range p.Files {
			ast.Inspect(f, func(n ast.Node) bool {
				vs, ok := n.(*ast.ValueSpec)
				if !ok {
					return true
				}

				vsTypeIdent, ok := vs.Type.(*ast.Ident)
				if !ok {
					return true
				}

				c, ok := checks[vsTypeIdent.Name]
				if !ok {
					return true
				}

				// We can error out now because we're certainly on something we want to check.
				for _, v := range vs.Values {
					bl, ok := v.(*ast.BasicLit)
					require.True(t, ok, "linter can only handle pg enums as basic lits")
					require.Equal(t, token.STRING, bl.Kind, "linter can only handle lit strings")
					c.goMembers[strings.Trim(bl.Value, "\"'`")] = true
				}

				return true
			})
		}
	}

	for _, c := range checks {
		for name := range c.ignore {
			delete(c.postgresMembers, name)
			delete(c.goMembers, name)
		}

		pb, err := json.Marshal(c.postgresMembers)
		require.NoError(t, err)
		gb, err := json.Marshal(c.goMembers)
		require.NoError(t, err)

		// Gives pretty diff.
		require.JSONEq(t, string(pb), string(gb))
	}
}

func TestAddTask(t *testing.T) {
	require.NoError(t, etc.SetRootPath(RootFromDB))
	db := MustResolveTestPostgres(t)
	MustMigrateTestPostgres(t, db, MigrationsFromDB)
	ctx := context.Background()

	u := RequireMockUser(t, db)

	jID := model.NewJobID()
	jIn := &model.Job{
		JobID:   jID,
		JobType: model.JobTypeExperiment,
		OwnerID: &u.ID,
		QPos:    decimal.New(0, 0),
	}
	err := AddJobTx(ctx, Bun(), jIn)
	require.NoError(t, err, "failed to add job")

	// Add a task.
	tID := model.NewTaskID()
	tIn := &model.Task{
		TaskID:     tID,
		JobID:      &jID,
		TaskType:   model.TaskTypeTrial,
		StartTime:  time.Now().UTC().Truncate(time.Millisecond),
		LogVersion: model.TaskLogVersion0,
	}
	err = AddTask(context.Background(), tIn)
	require.NoError(t, err, "failed to add task")
}

func TestTaskByID(t *testing.T) {
	mockT := mockTask(t)

	task, err := TaskByID(context.Background(), mockT.TaskID)
	require.NoError(t, err)
	require.Equal(t, mockT, task)
}

func TestAddNonExperimentTasksContextDirectory(t *testing.T) {
	ctx := context.Background()
	mockT := mockTask(t)
	b := []byte(`testing123`)

	err := AddNonExperimentTasksContextDirectory(ctx, mockT.TaskID, b)
	require.NoError(t, err)

	var taskCtxDir model.TaskContextDirectory
	err = Bun().NewSelect().Model(&taskCtxDir).Where("task_id = ?", mockT.TaskID).Scan(ctx, &taskCtxDir)
	require.NoError(t, err)
	require.Equal(t, b, taskCtxDir.ContextDirectory)
}

func TestTaskCompleted(t *testing.T) {
	ctx := context.Background()
	require.NoError(t, etc.SetRootPath(RootFromDB))
	db := MustResolveTestPostgres(t)
	MustMigrateTestPostgres(t, db, MigrationsFromDB)

	mockT := mockTask(t)

	completed, err := TaskCompleted(ctx, mockT.TaskID)
	require.False(t, completed)
	require.NoError(t, err)

	err = db.CompleteTask(mockT.TaskID, time.Now().UTC().Truncate(time.Millisecond))
	require.NoError(t, err)

	completed, err = TaskCompleted(ctx, mockT.TaskID)
	require.True(t, completed)
	require.NoError(t, err)
}

func TestAddAllocation(t *testing.T) {
	require.NoError(t, etc.SetRootPath(RootFromDB))
	db := MustResolveTestPostgres(t)
	MustMigrateTestPostgres(t, db, MigrationsFromDB)

	mockT := mockTask(t)

	a := model.Allocation{
		AllocationID: model.AllocationID(fmt.Sprintf("%s-1", mockT.TaskID)),
		TaskID:       mockT.TaskID,
		StartTime:    ptrs.Ptr(time.Now().UTC()),
		State:        ptrs.Ptr(model.AllocationStateTerminated),
	}
	err := db.AddAllocation(&a)
	require.NoError(t, err, "failed to add allocation")

	logrus.Debugf("%s", Bun().NewSelect().Table("allocations"))

	var tmp model.Allocation
	err = Bun().NewSelect().Table("allocations").Where("allocation_id = ?", string(a.AllocationID)).
		Scan(context.Background(), &tmp)
	require.NoError(t, err)
	require.Equal(t, a.AllocationID, tmp.AllocationID)
	require.Equal(t, a.TaskID, tmp.TaskID)
	require.Equal(t, a.StartTime, tmp.StartTime)
	require.Equal(t, a.State, tmp.State)
}

func TestAddAllocationExitStatus(t *testing.T) {
	mockT := mockTask(t)
	mockA := mockAllocation(t, mockT.TaskID)

	statusCode := int32(1)
	exitReason := "testing-exit-reason"
	exitErr := "testing-exit-err"

	mockA.ExitReason = &exitReason
	mockA.ExitErr = &exitErr
	mockA.StatusCode = &statusCode

	err := AddAllocationExitStatus(context.Background(), mockA)
	require.NoError(t, err)

	alloc := getAllocationByID(t, mockA.AllocationID)
	require.Equal(t, mockA.ExitErr, alloc.ExitErr)
	require.Equal(t, mockA.ExitReason, alloc.ExitReason)
	require.Equal(t, mockA.StatusCode, alloc.StatusCode)
}

func TestCompleteAllocation(t *testing.T) {
	require.NoError(t, etc.SetRootPath(RootFromDB))
	db := MustResolveTestPostgres(t)
	MustMigrateTestPostgres(t, db, MigrationsFromDB)

	mockT := mockTask(t)
	mockA := mockAllocation(t, mockT.TaskID)

	mockA.EndTime = ptrs.Ptr(time.Now().UTC())

	err := db.CompleteAllocation(mockA)
	require.NoError(t, err)

	alloc := getAllocationByID(t, mockA.AllocationID)
	require.Equal(t, mockA.EndTime, alloc.EndTime)
}

func TestCompleteAllocationTelemetry(t *testing.T) {
	require.NoError(t, etc.SetRootPath(RootFromDB))
	db := MustResolveTestPostgres(t)
	MustMigrateTestPostgres(t, db, MigrationsFromDB)

	mockT := mockTask(t)
	mockA := mockAllocation(t, mockT.TaskID)

	err := db.AddAllocation(mockA)
	require.NoError(t, err, "failed to add allocation")
	bytes, err := db.CompleteAllocationTelemetry(mockA.AllocationID)
	require.NoError(t, err)
	require.Contains(t, string(bytes), string(mockA.AllocationID))
	require.Contains(t, string(bytes), string(*mockT.JobID))
	require.Contains(t, string(bytes), string(mockT.TaskType))
}

func TestAllocationByID(t *testing.T) {
	require.NoError(t, etc.SetRootPath(RootFromDB))
	db := MustResolveTestPostgres(t)
	MustMigrateTestPostgres(t, db, MigrationsFromDB)

	mockT := RequireMockTask(t)
	mockA := RequireMockAllocation(t, db, mockT.TaskID)

	a, err := db.AllocationByID(mockA.AllocationID)
	require.NoError(t, err)
	require.Equal(t, mockA, a)
}

func TestAllocationSessionFlow(t *testing.T) {
	require.NoError(t, etc.SetRootPath(RootFromDB))
	db := MustResolveTestPostgres(t)
	MustMigrateTestPostgres(t, db, MigrationsFromDB)

	mockU := RequireMockUser(t, db)
	mockT := RequireMockTask(t)
	mockA := RequireMockAllocation(t, db, mockT.TaskID)

	tok, err := db.StartAllocationSession(mockA.AllocationID, &mockU)
	require.NoError(t, err)
	require.NotNil(t, tok)

	alloc := getAllocationByID(t, mockA.AllocationID)
	require.Equal(t, mockA, alloc)

	running := model.AllocationStateRunning
	mockA.State = &running
	err = db.UpdateAllocationState(*mockA)
	require.NoError(t, err)

	alloc = getAllocationByID(t, mockA.AllocationID)
	require.Equal(t, mockA.State, alloc.State)

	err = db.DeleteAllocationSession(mockA.AllocationID)
	require.NoError(t, err)
}

func TestUpdateAllocation(t *testing.T) {
	require.NoError(t, etc.SetRootPath(RootFromDB))
	db := MustResolveTestPostgres(t)
	MustMigrateTestPostgres(t, db, MigrationsFromDB)

	mockT := RequireMockTask(t)
	mockA := RequireMockAllocation(t, db, mockT.TaskID)

	// Testing UpdateAllocation Ports
	mockA.Ports = map[string]int{"abc": 123, "def": 456}

	err := UpdateAllocationPorts(*mockA)
	require.NoError(t, err)

	alloc := getAllocationByID(t, mockA.AllocationID)
	require.Equal(t, mockA.Ports, alloc.Ports)

	// Testing UpdateAllocationStartTime
	newStartTime := ptrs.Ptr(time.Now().UTC())
	mockA.StartTime = newStartTime

	err = db.UpdateAllocationStartTime(*mockA)
	require.NoError(t, err)

	alloc = getAllocationByID(t, mockA.AllocationID)
	require.Equal(t, mockA.StartTime, alloc.StartTime)

	// Testing UpdateAllocationProxyAddress
	proxyAddr := "here"
	mockA.ProxyAddress = &proxyAddr

	err = db.UpdateAllocationProxyAddress(*mockA)
	require.NoError(t, err)

	alloc = getAllocationByID(t, mockA.AllocationID)
	require.Equal(t, mockA.ProxyAddress, alloc.ProxyAddress)
}

func TestCloseOpenAllocations(t *testing.T) {
	require.NoError(t, etc.SetRootPath(RootFromDB))
	db := MustResolveTestPostgres(t)
	MustMigrateTestPostgres(t, db, MigrationsFromDB)

	mockT := mockTask(t)
	mockA1 := mockAllocation(t, mockT.TaskID)
	mockA2 := mockAllocation(t, mockT.TaskID)

	terminated := model.AllocationStateTerminated

	mockA2.State = &terminated
	mockA2.State = &terminated

	err := db.CloseOpenAllocations([]model.AllocationID{mockA1.AllocationID})
	require.NoError(t, err)

	alloc1 := getAllocationByID(t, mockA1.AllocationID)
	require.Nil(t, alloc1.EndTime)

	alloc2 := getAllocationByID(t, mockA2.AllocationID)
	logrus.Debugf("ALLOCATION 2: %v", alloc2)
	require.NotNil(t, alloc2.EndTime)

	err = db.CloseOpenAllocations([]model.AllocationID{})
	require.NoError(t, err)

	alloc1 = getAllocationByID(t, mockA1.AllocationID)
	require.NotNil(t, alloc1.EndTime)

	alloc2 = getAllocationByID(t, mockA2.AllocationID)
	require.NotNil(t, alloc2.EndTime)
}

func TestTaskLogsFlow(t *testing.T) {
	require.NoError(t, etc.SetRootPath(RootFromDB))
	db := MustResolveTestPostgres(t)
	MustMigrateTestPostgres(t, db, MigrationsFromDB)

	mockT1 := mockTask(t)
	mockT2 := mockTask(t)

	// Test AddTaskLogs & TaskLogCounts
	taskLog1 := RequireMockTaskLog(t, mockT1.TaskID, "1")
	taskLog2 := RequireMockTaskLog(t, mockT1.TaskID, "2")
	taskLog3 := RequireMockTaskLog(t, mockT2.TaskID, "3")

	err := db.AddTaskLogs([]*model.TaskLog{taskLog1})
	require.NoError(t, err)

	count, err := db.TaskLogsCount(mockT1.TaskID, []api.Filter{})
	require.NoError(t, err)
	require.Equal(t, count, 1)

	err = db.AddTaskLogs([]*model.TaskLog{taskLog2, taskLog3})
	require.NoError(t, err)

	count, err = db.TaskLogsCount(mockT1.TaskID, []api.Filter{})
	require.NoError(t, err)
	require.Equal(t, count, 2)

	// Test TaskLogs.
	logs, _, err := db.TaskLogs(mockT1.TaskID, 1, []api.Filter{}, apiv1.OrderBy_ORDER_BY_UNSPECIFIED, nil)
	require.NoError(t, err)
	require.Equal(t, 1, len(logs))
	require.Equal(t, logs[0].TaskID, string(mockT1.TaskID))
	require.Contains(t, []string{"1", "2"}, *logs[0].ContainerID)

	logs, _, err = db.TaskLogs(mockT1.TaskID, 5, []api.Filter{}, apiv1.OrderBy_ORDER_BY_UNSPECIFIED, nil)
	require.NoError(t, err)
	require.Equal(t, 2, len(logs))

	// Test DeleteTaskLogs.
	err = db.DeleteTaskLogs([]model.TaskID{mockT2.TaskID})
	require.NoError(t, err)

	count, err = db.TaskLogsCount(mockT2.TaskID, []api.Filter{})
	require.NoError(t, err)
	require.Equal(t, 0, count)
}

// TODO CAROLINA: write tests for EndAllTaskStats & TaskLogsFields

// mockTask returns a mock task.
func mockTask(t *testing.T) *model.Task {
	require.NoError(t, etc.SetRootPath(RootFromDB))
	db := MustResolveTestPostgres(t)
	MustMigrateTestPostgres(t, db, MigrationsFromDB)
	ctx := context.Background()

	u := RequireMockUser(t, db)

	jID := model.NewJobID()
	jIn := &model.Job{
		JobID:   jID,
		JobType: model.JobTypeExperiment,
		OwnerID: &u.ID,
		QPos:    decimal.New(0, 0),
	}
	err := AddJobTx(ctx, Bun(), jIn)
	require.NoError(t, err, "failed to add job")

	// Add a task.
	tID := model.NewTaskID()
	tIn := &model.Task{
		TaskID:     tID,
		JobID:      &jID,
		TaskType:   model.TaskTypeTrial,
		StartTime:  time.Now().UTC().Truncate(time.Millisecond),
		LogVersion: model.TaskLogVersion0,
	}
	err = AddTask(context.Background(), tIn)
	require.NoError(t, err, "failed to add task")
	return tIn
}

func mockAllocation(t *testing.T, tID model.TaskID) *model.Allocation {
	require.NoError(t, etc.SetRootPath(RootFromDB))
	db := MustResolveTestPostgres(t)
	MustMigrateTestPostgres(t, db, MigrationsFromDB)

	a := model.Allocation{
		AllocationID: model.AllocationID(fmt.Sprintf("%s-%s", tID, uuid.NewString())),
		TaskID:       tID,
		StartTime:    ptrs.Ptr(time.Now().UTC()),
		State:        ptrs.Ptr(model.AllocationStateTerminated),
	}
	err := db.AddAllocation(&a)
	require.NoError(t, err, "failed to add allocation")
	return &a
}

func getAllocationByID(t *testing.T, aID model.AllocationID) *model.Allocation {
	var alloc model.Allocation
	err := Bun().NewSelect().Table("allocations").
		Where("allocation_id = ?", aID).Scan(context.Background(), &alloc)
	require.NoError(t, err)
	return &alloc
}

func RequireMockTaskLog(t *testing.T, tID model.TaskID, suffix string) *model.TaskLog {
	mockA := mockAllocation(t, tID)
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
