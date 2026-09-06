package tikv

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lkfox993/orkio-core/internal/jobs"
	"github.com/tikv/client-go/v2/txnkv"
)

type JobRepository struct {
	db *txnkv.Client
}

func NewJobRepository(db *txnkv.Client) *JobRepository {
	return &JobRepository{
		db: db,
	}
}

func (r *JobRepository) Get(ctx context.Context, id string) (*jobs.Job, error) {

	key := jobKey(id)

	txn, err := r.db.Begin()

	if err != nil {
		return nil, err
	}

	defer txn.Rollback()

	value, err := txn.Get(ctx, []byte(key))
	if err != nil {
		return nil, fmt.Errorf("get job %q: %w", id, err)
	}

	var job jobs.Job

	if err := json.Unmarshal(value, &job); err != nil {
		return nil, fmt.Errorf("unmarshal job %q: %w", id, err)
	}

	return &job, nil
}

func (r *JobRepository) Create(ctx context.Context, job *jobs.Job) error {
	key := jobKey(job.ID)

	value, err := json.Marshal(job)
	if err != nil {
		return fmt.Errorf("marshal job %q: %w", job.ID, err)
	}

	txn, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	// Не перезаписываем существующий Job.
	exists, err := txn.Get(ctx, []byte(key))
	if err == nil && exists != nil {
		return fmt.Errorf("job %q already exists", job.ID)
	}

	if err != nil && !isNotFound(err) {
		return fmt.Errorf("check job %q: %w", job.ID, err)
	}

	if err := txn.Set([]byte(key), value); err != nil {
		return fmt.Errorf("set job %q: %w", job.ID, err)
	}

	if err := txn.Commit(ctx); err != nil {
		return fmt.Errorf("commit job %q: %w", job.ID, err)
	}

	return nil
}

func (r *JobRepository) Update(ctx context.Context, job *jobs.Job) error {
	key := jobKey(job.ID)

	value, err := json.Marshal(job)
	if err != nil {
		return fmt.Errorf("marshal job %q: %w", job.ID, err)
	}

	txn, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	if err := txn.Set([]byte(key), value); err != nil {
		return fmt.Errorf("update job %q: %w", job.ID, err)
	}

	if err := txn.Commit(ctx); err != nil {
		return fmt.Errorf("commit job %q: %w", job.ID, err)
	}

	return nil
}

func (r *JobRepository) Complete(ctx context.Context, id string) error {
	job, err := r.Get(ctx, id)
	if err != nil {
		return err
	}

	job.Status = jobs.StatusCompleted

	return r.Update(ctx, job)
}

func (r *JobRepository) Fail(
	ctx context.Context,
	id string,
	reason string,
) error {
	job, err := r.Get(ctx, id)
	if err != nil {
		return err
	}

	job.Status = jobs.StatusFailed
	job.Error = reason

	return r.Update(ctx, job)
}

func jobKey(id string) string {
	return "jobs/" + id
}

func isNotFound(err error) bool {
	// Реализовать через конкретную ошибку TiKV.
	// Зависит от версии client-go.
	return false
}
