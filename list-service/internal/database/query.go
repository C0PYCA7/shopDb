package database

import (
	"log"
	"shop/list-service/internal/models"
)

func GetAllProducts() []models.Product {
	rows, err := db.Query("SELECT * FROM get_product_details()")
	if err != nil {
		log.Print("Query error:", err)
	}
	defer rows.Close()

	var product []models.Product

	for rows.Next() {
		var detail models.Product
		if err := rows.Scan(&detail.Name, &detail.Type, &detail.Price, &detail.Count); err != nil {
			log.Print("Scan error:", err)
		}
		product = append(product, detail)
	}

	log.Print(product)
	return product
}
