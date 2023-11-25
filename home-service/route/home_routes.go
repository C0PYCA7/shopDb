package route

import (
	"github.com/labstack/echo/v4"
	"shop/home-service/internal/handlers"
)

func SetupRoutesForHome(e *echo.Echo) {
	e.GET("/home", handlers.HomeGetHandler)
	e.POST("/home", handlers.HomePostHandler)

}
