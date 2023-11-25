package database

import (
	"database/sql"
	"log"
)

var db *sql.DB

func ConnectDb() {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=root host=localhost dbname=courseProject sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDb() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}
