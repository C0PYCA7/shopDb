package route

import (
	"github.com/labstack/echo/v4"
	"shop/auth-service/internal/handlers"
)

func SetupRoutesForAuth(e *echo.Echo) {
	e.GET("/authorization", handlers.AuthorizationGetHandler)
	e.POST("/authorization", handlers.AuthorizationPostHandler)
}
