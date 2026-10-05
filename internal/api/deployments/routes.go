package deployments

import "github.com/labstack/echo/v5"

func RegisterRoutes(e *echo.Echo, handler *Handler) {

	group := e.Group("/deployments")

	group.GET("", handler.GetAll)
	group.GET("/:id", handler.GetByID)
	group.POST("", handler.Create)

}
