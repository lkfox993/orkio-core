package deployments

import "io"

type CreateDeploymentInput struct {
	Name      string
	Resources []DeploymentResource
}

type UpdateDeploymentInput struct {
	Name string
}

type DeploymentResource struct {
	Name        string
	ContentType string
	Open        func() (io.ReadCloser, error)
}
