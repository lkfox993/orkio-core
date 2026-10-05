package api

import (
	"github.com/lkfox993/orkio-core/internal/api/deployments"
	"github.com/lkfox993/orkio-core/internal/storage/postgres"
	"github.com/lkfox993/orkio-core/internal/storage/tikv"
	"github.com/lkfox993/orkio-core/internal/transports"
)

func RegisterModules(http *transports.HttpServer, db *postgres.Database, store *tikv.Database) error {

	deployments.NewModule(http, db, store).Init()

	return nil
}
