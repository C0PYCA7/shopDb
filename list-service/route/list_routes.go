package route

import (
	"github.com/labstack/echo/v4"
	"shop/list-service/internal/handlers"
)

func SetupRoutesForList(e *echo.Echo) {
	e.GET("/list", handlers.ListGetHandler)
	e.POST("/list", handlers.SearchHandler)
}
