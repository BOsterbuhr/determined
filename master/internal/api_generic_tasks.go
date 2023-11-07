package internal

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/uptrace/bun"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	k8sV1 "k8s.io/api/core/v1"

	"github.com/ghodss/yaml"

	"github.com/determined-ai/determined/master/internal/config"
	"github.com/determined-ai/determined/master/internal/db"
	"github.com/determined-ai/determined/master/internal/grpcutil"
	"github.com/determined-ai/determined/master/internal/sproto"
	"github.com/determined-ai/determined/master/internal/user"
	"github.com/determined-ai/determined/master/pkg/archive"
	"github.com/determined-ai/determined/master/pkg/check"
	pkgCommand "github.com/determined-ai/determined/master/pkg/command"
	"github.com/determined-ai/determined/master/pkg/model"
	"github.com/determined-ai/determined/master/pkg/ptrs"
	"github.com/determined-ai/determined/master/pkg/schemas"
	"github.com/determined-ai/determined/master/pkg/schemas/expconf"
	"github.com/determined-ai/determined/master/pkg/tasks"
	"github.com/determined-ai/determined/proto/pkg/apiv1"
	"github.com/determined-ai/determined/proto/pkg/utilv1"
)

func parseJustGenericResources(configBytes []byte) model.TaskResourcesConfig {
	// Make this function usable on experiment or command configs.
	type DummyConfig struct {
		Resources model.TaskResourcesConfig `json:"resources"`
	}

	dummy := DummyConfig{
		Resources: model.TaskResourcesConfig{
			SlotsPerTask: 1,
		},
	}

	// Don't throw errors; validation should happen elsewhere.
	_ = yaml.Unmarshal(configBytes, &dummy)

	return dummy.Resources
}

