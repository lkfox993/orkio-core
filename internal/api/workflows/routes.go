package workflows

import "github.com/labstack/echo/v5"

func RegisterRoutes(e *echo.Group, handler *Handler) {

	group := e.Group("/workflows")

	group.GET("", handler.GetAll)
	group.GET("/:id", handler.GetByID)
	group.POST("", handler.Create)

}
