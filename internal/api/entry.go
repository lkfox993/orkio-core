package api

import (
	"github.com/lkfox993/orkio-core/internal/api/deployments"
	"github.com/lkfox993/orkio-core/internal/api/incidents"
	"github.com/lkfox993/orkio-core/internal/api/processes"
	"github.com/lkfox993/orkio-core/internal/api/workflows"
	"github.com/lkfox993/orkio-core/internal/transports"
)

func RegisterModules(http *transports.HttpServer) error {

	deployments.NewModule(http).Init()
	workflows.NewModule(http).Init()
	processes.NewModule(http).Init()
	incidents.NewModule(http).Init()

	return nil
}
