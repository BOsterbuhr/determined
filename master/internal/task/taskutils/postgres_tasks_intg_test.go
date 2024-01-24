package taskutils

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slices"

	"github.com/determined-ai/determined/master/internal/db"
	"github.com/determined-ai/determined/master/internal/user"
	"github.com/determined-ai/determined/master/pkg/cproto"
	"github.com/determined-ai/determined/master/pkg/etc"
	"github.com/determined-ai/determined/master/pkg/model"
	"github.com/determined-ai/determined/master/pkg/ptrs"
	"github.com/determined-ai/determined/proto/pkg/taskv1"
)

// RootFromDB returns the relative path from db to root.
const RootFromDB = "../../../static/srv"

var pgDB *db.PgDB

func TestMain(m *testing.M) {
	pgDB, err := db.ResolveTestPostgres()
	if err != nil {
		log.Panicln(err)
	}

	err = db.MigrateTestPostgres(pgDB, "file://../../../static/migrations", "up")
	if err != nil {
		log.Panicln(err)
	}

	err = etc.SetRootPath(RootFromDB)
	if err != nil {
		log.Panicln(err)
	}

	os.Exit(m.Run())
}

func TestAllocationState(t *testing.T) {
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
		require.NoError(t, AddTask(context.Background(), task), "failed to add task")

		s := state
		a := &model.Allocation{
			TaskID:       tID,
			AllocationID: model.AllocationID(tID + "allocationID"),
			ResourcePool: "default",
			State:        &s,
		}
		require.NoError(t, AddAllocation(a), "failed to add allocation")

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
	err = AddTask(context.Background(), tIn)
	require.NoError(t, err, "failed to add task")

	// Retrieve it back and make sure the mapping is exhaustive.
	tOut, err := TaskByID(ctx, tID)
	require.NoError(t, err, "failed to retrieve task")
	require.True(t, reflect.DeepEqual(tIn, tOut), pprintedExpect(tIn, tOut))

	// Complete it.
	tIn.EndTime = ptrs.Ptr(time.Now().UTC().Truncate(time.Millisecond))
	err = CompleteTask(tID, *tIn.EndTime)
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
	err = AddAllocation(aIn)
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
	err = CompleteAllocation(aIn)
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

	err = AddTask(context.Background(), tIn)
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

	err = AddAllocation(aIn)
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
	require.NoError(t, CloseOpenAllocations(nil))

	// Retrieve the open allocation and check if end time is set to cluster_heartbeat
	aOut, err := AllocationByID(context.Background(), aIn.AllocationID)
	require.NoError(t, err)
	require.NotNil(t, aOut, "aOut is Nil")
	require.NotNil(t, aOut.EndTime, "aOut.EndTime is Nil")
	require.Equal(t, *aOut.EndTime, clusterHeartbeat,
		"Expected end time of open allocation is = %q but it is = %q instead",
		clusterHeartbeat.String(), aOut.EndTime.String())
}

