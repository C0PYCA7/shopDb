package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"shop/adduser-service/internal/database"
	"shop/adduser-service/internal/models"
	"strings"
)

var logger *log.Logger

func init() {
	file, err := os.OpenFile("addUser.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Не удалось открыть файл для записи логов:", err)
	}
	logger = log.New(file, "ADDUSER: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func AddUserGetHandler(c echo.Context) error {
	return c.File("adduser-service/web/template/addUser.html")
}

func AddUserPostHandler(c echo.Context) error {
	user := models.AddUser{
		Name:     toUpperCase(c.FormValue("name")),
		Surname:  toUpperCase(c.FormValue("surname")),
		Post:     toUpperCase(c.FormValue("post")),
		Login:    c.FormValue("login"),
		Password: hash(c.FormValue("password")),
	}
	log.Print(user)

	_, login, _, err := readCookie(c)
	if err != nil {
		log.Print(err)
	}
	exists, err := database.CheckLoginExists(user.Login)
	if err != nil {
		log.Print(err)
	}

	log.Print(exists)

	if !exists {
		err := database.InsertUser(user.Name, user.Surname, user.Login, user.Password, user.Post)
		if err != nil {
			return err
		}
		logger.Printf("Пользователь %s добавил пользователя %s %s ", login, user.Name, user.Surname)
		return c.String(http.StatusOK, "Данные добавлены")
	} else {
		return c.String(http.StatusBadRequest, "Логин занят")
	}
}

func toUpperCase(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func hash(s string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		log.Print(err)
	}
	return string(hash)
}

func readCookie(c echo.Context) (string, string, string, error) {
	cookie, err := c.Cookie("user")
	if err != nil {
		return "", "", "", err
	}

	values := strings.Split(cookie.Value, "|")
	if len(values) != 3 {
		return "", "", "", errors.New("неверный формат куки")
	}
	return values[0], values[1], values[2], nil
}
