package job

import "context"

type Repository interface {
	Get(ctx context.Context, id string) (*Job, error)
	Create(ctx context.Context, j *Job) error
	Update(ctx context.Context, j *Job) error

	Complete(ctx context.Context, id string) error
	Fail(ctx context.Context, id string, reason string) error
}
