package agentrm

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"

	"github.com/determined-ai/determined/master/internal/config"
	"github.com/determined-ai/determined/master/internal/connsave"
	"github.com/determined-ai/determined/master/internal/sproto"
	"github.com/determined-ai/determined/master/pkg/actor"
	"github.com/determined-ai/determined/master/pkg/actor/api"
	"github.com/determined-ai/determined/master/pkg/aproto"
	"github.com/determined-ai/determined/master/pkg/check"
	"github.com/determined-ai/determined/master/pkg/model"
	"github.com/determined-ai/determined/master/pkg/syncx/queue"
	"github.com/determined-ai/determined/proto/pkg/apiv1"
)

type agentService interface {
	list() map[agentID]*agentState
}

type actorAgentService struct {
	syslog    *logrus.Entry
	agentsRef *actor.Ref
	system    *actor.System
}

func newActorAgentService(system *actor.System, agentsRef *actor.Ref) *actorAgentService {
	return &actorAgentService{
		syslog:    logrus.WithField("component", "agents"),
		agentsRef: agentsRef,
		system:    system,
	}
}

// list implements agentService.
func (a *actorAgentService) list() map[agentID]*agentState {
	agents := a.agentsRef.Children()

	responses := a.system.AskAll(getAgentState{}, agents...).GetAll()

	result := make(map[agentID]*agentState, len(agents))
	for ref, msg := range responses {
		switch msg := msg.(type) {
		case *agentState:
			result[msg.agentID()] = msg
		case error:
			a.syslog.WithField("component", "agents").WithError(msg).
				Warnf("failed to get agent state for agent %s", ref.Address().Local())
		default:
			a.syslog.Warnf("bad agent state response for agent %s", ref.Address().Local())
		}
	}

	return result
}

type agentUpdatedEvent struct {
	resourcePool string
}

// initializeAgents creates a new global agents actor.
func initializeAgents(
	system *actor.System, poolConfigs []config.ResourcePoolConfig, e *echo.Echo, opts *aproto.MasterSetAgentOptions,
) (agentService, *queue.Queue[agentUpdatedEvent]) {
	updates := queue.New[agentUpdatedEvent]()
	agentsRef, ok := system.ActorOf(sproto.AgentsAddr, &agents{
		updates:     updates,
		poolConfigs: poolConfigs,
		opts:        opts,
	})
	check.Panic(check.True(ok, "agents address already taken"))
	system.Ask(agentsRef, actor.Ping{}).Get()
	e.GET("/agent*", func(c echo.Context) error {
		if c.IsWebSocket() {
			handler := api.Route(system, nil)
			if err := handler(c); err != nil {
				return err
			}
			return nil
		}
		return echo.ErrNotFound
	})
	return newActorAgentService(system, agentsRef), updates
}

type agents struct {
	updates     *queue.Queue[agentUpdatedEvent]
	poolConfigs []config.ResourcePoolConfig
	opts        *aproto.MasterSetAgentOptions
}