func (a *apiServer) getGenericTaskLaunchParameters(
	ctx context.Context,
	contextDirectory []*utilv1.File,
	configYAML string,
	projectID *int,
) (
	*tasks.GenericTaskSpec, []pkgCommand.LaunchWarning, error,
) {
	defer func() {
		debug.PrintStack()
	}()

	fmt.Println("Getting to master", configYAML)
	var err error
	genericTaskSpec := &tasks.GenericTaskSpec{}

	genericTaskSpec.ProjectID = model.DefaultProjectID
	if projectID != nil {
		genericTaskSpec.ProjectID = *projectID
	}

	// Validate the userModel and get the agent userModel group.
	userModel, _, err := grpcutil.GetUser(ctx)
	if err != nil {
		return nil,
			nil,
			status.Errorf(codes.Unauthenticated, "failed to get the user: %s", err)
	}

	workspaceID := 1 // TODO convert projectID to workspaceID here
	agentUserGroup, err := user.GetAgentUserGroup(ctx, userModel.ID, workspaceID)
	if err != nil {
		return nil, nil, err
	}

	// Validate the resource configuration.
	resources := parseJustGenericResources([]byte(configYAML))

	poolName, err := a.m.rm.ResolveResourcePool(
		resources.ResourcePool, workspaceID, resources.SlotsPerTask)
	if err != nil {
		return nil, nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	launchWarnings, err := a.m.rm.ValidateResourcePoolAvailability(
		&sproto.ValidateResourcePoolAvailabilityRequest{
			Name:  poolName,
			Slots: resources.SlotsPerTask,
		},
	)
	if err != nil {
		return nil, nil, fmt.Errorf("checking resource availability: %v", err.Error())
	}
	if a.m.config.ResourceManager.AgentRM != nil &&
		a.m.config.LaunchError &&
		len(launchWarnings) > 0 {
		return nil, nil, fmt.Errorf("slots requested exceeds cluster capacity")
	}

	// Get the base TaskSpec.
	taskContainerDefaults, err := a.m.rm.TaskContainerDefaults(
		poolName,
		a.m.config.TaskContainerDefaults,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("getting TaskContainerDefaults: %v", err)
	}
	taskSpec := *a.m.taskSpec
	taskSpec.TaskContainerDefaults = taskContainerDefaults
	taskSpec.AgentUserGroup = agentUserGroup
	taskSpec.Owner = userModel

	// Get the full configuration.
	taskConfig := model.DefaultConfigGenericTaskConfig(&taskSpec.TaskContainerDefaults)
	workDirInDefaults := taskConfig.WorkDir

	// TODO don't remove defaults with this.
	if err := yaml.Unmarshal([]byte(configYAML), &taskConfig); err != nil {
		return nil, nil, fmt.Errorf("yaml unmarshaling generic task config: %w", err)
	}

	// Copy discovered (default) resource pool name and slot count.
	taskConfig.Resources.ResourcePool = poolName
	taskConfig.Resources.SlotsPerTask = resources.SlotsPerTask

	taskContainerPodSpec := taskSpec.TaskContainerDefaults.GPUPodSpec
	if taskConfig.Resources.SlotsPerTask == 0 {
		taskContainerPodSpec = taskSpec.TaskContainerDefaults.CPUPodSpec
	}
	taskConfig.Environment.PodSpec = (*k8sV1.Pod)(schemas.Merge(
		(*expconf.PodSpec)(taskConfig.Environment.PodSpec),
		(*expconf.PodSpec)(taskContainerPodSpec),
	))

	var contextDirectoryBytes []byte
	if len(contextDirectory) > 0 {
		userFiles := filesToArchive(contextDirectory)

		workdirSetInReq := taskConfig.WorkDir != nil &&
			(workDirInDefaults == nil || *workDirInDefaults != *taskConfig.WorkDir)
		if workdirSetInReq {
			return nil, nil, status.Errorf(codes.InvalidArgument,
				"cannot set work_dir and context directory at the same time")
		}
		taskConfig.WorkDir = nil

		contextDirectoryBytes, err = archive.ToTarGz(userFiles)
		if err != nil {
			return nil, nil, status.Errorf(codes.InvalidArgument,
				fmt.Errorf("compressing files context files: %w", err).Error())
		}
	}

	extConfig := config.GetMasterConfig().InternalConfig.ExternalSessions
	var token string
	if extConfig.Enabled() {
		token, err = grpcutil.GetUserExternalToken(ctx)
		if err != nil {
			return nil, nil, status.Errorf(codes.Internal,
				errors.Wrapf(err,
					"unable to get external user token").Error())
		}
		err = nil
	} else {
		token, err = user.StartSession(ctx, userModel)
		if err != nil {
			return nil, nil, status.Errorf(codes.Internal,
				errors.Wrapf(err,
					"unable to create user session inside task").Error())
		}
	}
	taskSpec.UserSessionToken = token

	genericTaskSpec.Base = taskSpec
	genericTaskSpec.GenericTaskConfig = taskConfig

	taskID := model.NewTaskID()
	jobID := model.NewJobID()
	err = db.Bun().RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		// TODO these actually aren't in the transcation.
		if err := a.m.db.AddJob(&model.Job{
			JobID:   jobID,
			JobType: model.JobTypeGeneric,
			OwnerID: &genericTaskSpec.Base.Owner.ID,
		}); err != nil {
			return errors.Wrapf(err, "persisting job %v", taskID)
		}

		if err := a.m.db.AddTask(&model.Task{
			TaskID:     taskID,
			TaskType:   model.TaskTypeGeneric,
			StartTime:  time.Now(), // start time is submit time?
			JobID:      &jobID,
			LogVersion: model.CurrentTaskLogVersion,
		}); err != nil {
			return errors.Wrapf(err, "persisting task %v", taskID)
		}

		// TODO persist config elemnts

		if _, err := tx.NewInsert().Model(&model.TaskContextDirectory{
			TaskID:           taskID,
			ContextDirectory: contextDirectoryBytes,
		}).Exec(context.TODO()); err != nil {
			return fmt.Errorf("persisting context directory files: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, nil, fmt.Errorf("persisting task information: %w", err)
	}

	return genericTaskSpec, launchWarnings, nil
}

func (a *apiServer) CreateGenericTask(
	ctx context.Context, req *apiv1.CreateGenericTaskRequest,
) (*apiv1.CreateGenericTaskResponse, error) {
	// Parse launch commnads.
	var projectID *int
	if req.ProjectId != nil {
		projectID = ptrs.Ptr(int(*req.ProjectId))
	}
	genericTaskSpec, warnings, err := a.getGenericTaskLaunchParameters(
		ctx, req.ContextDirectory, req.Config, projectID,
	)
	if err != nil {
		return nil, err
	}
	if len(warnings) > 0 {
		return nil, nil // TODO warnings
	}
	// TODO rbac check.

	// Maybe fill in description?

	// TODO do we need to wrap entrypoint with a custom wrapper?
	// If we do it feels like a weird place to do it

	if err := check.Validate(genericTaskSpec.GenericTaskConfig); err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"invalid generic task config: %s",
			err.Error(),
		)
	}

	genericTaskSpec.Base.ExtraEnvVars = map[string]string{
		"DET_TASK_TYPE": string(model.TaskTypeGeneric),
	}

	// Persist the task.

	// TODO actually create the task

	return &apiv1.CreateGenericTaskResponse{}, nil
}
