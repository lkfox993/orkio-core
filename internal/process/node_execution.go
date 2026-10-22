package process

import (
	"time"

	"github.com/google/uuid"
)

type NodeExecutionStatus string

const (
	NodeStatusActive    NodeExecutionStatus = "active"
	NodeStatusCompleted NodeExecutionStatus = "completed"
	NodeStatusFailed    NodeExecutionStatus = "failed"
)

type NodeExecution struct {
	ID        uuid.UUID
	TenantID  uuid.UUID
	ProcessID uuid.UUID
	NodeID    string
	Status    NodeExecutionStatus
	// Attempt     int
	// Error       string
	StartedAt   time.Time
	CompletedAt *time.Time
}
