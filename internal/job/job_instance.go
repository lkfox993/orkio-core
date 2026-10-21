package job

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending   Status = "pending"
	StatusAvailable Status = "available"
	StatusClaimed   Status = "claimed"
	StatusRunning   Status = "running"
	StatusCompleted Status = "completed"
	StatusFailed    Status = "failed"
)

type Job struct {
	ID        uuid.UUID
	ProcessID uuid.UUID
	StepID    string
	Type      string

	Status      Status
	Attempt     int
	MaxAttempts int

	Error string

	CreatedAt time.Time
	Deadline  *time.Time
}
