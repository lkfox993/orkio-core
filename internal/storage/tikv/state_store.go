package tikv

import (
	"context"

	"github.com/google/uuid"
	"github.com/lkfox993/orkio-core/internal/job"
	"github.com/lkfox993/orkio-core/internal/process"
)

type StateStore struct {
	*ProcessStateStore
	*NodeExecutionStateStore
}

func NewStateStore(stateDB *Database) *StateStore {
	return &StateStore{
		ProcessStateStore: NewProcessStateStore(stateDB),
	}
}

func (s *StateStore) GetJob(ctx context.Context, id string) (*job.Job, error) {
	return nil, nil
}

func (s *StateStore) CreateJob(ctx context.Context, job *job.Job) error {
	return nil
}

func (s *StateStore) CompleteJob(ctx context.Context, id string) error {
	return nil
}

func (s *StateStore) GetProcess(ctx context.Context, id string) (*process.Process, error) {
	return nil, nil
}

func (s *StateStore) CreateProcess(ctx context.Context, p *process.Process) error {
	return s.ProcessStateStore.Create(ctx, p)
}

func (s *StateStore) CompleteProcess(ctx context.Context, id string) error {
	return nil
}

func (s *StateStore) CreateNodeExecution(
	ctx context.Context,
	execution *process.NodeExecution,
) error {
	return s.CreateNodeExecution(ctx, execution)
}

func (s *StateStore) CompleteNodeExecution(
	ctx context.Context,
	executionID uuid.UUID,
) error {
	return s.CompleteNodeExecution(ctx, executionID)
}
