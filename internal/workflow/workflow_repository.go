package workflow

import (
	"context"

	"github.com/google/uuid"
	"github.com/lkfox993/orkio-core/internal/storage/postgres"
	"gorm.io/gorm"
)

type WorkflowRepository struct {
	db *postgres.Database
}

func NewWorkflowRepository(db *postgres.Database) *WorkflowRepository {

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

func (r *WorkflowRepository) GetByID(ctx context.Context, id int64) (*Workflow, error) {

	var workflow Workflow

	err := r.db.WithContext(ctx).
		First(&workflow, id).
		Error

	if err != nil {
		return nil, err
	}

	return &workflow, nil
}

func (r *WorkflowRepository) GetAll(ctx context.Context) ([]Workflow, error) {
	var workflows []Workflow

	if err := r.db.WithContext(ctx).Find(&workflows).Error; err != nil {
		return nil, err
	}

	return workflows, nil
}

func (r *WorkflowRepository) GetByKey(
	ctx context.Context,
	tenantID uuid.UUID,
	key string,
) (*Workflow, error) {

	var workflow Workflow

	err := r.db.WithContext(ctx).
		Where("tenant_id = ? AND key = ?", tenantID, key).
		First(&workflow).
		Error

	if err != nil {
		return nil, err
	}

	return &workflow, nil
}

func (r *WorkflowRepository) Create(ctx context.Context, workflow *Workflow, tx *gorm.DB) error {

	var db = r.getDB(tx)

	return db.WithContext(ctx).
		Create(workflow).
		Error
}
