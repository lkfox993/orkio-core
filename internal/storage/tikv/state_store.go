package tikv

import (
	"context"

	"github.com/lkfox993/orkio-core/internal/job"
	"github.com/lkfox993/orkio-core/internal/process"
)

type StateStore struct {
	*ProcessStateStore
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
