package deployment

import (
	"context"

	"github.com/lkfox993/orkio-core/internal/storage/postgres"
	"github.com/lkfox993/orkio-core/internal/storage/postgres/models"
	"gorm.io/gorm"
)

type DeploymentRepository struct {
	DB *postgres.Database
}

func NewDeploymentRepository(db *postgres.Database) *DeploymentRepository {

	return &DeploymentRepository{
		DB: db,
	}
}

func (r *DeploymentRepository) getDB(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}

	return r.DB.DB
}

func (r *DeploymentRepository) GetByID(ctx context.Context, id int64) (*models.Deployment, error) {

	var deployment models.Deployment

	err := r.DB.WithContext(ctx).
		First(&deployment, id).
		Error

	if err != nil {
		return nil, err
	}

	return &deployment, nil
}

func (r *DeploymentRepository) GetAll(ctx context.Context) ([]models.Deployment, error) {
	var deployments []models.Deployment

	if err := r.DB.WithContext(ctx).Find(&deployments).Error; err != nil {
		return nil, err
	}

	return deployments, nil
}

func (r *DeploymentRepository) Create(ctx context.Context, deployment *models.Deployment, tx *gorm.DB) error {

	var db = r.getDB(tx)

	return db.WithContext(ctx).
		Create(deployment).
		Error
}

func (r *DeploymentRepository) Update(ctx context.Context, id int64, deployment *models.Deployment) (*models.Deployment, error) {
	var existing models.Deployment

	if err := r.DB.WithContext(ctx).
		First(&existing, id).
		Error; err != nil {
		return nil, err
	}

	existing.Name = deployment.Name

	if err := r.DB.WithContext(ctx).
		Save(&existing).
		Error; err != nil {
		return nil, err
	}

	return &existing, nil
}

func (r *DeploymentRepository) Delete(ctx context.Context, id int64) error {
	return r.DB.WithContext(ctx).Delete(&models.Deployment{}, id).Error
}
