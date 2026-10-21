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
}

func (WorkflowModel) TableName() string {
	return "workflows"
}

type WorkflowVersion struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	TenantID     uuid.UUID `json:"tenant_id" gorm:"type:uuid;not null"`
	WorkflowID   uuid.UUID `json:"workflow_id" gorm:"type:uuid;not null"`
	DeploymentID uuid.UUID `json:"deployment_id" gorm:"type:uuid;not null"`
	Version      int       `gorm:"not null"`
	ResourceName string
	Definition   string `gorm:"type:jsonb;not null"`
	CreatedAt    time.Time
}

// func (m *WorkflowModel) ToDomain() *workflow.Workflow {
// 	return &workflow.Workflow{
// 		ID: m.ID.String(),
// 	}
// }
