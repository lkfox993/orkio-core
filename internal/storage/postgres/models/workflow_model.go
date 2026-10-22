package models

import (
	"time"

	"github.com/google/uuid"
)

type WorkflowModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	TenantID  uuid.UUID
	Key       string
	Name      string
	CreatedAt time.Time
	Versions  []WorkflowVersionModel `json:"versions,omitempty" gorm:"foreignKey:WorkflowID"`
}

func (m *WorkflowModel) LoadedVersion() *WorkflowVersionModel {
	if len(m.Versions) == 0 {
		return nil
	}

	return &m.Versions[0]
}

func (WorkflowModel) TableName() string {
	return "workflows"
}

type WorkflowVersionModel struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	TenantID     uuid.UUID `json:"tenant_id" gorm:"type:uuid;not null"`
	WorkflowID   uuid.UUID `json:"workflow_id" gorm:"type:uuid;not null"`
	DeploymentID uuid.UUID `json:"deployment_id" gorm:"type:uuid;not null"`
	Version      int       `gorm:"not null"`
	ResourceName string
	Definition   string `gorm:"type:jsonb;not null"`
	CreatedAt    time.Time
}

func (WorkflowVersionModel) TableName() string {
	return "workflow_versions"
}
