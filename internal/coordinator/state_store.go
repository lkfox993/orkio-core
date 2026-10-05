package coordinator

import (
	"context"

	"github.com/lkfox993/orkio-core/internal/job"
	"github.com/lkfox993/orkio-core/internal/process"
	"github.com/lkfox993/orkio-core/internal/workflow"
)

type StateStore interface {
	GetJob(ctx context.Context, id string) (*job.Job, error)
	CompleteJob(ctx context.Context, id string) error
	CreateJob(ctx context.Context, job *job.Job) error
	GetProcessInstance(ctx context.Context, id string) (*execution.WorkflowInstance, error)
	CompleteWorkflow(ctx context.Context, id string) error
	GetWorkflow(ctx context.Context, key string) (workflow.Workflow, error)
	CreateProcessInstance(ctx context.Context, instance *process.ProcessInstance) error
}
