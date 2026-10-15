package workflows

import (
	"github.com/lkfox993/orkio-core/internal/app"
	"github.com/lkfox993/orkio-core/internal/storage/postgres"
	"github.com/lkfox993/orkio-core/internal/transports"
	"github.com/lkfox993/orkio-core/internal/workflow"
)

type Module struct {
	h *transports.HttpServer
}

func NewModule(h *transports.HttpServer) *Module {

	return &Module{
		h: h,
	}
}

func (m *Module) Init() {

	db, err := app.Resolve[*postgres.Database]("database")

	if err != nil {
		panic(err)
	}

	repository := workflow.NewWorkflowRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	RegisterRoutes(m.h.Echo, handler)

}
