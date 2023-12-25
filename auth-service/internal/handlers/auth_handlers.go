package handlers

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"shop/auth-service/internal/database"
	"shop/auth-service/internal/models"
)

var logger *log.Logger

func init() {
	file, err := os.OpenFile("auth.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Не удалось открыть файл для записи логов:", err)
	}
	logger = log.New(file, "AUTH: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func AuthorizationGetHandler(c echo.Context) error {
	return c.File("auth-service/web/templates/authorization.html")
}

func AuthorizationPostHandler(c echo.Context) error {
	user := models.User{
		Login:    c.FormValue("login"),
		Password: c.FormValue("password"),
	}

	exists, err := database.CheckUserExistence(user.Login)
	if err != nil {
		log.Print("Ошибка при проверке пользователя", err)
	}

	if exists {
		pass := database.GetUserPass(user.Login)

		compare, err := checkPassword(user.Password, pass)
		if err != nil {
			log.Print(err)
		}
		if compare {
			post, err := database.GetUserPost(user.Login)
			if err != nil {
				log.Print("Ошибка при получении должности", err)
			}

			log.Print(post)

			isAdmin, err := database.GetIsAdmin(user.Login)
			if err != nil {
				log.Print("Ошибка при проверке админа", err)
			}

			log.Print(isAdmin)

			uName := database.GetUserName(user.Login)

			log.Print(uName)

			writeCookie(uName, post, isAdmin, c)
			if err != nil {
				return c.String(http.StatusInternalServerError, "Ошибка сервера")
			}

			logger.Printf("пользователь %s зашел на сайт", user.Login)

			return c.Redirect(http.StatusSeeOther, "http://localhost:8081/home")
		} else {
			return c.String(http.StatusBadRequest, "Неверный пароль")
		}

	} else {
		return c.String(http.StatusUnauthorized, "Неверный логин")
	}
}

func writeCookie(username string, role string, isAdmin bool, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = role + "|" + username + "|" + encodeBool(isAdmin)
	c.SetCookie(cookie)
}

func encodeBool(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func checkPassword(inputPass, passDb string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(passDb), []byte(inputPass))
	if err != nil {
		log.Print(err)
		return false, err
	}
	return true, nil
}
