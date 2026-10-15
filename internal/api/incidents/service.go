package incidents

import (
	"context"

	"github.com/lkfox993/orkio-core/internal/incident"
	"github.com/lkfox993/orkio-core/internal/storage/postgres/models"
)

type Service struct {
	repository *incident.IncidentRepository
}

func NewService(repository *incident.IncidentRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetByID(ctx context.Context, id int64) (*models.Incident, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]models.Incident, error) {
	return s.repository.GetAll(ctx)
}
