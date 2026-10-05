package deployments

import (
	"github.com/lkfox993/orkio-core/internal/storage/postgres"
	"github.com/lkfox993/orkio-core/internal/storage/tikv"
	"github.com/lkfox993/orkio-core/internal/transports"
)

type Module struct {
	h     *transports.HttpServer
	db    *postgres.Database
	store *tikv.Database
}

func NewModule(h *transports.HttpServer, db *postgres.Database, store *tikv.Database) *Module {

	return &Module{
		h:     h,
		db:    db,
		store: store,
	}
}

func (m *Module) Init() {

	repository := NewRepository(m.db)
	service := NewService(repository)
	handler := NewHandler(service)

	RegisterRoutes(m.h.Echo, handler)

}
