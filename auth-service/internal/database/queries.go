package database

import "log"

func CheckUserExistence(login, password string) (bool, error) {
	var exists bool
	query := "SELECT checkUserExistence($1, $2)"
	err := db.QueryRow(query, login, password).Scan(&exists)
	if err != nil {
		log.Print(err)
		return false, err
	}
	return exists, nil
}

func GetUserPost(login, password string) (string, error) {
	var post string
	query := "SELECT getUserPost($1, $2)"
	err := db.QueryRow(query, login, password).Scan(&post)
	if err != nil {
		log.Print(err)
		return "", err
	}
	return post, nil
}
