package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func New(dsn string) (*Database, error) {

	dial := postgres.Open(dsn)
	db, err := gorm.Open(dial, nil)

	if err != nil {
		return nil, err
	}

	return &Database{
		DB: db,
	}, nil
}
