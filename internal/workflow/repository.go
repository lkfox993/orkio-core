package workflow

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Get(ctx context.Context, id string) (*Workflow, error)
	Create(ctx context.Context, j *Workflow) error
	Update(ctx context.Context, j *Workflow) error
	GetByKey(ctx context.Context, tenantID uuid.UUID, key string) (*Workflow, error)
}
