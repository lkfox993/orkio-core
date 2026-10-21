package coordinator

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/lkfox993/orkio-core/internal/events"
	"github.com/lkfox993/orkio-core/internal/job"
	"github.com/lkfox993/orkio-core/internal/process"
	"github.com/lkfox993/orkio-core/internal/storage/postgres"
	"github.com/lkfox993/orkio-core/internal/workflow"
	"github.com/lkfox993/orkio-core/internal/workflow/nodes"
)

const TenantID = "592cf5ef-f12e-4ca2-b66b-27f5f42ac07c"

type CreateProcessInput struct {
	WorkflowID string
	Version    *int
	Variables  map[string]any
}

type Engine struct {
	db    *postgres.Database
	state StateStore
}

func NewEngine(state StateStore, db *postgres.Database) *Engine {
	return &Engine{
		db:    db,
		state: state,
	}
}

func (e *Engine) CreateProcess(
	ctx context.Context,
	workflowKey string,
	variables map[string]any,
) (*process.Process, error) {

	tenantID, err := uuid.Parse(TenantID)

	if err != nil {
		return nil, err
	}

	workflowRepository := workflow.NewWorkflowRepository(e.db)

	workflow, err := workflowRepository.GetByKey(ctx, tenantID, workflowKey)

	if err != nil {
		return nil, fmt.Errorf("get workflow: %w", err)
	}

	p := &process.Process{
		ID:               uuid.New(),
		WorkflowID:       workflow.ID,
		InitialVariables: variables,
		Status:           process.StatusPending,
	}

	processRepository := process.NewProcessRepository(e.db)

	if err := processRepository.Create(ctx, p, nil); err != nil {
		return nil, err
	}

	if err := e.state.CreateProcess(ctx, p); err != nil {
		return nil, fmt.Errorf("create process instance: %w", err)
	}

	return nil, nil

	// // найти первый executable node
	// startNodes := workflow.Definiton.StartNodes()

	// if len(startNodes) == 0 {
	// 	return nil, fmt.Errorf("workflow has no start node")
	// }

	// // тут надо подумать так как бывают start_node from message
	// for _, startNode := range startNodes {
	// 	if err := e.activateNode(ctx, process, startNode); err != nil {
	// 		return nil, err
	// 	}
	// }

	// return process, nil
}

func (e *Engine) activateNode(
	ctx context.Context,
	instance *process.Process,
	node nodes.Node,
) error {
	// switch n := node.(type) {

	// case *nodes.StartEvent:
	// 	return e.activateNode(ctx, instance, n.Next())

	// case *nodes.ServiceTask:
	// 	return e.createJob(ctx, instance, n)

	// case *nodes.EndEvent:
	// 	return e.state.CompleteProcess(ctx, instance.ID)

	// default:
	// 	return fmt.Errorf("unsupported node type: %T", node)
	// }

	return nil
}

func (e *Engine) createJob(
	ctx context.Context,
	process *process.Process,
	node nodes.Node,
) error {

	j := &job.Job{
		ID:        uuid.New(),
		ProcessID: process.ID,
		StepID:    node.Base().ID,
		Type:      node.NodeType(),
		Status:    job.StatusPending,
		Attempt:   0,
	}

	switch n := node.(type) {
	case *nodes.ServiceTask:
		j.MaxAttempts = n.Retry.Attempts

	case *nodes.UserTask:
		j.MaxAttempts = n.Retry.Attempts

	default:
		return fmt.Errorf("node %q cannot create job", node.Base().ID)
	}

	if err := e.state.CreateJob(ctx, j); err != nil {
		return fmt.Errorf("create job: %w", err)
	}

	// publish ExecuteJob в Pulsar

	return nil

}

func (e *Engine) HandleJobCompleted(ctx context.Context, event events.JobCompleted) error {

	// j, err := e.state.GetJob(ctx, event.JobID)

	// if err != nil {
	// 	return fmt.Errorf("get job: %w", err)
	// }

	// // idempotency
	// if j.Status == job.StatusCompleted {
	// 	return nil
	// }

	// if err := e.state.CompleteJob(ctx, j.ID); err != nil {
	// 	return fmt.Errorf("complete job: %w", err)
	// }

	// process, err := e.state.GetProcess(ctx, j.ProcessID) //FIXME - GetWorkflowInstance лучьше переименовать в GetProcessInstance

	// if err != nil {
	// 	return fmt.Errorf("get process instance: %w", err)
	// }

	// nextStep := process.NextStep(j.StepID)

	// if nextStep == nil {

	// 	if err := e.state.CompleteProcess(ctx, process.ID); err != nil {
	// 		return fmt.Errorf("complete workflow: %w", err)
	// 	}

	// 	return nil
	// }

	return nil
}

func (e *Engine) HandleJobFailed(ctx context.Context, event events.JobFailed) error {

	// j, err := e.state.GetJob(ctx, event.JobID)

	// if err != nil {
	// 	return fmt.Errorf("get job: %w", err)
	// }

	// if j.Status == job.StatusFailed {
	// 	return nil
	// }

	// if j.Attempt < j.MaxAttempts {
	// 	return e.retryJob(ctx, j)
	// }

	// if err := e.state.FailJob(ctx, job.ID, err); err != nil {
	// 	return fmt.Errorf("fail job: %w", err)
	// }

	// if err := e.state.FailWorkflow(
	// 	ctx,
	// 	j.ProcessID,
	// ); err != nil {
	// 	return fmt.Errorf("fail process: %w", err)
	// }

	return nil
}

func (e *Engine) HandleJobTimedOut(ctx context.Context, event events.JobTimedOut) error {

	// j, err := e.state.GetJob(ctx, event.JobID)

	// if err != nil {
	// 	return fmt.Errorf("get job: %w", err)
	// }

	// if j.Status == job.StatusRunning {
	// 	return nil
	// }

	// if j.Attempt < j.MaxAttempts {
	// 	return e.retryJob(ctx, j)
	// }

	// if err := e.state.FailJob(ctx, j.ID, "job timeout"); err != nil {
	// 	return fmt.Errorf("fail job: %w", err)
	// }

	// if err := e.state.FailWorkflow(ctx, j.WorkflowInstanceID); err != nil {
	// 	return fmt.Errorf("fail workflow: %w", err)
	// }

	return nil
}
