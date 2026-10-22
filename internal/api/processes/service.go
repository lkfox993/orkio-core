package processes

import (
	"context"

	"github.com/lkfox993/orkio-core/internal/app"
	"github.com/lkfox993/orkio-core/internal/coordinator"
	"github.com/lkfox993/orkio-core/internal/process"
	"github.com/lkfox993/orkio-core/internal/storage/postgres/models"
)

type Service struct {
	repository *process.ProcessRepository
}

func NewService(repository *process.ProcessRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetByID(ctx context.Context, id int64) (*models.ProcessModel, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]models.ProcessModel, error) {
	return s.repository.GetAll(ctx)
}

func (s *Service) Create(ctx context.Context, input coordinator.CreateProcessInput) (*process.Process, error) {

	coordinator, err := app.Resolve[*coordinator.Engine]("coordinator")

	if err != nil {
		panic(err)
	}

	p, err := coordinator.CreateProcess(ctx, input.WorkflowID, input.Variables)

	if err != nil {
		return nil, err
	}

	return p, nil
}
