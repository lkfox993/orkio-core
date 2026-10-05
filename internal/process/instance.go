package process

type Status string

const (
	StatusRunning   Status = "running"
	StatusCompleted Status = "completed"
	StatusFailed    Status = "failed"
)

type ProcessInstance struct {
	ID         string
	WorkflowID string
	Status     Status
	Variables  map[string]any
}
