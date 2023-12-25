package handlers

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"shop/listofsale-service/internal/database"
)

func ListOfSaleGetHandler(c echo.Context) error {
	list := database.GetAllSale()
	tmpl, err := template.ParseFiles("listofsale-service/web/templates/listOfSale.html")
	if err != nil {
		return err
	}
	return tmpl.Execute(c.Response(), list)
}
