package database

import (
	"log"
	"time"
)

func CheckProductExists(productName string) (bool, error) {
	var productExists bool
	err := db.QueryRow("SELECT check_product_exists($1)", productName).Scan(&productExists)
	if err != nil {
		return false, err
	}
	return productExists, nil
}

func GetProductCount(productName string) int {
	var count int
	err := db.QueryRow("SELECT get_product_count($1)", productName).Scan(&count)
	if err != nil {
		return 0
	}
	return count
}

func GetProductPrice(productName string) int {
	var price int
	err := db.QueryRow("SELECT get_product_price($1)", productName).Scan(&price)
	if err != nil {
		return 0
	}
	return price
}

func InsertSaleData(employeeName, customerName, customerSurname, productName string, count, price int) error {
	date := time.Now()
	err := db.QueryRow("SELECT insert_sale_data($1,$2,$3,$4,$5,$6, $7)", employeeName, customerName, customerSurname, productName, count, price, date)
	if err != nil {
		log.Print(err)
	}
	return nil
}
