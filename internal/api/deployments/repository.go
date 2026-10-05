package deployments

import (
	"context"

	"github.com/lkfox993/orkio-core/internal/storage/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *postgres.Database
}

func NewRepository(db *postgres.Database) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) getDB(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}

	return r.db.DB
}

func (r *Repository) GetByID(ctx context.Context, id int64) (*Deployment, error) {

	var deployment Deployment

	err := r.db.WithContext(ctx).
		First(&deployment, id).
		Error

	if err != nil {
		return nil, err
	}

	return &deployment, nil
}

func (r *Repository) GetAll(ctx context.Context) ([]Deployment, error) {
	var deployments []Deployment

	if err := r.db.WithContext(ctx).Find(&deployments).Error; err != nil {
		return nil, err
	}

	return deployments, nil
}

func (r *Repository) Create(ctx context.Context, deployment *Deployment, tx *gorm.DB) error {

	var db = r.getDB(tx)

	return db.WithContext(ctx).
		Create(deployment).
		Error
}

func (r *Repository) Update(ctx context.Context, id int64, deployment *Deployment) (*Deployment, error) {
	var existing Deployment

	if err := r.db.WithContext(ctx).
		First(&existing, id).
		Error; err != nil {
		return nil, err
	}

	existing.Name = deployment.Name

	if err := r.db.WithContext(ctx).
		Save(&existing).
		Error; err != nil {
		return nil, err
	}

	return &existing, nil
}

func (r *Repository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&Deployment{}, id).Error
}
