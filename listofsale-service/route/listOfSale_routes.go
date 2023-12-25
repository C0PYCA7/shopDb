package route

import (
	"github.com/labstack/echo/v4"
	"shop/listofsale-service/internal/handlers"
)

func SetupRoutesForStory(e *echo.Echo) {
	e.GET("/story", handlers.ListOfSaleGetHandler)
}
