package incidents

import "github.com/labstack/echo/v5"

func RegisterRoutes(e *echo.Group, handler *Handler) {

	group := e.Group("/incidents")

	group.GET("", handler.GetAll)
	group.GET("/:id", handler.GetByID)
	group.POST("", handler.Create)

}
