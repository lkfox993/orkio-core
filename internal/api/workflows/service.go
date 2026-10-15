package workflows

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

func (s *Service) GetByID(ctx context.Context, id int64) (*Workflow, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]Workflow, error) {
	return s.repository.GetAll(ctx)
}
