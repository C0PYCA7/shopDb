package database

import "log"

func InsertUser(name, surname, login, password, post string) error {
	_, err := db.Exec("SELECT add_employee($1,$2,$3,$4,$5)", name, surname, login, password, post)
	if err != nil {
		log.Print(err)
	}
	return err
}

func CheckLoginExists(login string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT check_login_exists($1)", login).Scan(&exists)
	if err != nil {
		log.Print(err)
		return false, err
	}
	return exists, nil
}
