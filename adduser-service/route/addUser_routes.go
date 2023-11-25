package route

import (
	"github.com/labstack/echo/v4"
	"shop/adduser-service/internal/handlers"
)

func SetupRoutesForAddUser(e *echo.Echo) {
	e.GET("/addUser", handlers.AddUserGetHandler)
	e.POST("/addUser", handlers.AddUserPostHandler)
}
