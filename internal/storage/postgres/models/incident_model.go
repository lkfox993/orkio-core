package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type IncidentType string
type IncidentStatus string

const (
	IncidentJobFailed        IncidentType = "job_failed"
	IncidentJobTimeout       IncidentType = "job_timeout"
	IncidentRetriesExhausted IncidentType = "retries_exhausted"
)

const (
	IncidentStatusOpen     IncidentStatus = "open"
	IncidentStatusResolved IncidentStatus = "resolved"
)

type Incident struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	TenantID   uuid.UUID `gorm:"type:uuid;not null"`
	ProcessID  uuid.UUID
	WorkflowID uuid.UUID
	JobID      uuid.UUID
	Type       IncidentType // что произошло?
	Details    datatypes.JSON
	Message    string
	Status     IncidentStatus // что сейчас с инцидентом?
	CreatedAt  time.Time
	ResolvedAt *time.Time
}

// CREATE TABLE incidents (
//     id UUID PRIMARY KEY,
//     process_instance_id UUID NOT NULL,
//     workflow_id UUID NOT NULL,
//     job_id UUID,
//     activity_id TEXT,

//     type TEXT NOT NULL,
//     message TEXT NOT NULL,
//     details JSONB,

//     status TEXT NOT NULL DEFAULT 'open',
//     created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
//     resolved_at TIMESTAMPTZ
// );

// CREATE INDEX idx_incidents_open
//     ON incidents (created_at DESC)
//     WHERE status = 'open';

// CREATE INDEX idx_incidents_process
//     ON incidents (process_instance_id);
