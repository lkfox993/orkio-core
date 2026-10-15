package incidents

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetAll(c *echo.Context) error {
	incidents, err := h.service.GetAll(c.Request().Context())

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, incidents)
}

func (h *Handler) GetByID(c *echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	incident, err := h.service.repository.GetByID(c.Request().Context(), id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, incident)
}