func (a *agents) Receive(ctx *actor.Context) error {
	switch msg := ctx.Message().(type) {
	case actor.PreStart:
		// TODO(ilia): only restore the agents which have some non-zero state.
		// Currently, if an agent tries to reconnect and it was not restored here,
		// then it'd be told it must restart and do a fresh connection.
		agentStates, err := retrieveAgentStates()
		if err != nil {
			ctx.Log().WithError(err).Warnf("failed to retrieve agent states")
		}

		ctx.Log().Debugf("agent states to restore: %d", len(agentStates))
		badAgentIds := []agentID{}

		for agentID := range agentStates {
			state := agentStates[agentID]
			agentRef, err := a.createAgentActor(
				ctx, agentID, state.resourcePoolName, a.opts, &state)
			if err != nil {
				ctx.Log().WithError(err).Warnf("failed to create agent %s", agentID)
				badAgentIds = append(badAgentIds, agentID)
				continue
			}
			ctx.Ask(agentRef, actor.Ping{}).Get()
			ctx.Log().Debugf("restored agent state: %s", agentID)
		}

		if len(badAgentIds) > 0 {
			ctx.Log().Debugf("cleaning %d bad agent states", len(badAgentIds))
			if err := clearAgentStates(badAgentIds); err != nil {
				ctx.Log().WithError(err).Warnf("failed to clean bad agent states")
			}
		}
	case api.WebSocketRequest:
		cmuxConn := connsave.GetConn(msg.Ctx.Request().Context()).(*cmux.MuxConn)
		// Here, we just have to check that there are any certificates at all, since the top-level TLS
		// config verifies that any certificates that are provided are valid.
		if tlsConn, ok := cmuxConn.Conn.(*tls.Conn); ok {
			requireAuth := config.GetMasterConfig().ResourceManager.AgentRM.RequireAuthentication
			missingAuth := len(tlsConn.ConnectionState().PeerCertificates) == 0
			if requireAuth && missingAuth {
				ctx.Log().WithField("remote-addr", tlsConn.RemoteAddr()).
					Warnf("rejecting agent WebSocket request with no certificates")
				ctx.Respond(echo.ErrForbidden)
				return nil
			}
		}

		id := msg.Ctx.QueryParam("id")
		existingRef := ctx.Child(id)
		reconnect, err := msg.IsReconnect()
		if err != nil {
			ctx.Respond(errors.Wrapf(err, "parsing reconnect query param"))
			return nil
		}

		// If the agent actor is still alive on our side when an agent tries to reconnect,
		// accept it. Whether it is a network failure or a crash/restart, we will just try
		// to reattach whatever containers still exist.
		// That logic is located in agent.receive(ws.WebSocketRequest).
		if existingRef != nil {
			ctx.Log().WithField("reconnect", reconnect).Infof("restoring agent id: %s", id)
			ctx.Respond(ctx.Ask(existingRef, msg).Get())
			return nil
		}

		// If the agent actor is _not_ alive on our side and the agent is trying to reconnect,
		// continue to deny it. This case is nearly impossible (master waits longer than agent
		// tries, to avoid it).
		if reconnect {
			ctx.Respond(aproto.ErrAgentMustReconnect)
			return nil
		}

		// Finally, this must not be a recovery flow, so just create the agent actor.
		resourcePool := msg.Ctx.QueryParam("resource_pool")
		if ref, err := a.createAgentActor(ctx, agentID(id), resourcePool, a.opts, nil); err != nil {
			ctx.Respond(err)
		} else {
			ctx.Respond(ctx.Ask(ref, msg).Get())
		}

	case *apiv1.GetAgentsRequest:
		response := &apiv1.GetAgentsResponse{}
		for _, a := range a.summarize(ctx) {
			response.Agents = append(response.Agents, a.ToProto())
		}
		ctx.Respond(response)
	case echo.Context:
		a.handleAPIRequest(ctx, msg)
	case actor.PostStop:
	default:
		return actor.ErrUnexpectedMessage(ctx)
	}
	return nil
}

func (a *agents) createAgentActor(
	ctx *actor.Context,
	id agentID,
	resourcePool string,
	opts *aproto.MasterSetAgentOptions,
	restoredAgentState *agentState,
) (*actor.Ref, error) {
	if id == "" {
		return nil, errors.Errorf("invalid agent id specified: %s", id)
	}
	if resourcePool == "" {
		ctx.Log().Info("resource pool is empty; using default resource pool: default")
		resourcePool = "default"
	}

	var poolConfig *config.ResourcePoolConfig
	for _, pc := range a.poolConfigs {
		pc := pc
		if pc.PoolName == resourcePool {
			poolConfig = &pc
			break
		}
	}
	if poolConfig == nil {
		return nil, fmt.Errorf("cannot find specified resource pool %s for agent %s", resourcePool, id)
	}

	ref, ok := ctx.ActorOf(id, &agent{
		updates:               a.updates,
		resourcePoolName:      resourcePool,
		maxZeroSlotContainers: poolConfig.MaxAuxContainersPerAgent,
		agentReconnectWait:    time.Duration(poolConfig.AgentReconnectWait),
		opts:                  opts,
		agentState:            restoredAgentState,
	})
	if !ok {
		return nil, errors.Errorf("agent already connected: %s", id)
	}
	return ref, nil
}

func (a *agents) handleAPIRequest(ctx *actor.Context, apiCtx echo.Context) {
	switch apiCtx.Request().Method {
	case echo.GET:
		ctx.Respond(apiCtx.JSON(http.StatusOK, a.summarize(ctx)))
	default:
		ctx.Respond(echo.ErrMethodNotAllowed)
	}
}

func (a *agents) summarize(ctx *actor.Context) model.AgentsSummary {
	results := ctx.AskAll(model.AgentSummary{}, ctx.Children()...).GetAll()
	summary := make(map[string]model.AgentSummary, len(results))
	for ref, result := range results {
		summary[ref.Address().String()] = result.(model.AgentSummary)
	}
	return summary
}
