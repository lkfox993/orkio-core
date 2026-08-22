package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lkfox993/orkio-core/internal/transport/http/handlers"
)

func (s *Server) registerRoutes() {

	v1 := s.echo.Group("/v1")

	workflowHandler := handlers.NewWorkflowHandler()
	workflowHandler.RegisterRoutes(v1)

	deploymentHandler := handlers.NewDeploymentHandler()
	deploymentHandler.RegisterRoutes(v1)

	s.echo.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	s.echo.GET("/ready", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ready",
		})
	})
}
