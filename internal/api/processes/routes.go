package processes

import "github.com/labstack/echo/v5"

func RegisterRoutes(e *echo.Group, handler *Handler) {

	group := e.Group("/processes")

	group.GET("", handler.GetAll)
	group.GET("/:id", handler.GetByID)
	group.POST("", handler.Create)

}
