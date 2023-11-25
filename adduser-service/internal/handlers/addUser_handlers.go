package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"shop/adduser-service/internal/database"
	"shop/adduser-service/internal/models"
)

func AddUserGetHandler(c echo.Context) error {
	return c.File("adduser-service/web/template/addUser.html")
}

func AddUserPostHandler(c echo.Context) error {
	user := models.AddUser{
		OldLogin: c.FormValue("oldLogin"),
		Login:    c.FormValue("login"),
		Password: c.FormValue("password"),
	}
	log.Print(user)


	database.InsertUser(user.)
}
