package transports

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/lkfox993/orkio-core/internal/common"
)

type HttpServer struct {
	Echo *echo.Echo
}

func NewHttpServer() *HttpServer {

	e := echo.New()

	server := &HttpServer{
		Echo: e,
	}

	return server
}

func (s *HttpServer) Configure() error {

	s.Echo.Validator = common.NewCustomValidator()

	s.Echo.Use(middleware.CORS("*"))

	s.Echo.GET("/health", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	s.Echo.GET("/ready", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ready",
		})
	})

	return nil
}

func (s *HttpServer) Start(addr string) error {
	return s.Echo.Start(addr)
}

func (s *HttpServer) Shutdown() error {
	return nil
}
