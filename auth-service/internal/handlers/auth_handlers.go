package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"shop/auth-service/internal/database"
	"shop/auth-service/internal/models"
)

func AuthorizationGetHandler(c echo.Context) error {
	return c.File("auth-service/web/templates/authorization.html")
}

func AuthorizationPostHandler(c echo.Context) error {
	user := models.User{
		Login:    c.FormValue("login"),
		Password: c.FormValue("password"),
	}

	exists, err := database.CheckUserExistence(user.Login, user.Password)
	if err != nil {
		log.Print("Ошибка при проверке пользователя", err)
	}
	post, err := database.GetUserPost(user.Login, user.Password)
	if err != nil {
		log.Print("Ошибка при получении должности", err)
	}

	if exists {
		writeCookie(user.Login, post, c)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Ошибка сервера")
		}
		return c.Redirect(http.StatusSeeOther, "http://localhost:8081/home")
	} else {
		return c.String(http.StatusUnauthorized, "Ошибка авторизации")
	}
}

func writeCookie(username string, role string, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = role + "|" + username
	c.SetCookie(cookie)
}
