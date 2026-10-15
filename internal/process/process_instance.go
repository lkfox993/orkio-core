package process

type Status string

const (
	StatusRunning   Status = "running"
	StatusCompleted Status = "completed"
	StatusFailed    Status = "failed"
)

type Process struct {
	ID         string
	WorkflowID string
	Status     Status
	Variables  map[string]any
}

type CreateProcessInput struct {
	WorkflowID string
	Version    *int
	Variables  map[string]any
}
