package process

import "context"

type Repository interface {
	Get(ctx context.Context, id string) (*Process, error)
	Create(ctx context.Context, j *Process) error
	Update(ctx context.Context, j *Process) error
	Complete(ctx context.Context, id string) error
}
