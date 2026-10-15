package tikv

import (
	"context"

	"github.com/lkfox993/orkio-core/internal/job"
	"github.com/lkfox993/orkio-core/internal/process"
)

type StateStore struct {
}

func NewStateStore() *StateStore {
	return &StateStore{}
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

func (s *StateStore) CreateProcess(ctx context.Context, instance *process.Process) error {
	return nil
}

func (s *StateStore) CompleteProcess(ctx context.Context, id string) error {
	return nil
}
