package route

import (
	"github.com/labstack/echo/v4"
	"shop/addproduct-service/internal/handlers"
)

func SetupRoutesForAdd(e *echo.Echo) {
	e.GET("/addProduct", handlers.AddProductGetHandler)
	e.POST("/addProduct", handlers.AddProductPostHandler)
}
