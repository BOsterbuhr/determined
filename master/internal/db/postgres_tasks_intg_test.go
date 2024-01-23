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
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/determined-ai/determined/master/pkg/cproto"
	"github.com/determined-ai/determined/master/pkg/etc"
	"github.com/determined-ai/determined/master/pkg/model"
	"github.com/determined-ai/determined/master/pkg/ptrs"
)

func TestRecordAndEndTaskStats(t *testing.T) {
	require.NoError(t, etc.SetRootPath(RootFromDB))
	db := MustResolveTestPostgres(t)
	MustMigrateTestPostgres(t, db, MigrationsFromDB)

	tID := model.NewTaskID()
	require.NoError(t, AddTask(context.Background(), &model.Task{
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

}

func TestTaskByID(t *testing.T) {

}

func TestAddNonExperimentTasksContextDirectory(t *testing.T) {

}

func TestTaskCompleted(t *testing.T) {

}

func TestCompleteTask(t *testing.T) {

}

func TestAddAllocation(t *testing.T) {

}

func TestAddAllocationExitStatus(t *testing.T) {

}

func TestCompleteAllocation(t *testing.T) {

}

func TestCompleteAllocationTelemetry(t *testing.T) {

}

func TestCloseOpenAllocations(t *testing.T) {

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

// RequireMockTask returns a mock task.
func RequireMockTask(t *testing.T, db *PgDB, userID *model.UserID) *model.Task {
	jID := RequireMockJob(t, db, userID)

	// Add a task.
	tID := model.NewTaskID()
	tIn := &model.Task{
		TaskID:    tID,
		JobID:     &jID,
		TaskType:  model.TaskTypeTrial,
		StartTime: time.Now().UTC().Truncate(time.Millisecond),
	}
	err := AddTask(context.Background(), tIn)
	require.NoError(t, err, "failed to add task")
	return tIn
}
