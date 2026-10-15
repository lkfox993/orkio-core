package processes

type CreateProcessRequest struct {
	// Name string `json:"name" validate:"required,min=3,max=100"`
	WorkflowID string         `json:"workflowId" validate:"required,min=3,max=100"`
	Version    *int           `json:"version,omitempty" validate:"required,min=1"`
	Variables  map[string]any `json:"variables,omitempty"`
}
