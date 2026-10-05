package deployments

import (
	"io"
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

func (h *Handler) GetByID(c *echo.Context) error {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	campaign, err := h.service.repository.GetByID(c.Request().Context(), id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, campaign)
}

func (h *Handler) GetAll(c *echo.Context) error {

	campaigns, err := h.service.GetAll(c.Request().Context())

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, campaigns)
}

func (h *Handler) Create(c *echo.Context) error {

	form, err := c.MultipartForm()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid multipart form")
	}

	// req := new(CreateDeploymentRequest)

	req := CreateDeploymentRequest{
		Name: form.Value["name"][0],
	}

	// if err := c.Bind(req); err != nil {
	// 	return err
	// }

	if err := c.Validate(req); err != nil {
		return err
	}

	files := form.File["files"]
	resources := make([]DeploymentResource, 0, len(files))

	for _, file := range files {
		resources = append(resources, DeploymentResource{
			Name:        file.Filename,
			ContentType: file.Header.Get("Content-Type"),
			Open: func() (io.ReadCloser, error) {
				return file.Open()
			},
		})
	}

	deployment, err := h.service.Create(c.Request().Context(), CreateDeploymentInput{
		Name:      req.Name,
		Resources: resources,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, deployment)
}

func (h *Handler) Update(c *echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid campaign id")
	}

	req := new(UpdateDeploymentRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	campaign, err := h.service.Update(
		c.Request().Context(),
		id,
		UpdateDeploymentInput{
			Name: req.Name,
		},
	)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, campaign)
}

func (h *Handler) Delete(c *echo.Context) error {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid campaign id")
	}

	if err := h.service.Delete(c.Request().Context(), id); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
