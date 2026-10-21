package process

import (
	"context"

	"github.com/lkfox993/orkio-core/internal/storage/postgres"
	"github.com/lkfox993/orkio-core/internal/storage/postgres/models"
	"gorm.io/gorm"
)

type ProcessRepository struct {
	db *postgres.Database
}

func NewProcessRepository(db *postgres.Database) *ProcessRepository {
	return &ProcessRepository{
		db: db,
	}
}

func (r *ProcessRepository) getDB(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}

	return r.db.DB
}

func (r *ProcessRepository) GetByID(ctx context.Context, id int64) (*models.ProcessModel, error) {

	var process models.ProcessModel

	err := r.db.WithContext(ctx).
		First(&process, id).
		Error

	if err != nil {
		return nil, err
	}

	return &process, nil
}

func (r *ProcessRepository) GetAll(ctx context.Context) ([]models.ProcessModel, error) {
	var processes []models.ProcessModel

	if err := r.db.WithContext(ctx).Find(&processes).Error; err != nil {
		return nil, err
	}

	return processes, nil
}

func (r *ProcessRepository) Create(ctx context.Context, process *Process, tx *gorm.DB) error {

	var db = r.getDB(tx)

	return db.WithContext(ctx).
		Create(process).
		Error
}

func (r *ProcessRepository) Update(ctx context.Context, id int64, process *Process) (*Process, error) {
	var existing Process

	if err := r.db.WithContext(ctx).
		First(&existing, id).
		Error; err != nil {
		return nil, err
	}

	// existing.Name = process.Name

	if err := r.db.WithContext(ctx).
		Save(&existing).
		Error; err != nil {
		return nil, err
	}

	return &existing, nil
}

func (r *ProcessRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&Process{}, id).Error
}
