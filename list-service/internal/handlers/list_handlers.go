package handlers

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"shop/list-service/internal/database"
)

func ListGetHandler(c echo.Context) error {
	product := database.GetAllProducts()
	tmpl, err := template.ParseFiles("list-service/web/templates/listOfProduct.html")
	if err != nil {
		return err
	}

	return tmpl.Execute(c.Response(), product)
}
