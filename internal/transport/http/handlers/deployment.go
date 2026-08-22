package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type DeploymentHandler struct{}

func NewDeploymentHandler() *DeploymentHandler {
	return &DeploymentHandler{}
}

func (h *DeploymentHandler) RegisterRoutes(e *echo.Group) {
	r := e.Group("deployments")

	r.GET("", h.List)
}

func (h *DeploymentHandler) List(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func (h *DeploymentHandler) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
