package process

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/tikv/client-go/v2/txnkv"
)

type ProcessState struct {
	db *txnkv.Client
}

func NewProcessState(db *txnkv.Client) *ProcessState {
	return &ProcessState{
		db: db,
	}
}

func (s *ProcessState) Get(
	ctx context.Context,
	id string,
) (*Process, error) {
	key := processKey(id)

	txn, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()

	value, err := txn.Get(ctx, []byte(key))
	if err != nil {
		return nil, fmt.Errorf("get process %q: %w", id, err)
	}

	var instance Process

	if err := json.Unmarshal(value, &instance); err != nil {
		return nil, fmt.Errorf("unmarshal process %q: %w", id, err)
	}

	return &instance, nil
}

func processKey(id string) string {
	return "processes/" + id
}

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
