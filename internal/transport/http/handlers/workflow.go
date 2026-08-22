package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type WorkflowHandler struct{}

func NewWorkflowHandler() *WorkflowHandler {
	return &WorkflowHandler{}
}

func (h *WorkflowHandler) RegisterRoutes(e *echo.Group) {

	r := e.Group("/workflows")
	r.GET("", h.List)
	r.GET(":key", h.Get)

}

func (h *WorkflowHandler) List(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func (h *WorkflowHandler) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
