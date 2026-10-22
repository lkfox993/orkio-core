package tikv

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/lkfox993/orkio-core/internal/process"
	tikverr "github.com/tikv/client-go/v2/error"
)

type NodeExecutionStateStore struct {
	stateDB *Database
}

func NewNodeExecutionStateStore(
	stateDB *Database,
) *NodeExecutionStateStore {
	return &NodeExecutionStateStore{stateDB: stateDB}
}

func (s *NodeExecutionStateStore) Create(
	ctx context.Context,
	execution *process.NodeExecution,
) error {
	key := nodeExecutionKey(
		execution.TenantID,
		execution.ProcessID,
		execution.ID,
	)

	value, err := json.Marshal(execution)

	if err != nil {
		return fmt.Errorf("marshal node execution: %w", err)
	}

	txn, err := s.stateDB.Client.Begin()

	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer txn.Rollback()

	_, err = txn.Get(ctx, []byte(key))
	if err == nil {
		return fmt.Errorf("node execution %s already exists", execution.ID)
	}

	if !errors.Is(err, tikverr.ErrNotExist) {
		return fmt.Errorf("check node execution: %w", err)
	}

	if err := txn.Set([]byte(key), value); err != nil {
		return fmt.Errorf("set node execution: %w", err)
	}

	if err := txn.Commit(ctx); err != nil {
		return fmt.Errorf("commit node execution: %w", err)
	}

	return nil
}

func (s *NodeExecutionStateStore) Complete(
	ctx context.Context,
	processID uuid.UUID,
	executionID uuid.UUID,
) error {
	// Update status to completed
	return nil
}

func nodeExecutionKey(
	tenantID uuid.UUID,
	processID uuid.UUID,
	executionID uuid.UUID,
) string {
	return fmt.Sprintf(
		"tenants/%s/processes/%s/nodes/%s",
		tenantID,
		processID,
		executionID,
	)
}
