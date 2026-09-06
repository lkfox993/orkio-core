package coordinator

import (
	"context"

	"github.com/lkfox993/orkio-core/internal/job"
)

type StateStore interface {
	GetJob(ctx context.Context, id string) (*job.Job, error)
	CompleteJob(ctx context.Context, id string) error
	CreateJob(ctx context.Context, job job.Job) error
	GetWorkflowInstance(ctx context.Context, id string) (*execution.WorkflowInstance, error)
	CompleteWorkflow(ctx context.Context, id string) error
}
