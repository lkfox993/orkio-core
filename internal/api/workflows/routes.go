package workflows

import "github.com/labstack/echo/v5"

func RegisterRoutes(e *echo.Echo, handler *Handler) {

	group := e.Group("/workflows")

	group.GET("", handler.GetAll)
	group.GET("/:id", handler.GetByID)

}
