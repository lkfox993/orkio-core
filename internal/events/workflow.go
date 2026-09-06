package events

import "time"

const (
	TypeWorkflowStarted   Type = "workflow.started"
	TypeWorkflowCompleted Type = "workflow.completed"
	TypeWorkflowFailed    Type = "workflow.failed"
	TypeWorkflowCanceled  Type = "workflow.canceled"
)

type WorkflowStarted struct {
	WorkflowInstanceID string    `json:"workflow_instance_id"`
	DefinitionID       string    `json:"definition_id"`
	StartedAt          time.Time `json:"started_at"`
}

type WorkflowCompleted struct {
	WorkflowInstanceID string    `json:"workflow_instance_id"`
	CompletedAt        time.Time `json:"completed_at"`
}

type WorkflowFailed struct {
	WorkflowInstanceID string    `json:"workflow_instance_id"`
	JobID              string    `json:"job_id,omitempty"`
	StepID             string    `json:"step_id,omitempty"`
	Error              string    `json:"error"`
	FailedAt           time.Time `json:"failed_at"`
}

type WorkflowCanceled struct {
	WorkflowInstanceID string    `json:"workflow_instance_id"`
	CanceledAt         time.Time `json:"canceled_at"`
}
