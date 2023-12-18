package database

import "log"

func CheckUserExistence(login string) (bool, error) {
	var exists bool
	query := "SELECT check_login_exists($1)"
	err := db.QueryRow(query, login).Scan(&exists)
	if err != nil {
		log.Print(err)
		return false, err
	}
	return exists, nil
}

func GetUserPost(login string) (string, error) {
	var post string
	query := "SELECT getUserPost($1)"
	err := db.QueryRow(query, login).Scan(&post)
	if err != nil {
		log.Print(err)
		return "", err
	}
	return post, nil
}

func GetIsAdmin(login string) (bool, error) {
	var isAdmin bool
	query := "SELECT get_is_admin($1)"
	err := db.QueryRow(query, login).Scan(&isAdmin)
	if err != nil {
		log.Print(err)
		return false, err
	}
	return isAdmin, nil
}

func GetUserPass(login string) string {
	var pass string
	query := "SELECT get_password_by_login($1)"
	err := db.QueryRow(query, login).Scan(&pass)
	if err != nil {
		log.Print(err)
		return ""
	}
	return pass
}
