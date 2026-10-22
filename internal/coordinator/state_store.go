package coordinator

import (
	"context"

	"github.com/google/uuid"
	"github.com/lkfox993/orkio-core/internal/job"
	"github.com/lkfox993/orkio-core/internal/process"
)

type StateStore interface {
	GetJob(ctx context.Context, id string) (*job.Job, error)
	CompleteJob(ctx context.Context, id string) error
	CreateJob(ctx context.Context, job *job.Job) error
	GetProcess(ctx context.Context, id string) (*process.Process, error)
	CompleteProcess(ctx context.Context, id string) error
	CreateProcess(ctx context.Context, instance *process.Process) error
	CreateNodeExecution(
		ctx context.Context,
		execution *process.NodeExecution,
	) error

	CompleteNodeExecution(
		ctx context.Context,
		executionID uuid.UUID,
	) error
}
