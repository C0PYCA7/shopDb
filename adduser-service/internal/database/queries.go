package database

func InsertUser(name, surname, login, password, post string) error {
	_, err := db.Exec("SELECT add_user($1,$2,$3,$4,$5)", name, surname, login, password, post)
	return err
}
