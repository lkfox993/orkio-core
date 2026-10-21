package tikv

import (
	"fmt"

	"github.com/tikv/client-go/v2/txnkv"
)

type Config struct {
	Endpoints []string
}

type Database struct {
	Client *txnkv.Client
}

func New(cfg Config) (*Database, error) {

	if len(cfg.Endpoints) == 0 {
		return nil, fmt.Errorf("tikv endpoints are required")
	}

	client, err := txnkv.NewClient(cfg.Endpoints)

	if err != nil {
		return nil, fmt.Errorf("connect to tikv: %w", err)
	}

	return &Database{
		Client: client,
	}, nil
}

func (s *Database) Close() error {
	return s.Client.Close()
}
