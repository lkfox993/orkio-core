package coordinator

import (
	"context"
	"fmt"

	"github.com/lkfox993/orkio-core/internal/events"
	"github.com/lkfox993/orkio-core/internal/job"
)

type Engine struct {
	state StateStore
}

func NewEngine(state StateStore) *Engine {
	return &Engine{
		state: state,
	}
}

func (e *Engine) HandleJobCompleted(ctx context.Context, event events.JobCompleted) error {

	j, err := e.state.GetJob(ctx, event.JobID)

	if err != nil {
		return fmt.Errorf("get job: %w", err)
	}

	// idempotency
	if j.Status == job.StatusCompleted {
		return nil
	}

	if err := e.state.CompleteJob(ctx, j.ID); err != nil {
		return fmt.Errorf("complete job: %w", err)
	}

	workflow, err := e.state.GetWorkflowInstance(ctx, j.WorkflowInstanceID)

	if err != nil {
		return fmt.Errorf("get workflow instance: %w", err)
	}

	nextStep := workflow.NextStep(j.StepID)

	if nextStep == nil {

		if err := e.state.CompleteWorkflow(ctx, workflow.ID); err != nil {
			return fmt.Errorf("complete workflow: %w", err)
		}

		return nil
	}

	return nil
}

func (e *Engine) HandleJobFailed(ctx context.Context, event events.JobFailed) error {

	j, err := e.state.GetJob(ctx, event.JobID)

	if err != nil {
		return fmt.Errorf("get job: %w", err)
	}

	if j.Status == job.StatusFailed {
		return nil
	}

	if j.Attempt < j.MaxAttempts {
		return e.retryJob(ctx, j)
	}

	if err := e.state.FailJob(ctx, job.ID, err); err != nil {
		return fmt.Errorf("fail job: %w", err)
	}

	if err := e.state.FailWorkflow(
		ctx,
		job.WorkflowInstanceID,
	); err != nil {
		return fmt.Errorf("fail workflow: %w", err)
	}

	return nil
}

func (e *Engine) HandleJobTimedOut(ctx context.Context, event events.JobTimedOut) error {

	j, err := e.state.GetJob(ctx, event.JobID)

	if err != nil {
		return fmt.Errorf("get job: %w", err)
	}

	if j.Status == job.StatusRunning {
		return nil
	}

	if j.Attempt < j.MaxAttempts {
		return e.retryJob(ctx, j)
	}

	if err := e.state.FailJob(ctx, j.ID, "job timeout"); err != nil {
		return fmt.Errorf("fail job: %w", err)
	}

	if err := e.state.FailWorkflow(ctx, j.WorkflowInstanceID); err != nil {
		return fmt.Errorf("fail workflow: %w", err)
	}

	return nil
}
