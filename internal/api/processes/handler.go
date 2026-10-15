package processes

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
	processes, err := h.service.GetAll(c.Request().Context())

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, processes)
}

func (h *Handler) GetByID(c *echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	process, err := h.service.repository.GetByID(c.Request().Context(), id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, process)
}

func (h *Handler) Create(c *echo.Context) error {

	req := new(CreateProcessRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	// process, err := h.service.Create(c.Request().Context(), coordinator.CreateProcessInput{
	// 	WorkflowID: req.WorkflowID,
	// 	Version:    req.Version,
	// 	Variables:  req.Variables,
	// })

	// if err != nil {
	// 	return err
	// }

	// return c.JSON(http.StatusCreated, process)

	return nil
}
