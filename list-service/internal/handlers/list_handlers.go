package handlers

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"log"
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

func SearchHandler(c echo.Context) error {
	searchTerm := c.FormValue("search")
	searchResults, err := database.SearchProductByName(searchTerm)
	if err != nil {
		log.Print(err)
	}

	log.Print(searchTerm)
	tmpl, err := template.ParseFiles("list-service/web/templates/listOfProduct.html")
	if err != nil {
		return err
	}

	return tmpl.Execute(c.Response(), searchResults)
}
