package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"shop/sale-service/internal/database"
	"shop/sale-service/internal/models"
	"strconv"
	"strings"
)

func SaleGetHandler(c echo.Context) error {
	return c.File("sale-service/web/templates/sale.html")
}

func SalePostHandler(c echo.Context) error {

	_, name, _, err := readCookie(c)
	if err != nil {
		return c.String(http.StatusUnauthorized, "Ошибка авторизации")
	}

	quantityStr := c.FormValue("quantity")

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid quantity value")
	}

	user := models.User{
		SellerName:   toUpperCase(name),
		BuyerName:    toUpperCase(c.FormValue("buyerName")),
		BuyerSurname: toUpperCase(c.FormValue("buyerSurname")),
	}
	item := models.Item{
		Name:     c.FormValue("prodName"),
		Quantity: quantity,
	}

	log.Print(user)
	log.Print(item)

	exists, err := database.CheckProductExists(item.Name)
	if err != nil {
		log.Print(err)
	}

	if exists {
		count, price := database.GetProductCount(item.Name), database.GetProductPrice(item.Name)
		log.Print(count, price)

		if count >= item.Quantity && price != 0 {
			err := database.InsertSaleData(user.SellerName, user.BuyerName, user.BuyerSurname, item.Name, item.Quantity, price)
			if err != nil {
				log.Print(err)
			}
			return c.String(http.StatusOK, "Товар продан")
		} else {
			return c.String(http.StatusBadRequest, "Нет такого продукта")
		}
	} else {
		return c.String(http.StatusBadRequest, "Нет такого продукта")
	}
}

func toUpperCase(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
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
