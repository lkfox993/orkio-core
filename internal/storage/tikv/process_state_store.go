package tikv

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lkfox993/orkio-core/internal/process"
)

type ProcessStateStore struct {
	stateDB *Database
}

func NewProcessStateStore(stateDB *Database) *ProcessStateStore {
	return &ProcessStateStore{
		stateDB: stateDB,
	}
}

func (s *ProcessStateStore) Get(
	ctx context.Context,
	id string,
) (*process.Process, error) {
	key := processStateKey(id)

	txn, err := s.stateDB.Client.Begin()
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()

	value, err := txn.Get(ctx, []byte(key))
	if err != nil {
		return nil, fmt.Errorf("get process %q: %w", id, err)
	}

	var instance process.Process

	if err := json.Unmarshal(value, &instance); err != nil {
		return nil, fmt.Errorf("unmarshal process %q: %w", id, err)
	}

	return &instance, nil
}

func (s *ProcessStateStore) Create(
	ctx context.Context,
	process *process.Process,
) error {
	key := processStateKey(process.ID.String())

	value, err := json.Marshal(process)

	if err != nil {
		return fmt.Errorf("marshal process %q: %w", process.ID, err)
	}

	txn, err := s.stateDB.Client.Begin()

	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer txn.Rollback()

	_, err = txn.Get(ctx, []byte(key))

	switch {
	case err == nil:
		return fmt.Errorf("%w: %s", process, process.ID)

	case !isNotFound(err):
		return fmt.Errorf("check process %q: %w", process.ID, err)
	}

	if err := txn.Set([]byte(key), value); err != nil {
		return fmt.Errorf("set process %q: %w", process.ID, err)
	}

	if err := txn.Commit(ctx); err != nil {
		return fmt.Errorf("commit process %q: %w", process.ID, err)
	}

	return nil
}

func processStateKey(id string) string {
	return "processes/" + id
}

func isNotFound(err error) bool {
	// Реализовать через конкретную ошибку TiKV.
	// Зависит от версии client-go.
	return false
}

// https://chatgpt.com/c/6ad49f93-8a34-83ed-a6f3-8d8586331610?src=history_search <------- важно чтобы понять как реалзовать tikv state

// Как я бы организовал ключи Orkio

// Process

// processes/{processID}

// Общее состояние экземпляра: статус, текущая позиция и версия.

// Variables

// processes/{processID}/variables

// Данные заказа и результаты выполнения узлов.

// Node state

// processes/{processID}/nodes/{nodeID}

// Статус конкретного узла, попытки, ошибка и результат.

// Job

// jobs/{jobID}

// Статус задачи, исполнитель, попытки и другие runtime-данные.
