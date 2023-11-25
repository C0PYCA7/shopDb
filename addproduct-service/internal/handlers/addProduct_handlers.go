package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"shop/addproduct-service/internal/database"
	"shop/addproduct-service/internal/models"
	"strconv"
)

func AddProductGetHandler(c echo.Context) error {
	return c.File("addproduct-service/web/template/addProduct.html")
}

func AddProductPostHandler(c echo.Context) error {

	productPriceStr := c.FormValue("productPrice")
	productCountStr := c.FormValue("productCount")

	productPrice, err := strconv.ParseInt(productPriceStr, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid productPrice value")
	}

	productCount, err := strconv.ParseInt(productCountStr, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid productCount value")
	}

	addProduct := models.AddProduct{
		Name:  c.FormValue("productName"),
		Type:  c.FormValue("productType"),
		Price: productPrice,
		Count: productCount,
	}

	log.Print(addProduct)

	exists, err := database.CheckProductExists(addProduct.Name)
	if err != nil {
		log.Print(err)
	}

	if exists {
		err := database.InsertProductIfExists(addProduct.Name, addProduct.Type, addProduct.Count)
		if err != nil {
			log.Print(err)
		}
		return c.String(http.StatusOK, "Данные добавлены")
	} else {
		err := database.InsertProductIfNotExists(addProduct.Name, addProduct.Type, addProduct.Price, addProduct.Count)
		if err != nil {
			log.Print(err)
		}
		return c.String(http.StatusOK, "Данные добавлены")
	}
}
