package events

import "time"

const (
	TypeJobAvailable Type = "job.available"
	TypeJobCompleted Type = "job.completed"
	TypeJobFailed    Type = "job.failed"
	TypeJobTimedOut  Type = "job.timed_out"
)

type JobCompleted struct {
	JobID       string         `json:"job_id"`
	Output      map[string]any `json:"output,omitempty"`
	CompletedAt time.Time      `json:"completed_at"`
}

type JobFailed struct {
	JobID    string    `json:"job_id"`
	Error    string    `json:"error"`
	FailedAt time.Time `json:"failed_at"`
}

type JobTimedOut struct {
	JobID      string    `json:"job_id"`
	TimedOutAt time.Time `json:"timed_out_at"`
}
