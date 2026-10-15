package incident

import (
	"context"

	"github.com/lkfox993/orkio-core/internal/storage/postgres"
	"github.com/lkfox993/orkio-core/internal/storage/postgres/models"
	"gorm.io/gorm"
)

type IncidentRepository struct {
	db *postgres.Database
}

func NewIncidentRepository(db *postgres.Database) *IncidentRepository {

	return &IncidentRepository{
		db: db,
	}
}

func (r *IncidentRepository) getDB(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}

	return r.db.DB
}

func (r *IncidentRepository) GetByID(ctx context.Context, id int64) (*models.Incident, error) {

	var incident models.Incident

	err := r.db.WithContext(ctx).
		First(&incident, id).
		Error

	if err != nil {
		return nil, err
	}

	return &incident, nil
}

func (r *IncidentRepository) GetAll(ctx context.Context) ([]models.Incident, error) {
	var incidents []models.Incident

	if err := r.db.WithContext(ctx).Find(&incidents).Error; err != nil {
		return nil, err
	}

	return incidents, nil
}
