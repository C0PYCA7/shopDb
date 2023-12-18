package route

import (
	"github.com/labstack/echo/v4"
	"shop/sale-service/internal/handlers"
)

func SetupRoutesForSale(e *echo.Echo) {
	e.GET("/sale", handlers.SaleGetHandler)
	e.POST("/sale", handlers.SalePostHandler)
}
