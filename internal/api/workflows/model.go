package workflows

import (
	"time"

	"github.com/google/uuid"
)

type Workflow struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	TenantID  uuid.UUID
	Key       string
	Name      string
	CreatedAt time.Time
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
