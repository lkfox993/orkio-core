package workflow

import (
	"context"

	"github.com/google/uuid"
	"github.com/lkfox993/orkio-core/internal/storage/postgres"
	"github.com/lkfox993/orkio-core/internal/storage/postgres/models"
	"gorm.io/gorm"
)

type WorkflowRepository struct {
	db *postgres.Database
}

func NewWorkflowRepository(db *postgres.Database) *WorkflowRepository {

	if db == nil {
		panic("WorkflowRepository: db is nil")
	}

	return &WorkflowRepository{
		db: db,
	}
}

func (r *WorkflowRepository) getDB(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}

	return r.db.DB
}

func (r *WorkflowRepository) GetByID(ctx context.Context, id int64) (*models.WorkflowModel, error) {

	var workflow models.WorkflowModel

	err := r.db.WithContext(ctx).
		First(&workflow, id).
		Error

	if err != nil {
		return nil, err
	}

	return &workflow, nil
}

func (r *WorkflowRepository) GetAll(ctx context.Context) ([]models.WorkflowModel, error) {
	var workflows []models.WorkflowModel

	if err := r.db.WithContext(ctx).Find(&workflows).Error; err != nil {
		return nil, err
	}

	return workflows, nil
}

func (r *WorkflowRepository) GetByKey(
	ctx context.Context,
	tenantID uuid.UUID,
	key string,
	version *int,
) (*models.WorkflowModel, error) {

	var workflow models.WorkflowModel

	query := r.db.WithContext(ctx).
		Where("tenant_id = ? AND key = ?", tenantID, key)

	query = query.Preload("Versions", func(db *gorm.DB) *gorm.DB {

		if version == nil {
			return db.
				Where("tenant_id = ?", tenantID).
				Order("version DESC").
				Limit(1)
		}

		return db.
			Where("tenant_id = ? AND version = ?", tenantID, *version)
	})

	if err := query.First(&workflow).Error; err != nil {
		return nil, err
	}

	return &workflow, nil

	// var workflow models.WorkflowModel

	// err := r.db.WithContext(ctx).
	// 	Where("tenant_id = ? AND key = ?", tenantID, key).
	// 	First(&workflow).
	// 	Error

	// if err != nil {
	// 	return nil, err
	// }

	// return &workflow, nil
}

func (r *WorkflowRepository) Create(ctx context.Context, workflow *Workflow, tx *gorm.DB) error {

	var db = r.getDB(tx)

	return db.WithContext(ctx).
		Create(workflow).
		Error
}