func TestRecordAndEndTaskStats(t *testing.T) {
	tID := model.NewTaskID()
	require.NoError(t, AddTask(context.Background(), &model.Task{
		TaskID:    tID,
		TaskType:  model.TaskTypeTrial,
		StartTime: time.Now().UTC().Truncate(time.Millisecond),
	}), "failed to add task")

	allocationID := model.AllocationID(tID + "allocationID")
	require.NoError(t, AddAllocation(&model.Allocation{
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
		require.NoError(t, db.RecordTaskStatsBun(taskStats))

		taskStats.EndTime = ptrs.Ptr(time.Now().Truncate(time.Millisecond))
		require.NoError(t, db.RecordTaskEndStatsBun(taskStats))
		expected = append(expected, taskStats)
	}

	var actual []*model.TaskStats
	err := db.Bun().NewSelect().
		Model(&actual).
		Where("allocation_id = ?", allocationID).
		Scan(context.TODO(), &actual)
	require.NoError(t, err)

	require.ElementsMatch(t, expected, actual)
}

func TestNonExperimentTasksContextDirectory(t *testing.T) {
	ctx := context.Background()

	// Task doesn't exist.
	_, err := NonExperimentTasksContextDirectory(ctx, model.TaskID(uuid.New().String()))
	require.ErrorIs(t, err, sql.ErrNoRows)

	// Nil context directory.
	tID := model.NewTaskID()
	require.NoError(t, AddTask(context.Background(), &model.Task{
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
	require.NoError(t, AddTask(context.Background(), &model.Task{
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

/*
func TestExhaustiveEnums(t *testing.T) {
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
		rows, err := pgDB.sql.Queryx(q)
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
*/

func TestAddTask(t *testing.T) {
	ctx := context.Background()

	uuid := uuid.NewString()
	userID, err := user.Add(ctx, &model.User{Username: uuid}, nil)
	require.NoError(t, err)

	jID := model.NewJobID()
	jIn := &model.Job{
		JobID:   jID,
		JobType: model.JobTypeExperiment,
		OwnerID: &userID,
		QPos:    decimal.New(0, 0),
	}
	err = db.AddJobTx(ctx, db.Bun(), jIn)
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
	err = db.Bun().NewSelect().Model(&taskCtxDir).Where("task_id = ?", mockT.TaskID).Scan(ctx, &taskCtxDir)
	require.NoError(t, err)
	require.Equal(t, b, taskCtxDir.ContextDirectory)
}

func TestTaskCompleted(t *testing.T) {
	ctx := context.Background()
	mockT := mockTask(t)

	completed, err := TaskCompleted(ctx, mockT.TaskID)
	require.False(t, completed)
	require.NoError(t, err)

	err = CompleteTask(mockT.TaskID, time.Now().UTC().Truncate(time.Millisecond))
	require.NoError(t, err)

	completed, err = TaskCompleted(ctx, mockT.TaskID)
	require.True(t, completed)
	require.NoError(t, err)
}

func TestAddAllocation(t *testing.T) {
	mockT := mockTask(t)

	a := model.Allocation{
		AllocationID: model.AllocationID(fmt.Sprintf("%s-1", mockT.TaskID)),
		TaskID:       mockT.TaskID,
		StartTime:    ptrs.Ptr(time.Now().UTC()),
		State:        ptrs.Ptr(model.AllocationStateTerminated),
	}
	err := AddAllocation(&a)
	require.NoError(t, err, "failed to add allocation")

	logrus.Debugf("%s", db.Bun().NewSelect().Table("allocations"))

	var tmp model.Allocation
	err = db.Bun().NewSelect().Table("allocations").Where("allocation_id = ?", string(a.AllocationID)).
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
	mockT := mockTask(t)
	mockA := mockAllocation(t, mockT.TaskID)

	mockA.EndTime = ptrs.Ptr(time.Now().UTC())

	err := CompleteAllocation(mockA)
	require.NoError(t, err)

	alloc := getAllocationByID(t, mockA.AllocationID)
	require.Equal(t, mockA.EndTime, alloc.EndTime)
}

func TestCloseOpenAllocations(t *testing.T) {
	mockT := mockTask(t)
	mockA1 := mockAllocation(t, mockT.TaskID)
	mockA2 := mockAllocation(t, mockT.TaskID)

	err := CloseOpenAllocations([]model.AllocationID{mockA1.AllocationID})
	require.NoError(t, err)

	alloc1 := getAllocationByID(t, mockA1.AllocationID)
	require.Nil(t, alloc1.EndTime)

	alloc2 := getAllocationByID(t, mockA2.AllocationID)
	logrus.Debugf("ALLOCATION 2: %v", alloc2)
	require.NotNil(t, alloc2.EndTime)

	err = CloseOpenAllocations([]model.AllocationID{})
	require.NoError(t, err)

	alloc1 = getAllocationByID(t, mockA1.AllocationID)
	require.NotNil(t, alloc1.EndTime)

	alloc2 = getAllocationByID(t, mockA2.AllocationID)
	require.NotNil(t, alloc2.EndTime)
}

func TestTaskLogs(t *testing.T) {
}

func TestAddTaskLogs(t *testing.T) {
}

func TestDeleteTaskLogs(t *testing.T) {
}

func TestTaskLogCounts(t *testing.T) {
}

func TestRecordTaskEndStatsBun(t *testing.T) {
}

func TestEndAllTaskStats(t *testing.T) {
}

func TestTaskLogsFields(t *testing.T) {
}

// mockTask returns a mock task.
func mockTask(t *testing.T) *model.Task {
	ctx := context.Background()

	uuid := uuid.NewString()
	userID, err := user.Add(ctx, &model.User{Username: uuid}, nil)
	require.NoError(t, err)

	jID := model.NewJobID()
	jIn := &model.Job{
		JobID:   jID,
		JobType: model.JobTypeExperiment,
		OwnerID: &userID,
		QPos:    decimal.New(0, 0),
	}
	err = db.AddJobTx(ctx, db.Bun(), jIn)
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
	a := model.Allocation{
		AllocationID: model.AllocationID(fmt.Sprintf("%s-1", tID)),
		TaskID:       tID,
		StartTime:    ptrs.Ptr(time.Now().UTC()),
		State:        ptrs.Ptr(model.AllocationStateTerminated),
	}
	err := AddAllocation(&a)
	require.NoError(t, err, "failed to add allocation")
	return &a
}

func getAllocationByID(t *testing.T, aID model.AllocationID) *model.Allocation {
	var alloc model.Allocation
	err := db.Bun().NewSelect().Table("allocations").
		Where("allocation_id = ?", aID).Scan(context.Background(), &alloc)
	require.NoError(t, err)
	return &alloc
}
