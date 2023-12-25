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

func SearchProductByName(searchItem string) ([]models.Product, error) {
	var searchResults []models.Product

	query := "SELECT p.name, t.name as type, p.price, p.count FROM product p INNER JOIN type t ON p.id_type = t.id WHERE LOWER(p.name) LIKE LOWER($1)"
	rows, err := db.Query(query, "%"+searchItem+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.Name, &p.Type, &p.Price, &p.Count)
		if err != nil {
			return nil, err
		}
		searchResults = append(searchResults, p)
	}
	return searchResults, nil
}
