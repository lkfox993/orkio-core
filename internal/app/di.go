package app

import (
	"fmt"

	"github.com/goioc/di"
	// "github.com/lkfox993/orkio-core/internal/coordinator"
	"github.com/lkfox993/orkio-core/internal/coordinator"
	"github.com/lkfox993/orkio-core/internal/storage/postgres"
	"github.com/lkfox993/orkio-core/internal/storage/tikv"
)

func registerDependencies() error {

	db, err := postgres.New( /*os.Getenv("DATABASE_URL")*/ "postgres://orkio:password@ork-postgres:5432/orkio")

	if err != nil {
		return err
	}

	if _, err := di.RegisterBeanInstance("database", db); err != nil {
		return err
	}

	stateDB, err := tikv.New(tikv.Config{
		Endpoints: []string{
			"http://ork-tikv-pd:2379",
		},
	})

	if err != nil {
		return err
	}

	if _, err := di.RegisterBeanInstance("store", stateDB); err != nil {
		return err
	}

	stateStore := tikv.NewStateStore(stateDB)

	coordinator := coordinator.NewEngine(stateStore, db)

	if _, err := di.RegisterBeanInstance("coordinator", coordinator); err != nil {
		return err
	}

	return nil
}

func Resolve[T any](name string) (T, error) {
	var zero T

	instance, err := di.GetInstanceSafe(name)
	if err != nil {
		return zero, err
	}

	value, ok := instance.(T)
	if !ok {
		return zero, fmt.Errorf("bean %q has unexpected type", name)
	}

	return value, nil
}
