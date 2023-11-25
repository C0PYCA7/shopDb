package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
)

func HomeGetHandler(c echo.Context) error {
	return c.File("home-service/web/templtales/home.html")
}

func HomePostHandler(c echo.Context) error {

	post, _, err := readCookie(c)
	if err != nil {
		// Обработка ошибки при чтении куки
		return c.String(http.StatusUnauthorized, "Ошибка авторизации")
	}

	log.Print(post)

	action := c.FormValue("action")

	switch action {
	case "listOfProduct":
		return c.Redirect(http.StatusSeeOther, "http://localhost:8081/listOfProduct")

	case "addProduct":
		if post == "Manager" {
			return c.Redirect(http.StatusSeeOther, "http://localhost:8082/addProduct")
		} else {
			return c.String(http.StatusForbidden, "У вас недостаточно прав для выполнения этого действия")
		}

	case "sale":
		if post == "Seller" {
			return c.Redirect(http.StatusSeeOther, "http://localhost:8081/sale")
		} else {
			return c.String(http.StatusForbidden, "У вас недостаточно прав для выполнения этого действия")
		}
	}

	return nil
}

func readCookie(c echo.Context) (string, string, error) {
	cookie, err := c.Cookie("user")
	if err != nil {
		return "", "", err
	}

	values := strings.Split(cookie.Value, "|")
	if len(values) != 2 {
		return "", "", errors.New("неверный формат куки")
	}
	return values[0], values[1], nil
}