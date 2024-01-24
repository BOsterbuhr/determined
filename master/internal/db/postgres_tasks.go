package db

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/determined-ai/determined/master/internal/api"
	"github.com/determined-ai/determined/master/pkg/model"
	"github.com/determined-ai/determined/proto/pkg/apiv1"
)

func completeTrialsTasks(ex sqlx.Execer, trialID int, endTime time.Time) error {
	if _, err := ex.Exec(`
UPDATE tasks
SET end_time = $2
FROM trial_id_task_id
WHERE trial_id_task_id.task_id = tasks.task_id
  AND trial_id_task_id.trial_id = $1
  AND end_time IS NULL`, trialID, endTime); err != nil {
		return fmt.Errorf("completing task: %w", err)
	}
	return nil
}

// taskLogsFieldMap is used to map fields in filters to expressions. This was used historically
// in trial logs to either read timestamps or regex them out of logs.
var taskLogsFieldMap = map[string]string{}

type taskLogsFollowState struct {
	// The last ID returned by the query. Historically the trial logs API when streaming
	// repeatedly made a request like SELECT ... FROM trial_logs ... ORDER BY k OFFSET N LIMIT M.
	// Since offset is less than optimal (no filtering is done during the initial
	// index scan), we at least pass Postgres the ID and let it begin after a certain ID rather
	// than offset N into the query.
	id int64
}

// TaskLogs takes a task ID and log offset, limit and filters and returns matching logs.
func (db *PgDB) TaskLogs(
	taskID model.TaskID, limit int, fs []api.Filter, order apiv1.OrderBy, followState interface{},
) ([]*model.TaskLog, interface{}, error) {
	if followState != nil {
		fs = append(fs, api.Filter{
			Field:     "id",
			Operation: api.FilterOperationGreaterThan,
			Values:    []int64{followState.(*taskLogsFollowState).id},
		})
	}

	params := []interface{}{taskID, limit}
	fragment, params := filtersToSQL(fs, params, taskLogsFieldMap)
	query := fmt.Sprintf(`
SELECT
    l.id,
    l.task_id,
    l.allocation_id,
    l.agent_id,
    l.container_id,
    l.rank_id,
    l.timestamp,
    l.level,
    l.stdtype,
    l.source,
    l.log
FROM task_logs l
WHERE l.task_id = $1
%s
ORDER BY l.id %s LIMIT $2
`, fragment, OrderByToSQL(order))

	var b []*model.TaskLog
	if err := db.queryRows(query, &b, params...); err != nil {
		return nil, nil, err
	}

	if len(b) > 0 {
		lastLog := b[len(b)-1]
		followState = &taskLogsFollowState{id: int64(*lastLog.ID)}
	}

	return b, followState, nil
}

// AddTaskLogs adds a list of *model.TaskLog objects to the database with automatic IDs.
func (db *PgDB) AddTaskLogs(logs []*model.TaskLog) error {
	if len(logs) == 0 {
		return nil
	}

	var text strings.Builder
	text.WriteString(`
INSERT INTO task_logs
  (task_id, allocation_id, log, agent_id, container_id, rank_id, timestamp, level, stdtype, source)
VALUES
`)

	args := make([]interface{}, 0, len(logs)*10)

	for i, log := range logs {
		if i > 0 {
			text.WriteString(",")
		}
		// TODO(brad): We can do better.
		fmt.Fprintf(&text, " ($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)",
			i*10+1, i*10+2, i*10+3, i*10+4, i*10+5, i*10+6, i*10+7, i*10+8, i*10+9, i*10+10)

		args = append(args, log.TaskID, log.AllocationID, []byte(log.Log), log.AgentID, log.ContainerID,
			log.RankID, log.Timestamp, log.Level, log.StdType, log.Source)
	}

	if _, err := db.sql.Exec(text.String(), args...); err != nil {
		return errors.Wrapf(err, "error inserting %d task logs", len(logs))
	}

	return nil
}

// DeleteTaskLogs deletes the logs for the given tasks.
func (db *PgDB) DeleteTaskLogs(ids []model.TaskID) error {
	if _, err := db.sql.Exec(`
DELETE FROM task_logs
WHERE task_id IN (SELECT unnest($1::text [])::text);
`, ids); err != nil {
		return errors.Wrapf(err, "error deleting task logs for task %v", ids)
	}
	return nil
}

// TaskLogsCount returns the number of logs in postgres for the given task.
func (db *PgDB) TaskLogsCount(taskID model.TaskID, fs []api.Filter) (int, error) {
	params := []interface{}{taskID}
	fragment, params := filtersToSQL(fs, params, taskLogsFieldMap)
	query := fmt.Sprintf(`
SELECT count(*)
FROM task_logs
WHERE task_id = $1
%s
`, fragment)
	var count int
	if err := db.sql.QueryRow(query, params...).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

// RecordTaskStats record stats for tasks.
func (db *PgDB) RecordTaskStats(stats *model.TaskStats) error {
	return RecordTaskStatsBun(stats)
}

// RecordTaskStatsBun record stats for tasks with bun.
func RecordTaskStatsBun(stats *model.TaskStats) error {
	_, err := Bun().NewInsert().Model(stats).Exec(context.TODO())
	return err
}

// RecordTaskEndStats record end stats for tasks.
func (db *PgDB) RecordTaskEndStats(stats *model.TaskStats) error {
	return RecordTaskEndStatsBun(stats)
}

// RecordTaskEndStatsBun record end stats for tasks with bun.
func RecordTaskEndStatsBun(stats *model.TaskStats) error {
	query := Bun().NewUpdate().Model(stats).Column("end_time").
		Where("allocation_id = ?", stats.AllocationID).
		Where("event_type = ?", stats.EventType).
		Where("end_time IS NULL")
	if stats.ContainerID == nil {
		// Just doing Where("container_id = ?", stats.ContainerID) in the null case
		// generates WHERE container_id = NULL which doesn't seem to match on null rows.
		// We don't use this case anywhere currently but this feels like an easy bug to write
		// without this.
		query = query.Where("container_id IS NULL")
	} else {
		query = query.Where("container_id = ?", stats.ContainerID)
	}

	if _, err := query.Exec(context.TODO()); err != nil {
		return fmt.Errorf("recording task end stats %+v: %w", stats, err)
	}

	return nil
}

// EndAllTaskStats called at master starts, in case master previously crashed.
func (db *PgDB) EndAllTaskStats() error {
	_, err := db.sql.Exec(`
UPDATE task_stats SET end_time = greatest(cluster_heartbeat, task_stats.start_time)
FROM cluster_id, allocations
WHERE allocations.allocation_id = task_stats.allocation_id
AND allocations.end_time IS NOT NULL
AND task_stats.end_time IS NULL`)
	if err != nil {
		return fmt.Errorf("ending all task stats: %w", err)
	}

	return nil
}

// TaskLogsFields returns the unique fields that can be filtered on for the given task.
func (db *PgDB) TaskLogsFields(taskID model.TaskID) (*apiv1.TaskLogsFieldsResponse, error) {
	var fields apiv1.TaskLogsFieldsResponse
	err := db.QueryProto("get_task_logs_fields", &fields, taskID)
	return &fields, err
}

// MaxTerminationDelay is the max delay before a consumer can be sure all logs have been recevied.
// For Postgres, we don't need to wait very long at all; this was a hypothetical cap on fluent
// to DB latency prior to fluent's deprecation.
func (db *PgDB) MaxTerminationDelay() time.Duration {
	// TODO: K8s logs can take a bit to get to us, so much so we should investigate.
	return 5 * time.Second
}
