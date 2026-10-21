package workflows

import (
	"context"

	"github.com/lkfox993/orkio-core/internal/storage/postgres/models"
	"github.com/lkfox993/orkio-core/internal/workflow"
)

type Service struct {
	repository *workflow.WorkflowRepository
}

func NewService(repository *workflow.WorkflowRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetByID(ctx context.Context, id int64) (*models.WorkflowModel, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]models.WorkflowModel, error) {
	return s.repository.GetAll(ctx)
}
