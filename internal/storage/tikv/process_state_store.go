package tikv

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	tikverr "github.com/tikv/client-go/v2/error"

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
	tenantID uuid.UUID,
	processID uuid.UUID,
) (*process.Process, error) {

	key := processStateKey(tenantID, processID)

	txn, err := s.stateDB.Client.Begin()

	if err != nil {
		return nil, err
	}
	defer txn.Rollback()

	value, err := txn.Get(ctx, []byte(key))

	if err != nil {
		return nil, fmt.Errorf("get process %q: %w", processID, err)
	}

	var instance process.Process

	if err := json.Unmarshal(value, &instance); err != nil {
		return nil, fmt.Errorf("unmarshal process %q: %w", processID, err)
	}

	return &instance, nil
}

func (s *ProcessStateStore) Create(
	ctx context.Context,

	process *process.Process,
) error {
	key := processStateKey(process.ID, process.TenantID)

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
		return fmt.Errorf("process %q already exists", process.ID)

	case isNotFound(err):

		// Process не существует — создаём
		if err := txn.Set([]byte(key), value); err != nil {
			return fmt.Errorf("set process %q: %w", process.ID, err)
		}

		return txn.Commit(ctx)

	default:
		// Другая ошибка
		return fmt.Errorf("check process %q: %w", process.ID, err)
	}
}

func processStateKey(tenantID uuid.UUID, processID uuid.UUID) string {

	return fmt.Sprintf(
		"tenants/%s/processes/%s/state",
		tenantID,
		processID,
	)
}

func isNotFound(err error) bool {
	return errors.Is(err, tikverr.ErrNotExist)
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
