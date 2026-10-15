package models

import "github.com/google/uuid"

type Process struct {
	ID       string    `gorm:"type:uuid;primaryKey"`
	TenantID uuid.UUID `gorm:"type:uuid;not null"`
}
