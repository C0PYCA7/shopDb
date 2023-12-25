package database

import (
	"log"
	"shop/listofsale-service/internal/models"
	"time"
)

func GetAllSale() []models.SaleItem {
	query := "SELECT p.name AS product_name, s.quantity, c.name AS customer_name, c.surname AS customer_surname, e.name AS employee_name, pm.payment_date, pm.payment_amount " +
		"FROM payments pm " +
		"LEFT JOIN sale s ON pm.id = s.id_payment " +
		"LEFT JOIN product p ON s.id_product = p.id " +
		"LEFT JOIN customer c ON s.id_customer = c.id " +
		"LEFT JOIN employee e ON s.id_employee = e.id;"

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()

	var saleItem []models.SaleItem

	for rows.Next() {
		var item models.SaleItem
		var date time.Time
		if err := rows.Scan(&item.ProductName, &item.Quantity, &item.CustomerName, &item.CustomerSurname, &item.EmployeeName, &date, &item.Amount); err != nil {
			log.Print("Scan error:", err)
		}
		item.Date = date.Format("2006-01-02") // Форматируем дату
		saleItem = append(saleItem, item)
	}

	log.Print(saleItem)
	return saleItem
}
