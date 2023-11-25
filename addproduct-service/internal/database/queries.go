package database

func CheckProductExists(productName string) (bool, error) {
	var productExists bool
	err := db.QueryRow("SELECT check_product_exists($1)", productName).Scan(&productExists)
	if err != nil {
		return false, err
	}
	return productExists, nil
}

func InsertProductIfNotExists(productName, productType string, productPrice, productCount int64) error {
	_, err := db.Exec("SELECT add_product_if_not_exists($1, $2, $3, $4)", productName, productType, productPrice, productCount)

	return err
}

func InsertProductIfExists(productName, productType string, productCount int64) error {
	_, err := db.Exec("SELECT increase_product_count_by_name_and_type($1,$2,$3)", productName, productType, productCount)
	return err
}
