package app

import (
	"context"

	"github.com/goioc/di"
)

type App struct{}

func New(ctx context.Context) error {

	registerDependencies()

	if err := di.InitializeContainer(); err != nil {
		return err
	}

	return nil
}
