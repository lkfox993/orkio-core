package deployments

import (
	"time"

	"github.com/google/uuid"
)

type Deployment struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	TenantID  uuid.UUID `gorm:"type:uuid;not null"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
