package logretention

import (
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/determined-ai/determined/master/internal/db"
	"github.com/determined-ai/determined/master/pkg/model"
)

var log = logrus.WithField("component", "log-retention")

var scheduler gocron.Scheduler

func init() {
	limitOpt := gocron.WithLimitConcurrentJobs(1, gocron.LimitModeReschedule)
	var err error
	scheduler, err = gocron.NewScheduler(limitOpt)
	if err != nil {
		panic(errors.Wrapf(err, "failed to create logretention scheduler"))
	}
}

// Schedule begins a log deletion schedule according to the provided LogRetentionPolicy.
func Schedule(config model.LogRetentionPolicy, db *db.PgDB) error {
	var defaultLogRetentionDays int16 = -1
	if config.Days != nil {
		defaultLogRetentionDays = *config.Days
	}
	// Create a task that deletes expired task logs.
	task := gocron.NewTask(func() {
		log.WithField("default-duration", defaultLogRetentionDays).Trace("deleting expired task logs")
		count, err := db.DeleteExpiredTaskLogs(defaultLogRetentionDays)
		if err != nil {
			log.WithError(err).Error("failed to delete expired task logs")
		} else if count > 0 {
			log.WithField("count", count).Info("deleted expired task logs")
		}
	})
	// If cleanup on start is enabled, run the cleanup task immediately.
	if config.CleanupOnStart != nil && *config.CleanupOnStart {
		log.Debug("running task log cleanup on start")
		_, err := scheduler.NewJob(gocron.OneTimeJob(gocron.OneTimeJobStartImmediately()), task)
		if err != nil {
			return errors.Wrapf(err, "failed to schedule startup task log cleanup")
		}
	}
	// If a cleanup schedule is set, schedule the cleanup task.
	if config.Schedule != nil {
		if d, err := time.ParseDuration(*config.Schedule); err == nil {
			// Try to parse out a duration.
			log.WithField("duration", d).Debug("running task log cleanup with duration")
			_, err := scheduler.NewJob(gocron.DurationJob(d), task)
			if err != nil {
				return errors.Wrapf(err, "failed to schedule duration task log cleanup")
			}
		} else {
			// Otherwise, use a cron.
			log.WithField("cron", *config.Schedule).Debug("running task log cleanup with cron")
			_, err := scheduler.NewJob(gocron.CronJob(*config.Schedule, false), task)
			if err != nil {
				return errors.Wrapf(err, "failed to schedule cron task log cleanup")
			}
		}
	}
	// Start the scheduler.
	scheduler.Start()
	return nil
}
