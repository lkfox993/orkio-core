package processes

import (
	"context"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetByID(ctx context.Context, id int64) (*Process, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]Process, error) {
	return s.repository.GetAll(ctx)
}

func (s *Service) Create(ctx context.Context, input any /*coordinator.CreateProcessInput*/) (*Process, error) {
	// return s.engine.CreateProcess(ctx, input)
	return nil, nil
}
