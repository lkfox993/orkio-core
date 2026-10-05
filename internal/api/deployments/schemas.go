package deployments

type CreateDeploymentRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}

type UpdateDeploymentRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}
