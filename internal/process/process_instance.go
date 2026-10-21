package process

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Status string

const (
	StatusPending   Status = "pending"
	StatusRunning   Status = "running"
	StatusCompleted Status = "completed"
	StatusFailed    Status = "failed"
)

type Process struct {
	ID               uuid.UUID
	WorkflowID       uuid.UUID
	Status           Status
	InitialVariables datatypes.JSONMap
}

type CreateProcessInput struct {
	WorkflowID string
	Version    *int
	Variables  map[string]any
}
