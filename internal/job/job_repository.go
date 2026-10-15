package job

import "github.com/lkfox993/orkio-core/internal/storage/postgres"

type JobRepository struct {
	db *postgres.Database
}

func NewJobRepository(db *postgres.Database) *JobRepository {
	return &JobRepository{
		db: db,
	}
}
