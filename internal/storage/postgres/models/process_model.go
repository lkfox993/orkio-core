package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type ProcessModel struct {
	ID               uuid.UUID         `gorm:"type:uuid;primaryKey"`
	TenantID         uuid.UUID         `gorm:"type:uuid;not null"`
	WorkflowID       uuid.UUID         `gorm:"type:uuid;not null"`
	InitialVariables datatypes.JSONMap `gorm:"column:initial_variables;type:jsonb;not null;default:'{}'"`
	Status           string            `gorm:"type:varchar(32);not null;default:'pending'"`
}

func (ProcessModel) TableName() string {
	return "processes"
}
