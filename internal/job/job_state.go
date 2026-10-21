package job

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/tikv/client-go/v2/txnkv"
)

type JobState struct {
	db *txnkv.Client
}

func NewJobState(db *txnkv.Client) *JobState {
	return &JobState{
		db: db,
	}
}

func (s *JobState) Get(ctx context.Context, id string) (*Job, error) {

	key := jobKey(id)

	txn, err := s.db.Begin()

	if err != nil {
		return nil, err
	}

	defer txn.Rollback()

	value, err := txn.Get(ctx, []byte(key))

	if err != nil {
		return nil, fmt.Errorf("get job %q: %w", id, err)
	}

	var job Job

	if err := json.Unmarshal(value, &job); err != nil {
		return nil, fmt.Errorf("unmarshal job %q: %w", id, err)
	}

	return &job, nil
}

func (s *JobState) Create(ctx context.Context, job *Job) error {
	key := jobKey(job.ID.String())

	value, err := json.Marshal(job)

	if err != nil {
		return fmt.Errorf("marshal job %q: %w", job.ID, err)
	}

	txn, err := s.db.Begin()

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

func (s *JobState) Update(ctx context.Context, job *Job) error {
	key := jobKey(job.ID.String())

	value, err := json.Marshal(job)
	if err != nil {
		return fmt.Errorf("marshal job %q: %w", job.ID, err)
	}

	txn, err := s.db.Begin()

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

func (s *JobState) Complete(ctx context.Context, id string) error {
	j, err := s.Get(ctx, id)
	if err != nil {
		return err
	}

	j.Status = StatusCompleted

	return s.Update(ctx, j)
}

func (s *JobState) Fail(
	ctx context.Context,
	id string,
	reason string,
) error {
	j, err := s.Get(ctx, id)
	if err != nil {
		return err
	}

	j.Status = StatusFailed
	j.Error = reason

	return s.Update(ctx, j)
}

func jobKey(id string) string {
	return "jobs/" + id
}

func isNotFound(err error) bool {
	// Реализовать через конкретную ошибку TiKV.
	// Зависит от версии client-go.
	return false
}
