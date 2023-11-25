package main

import (
	"shop/adduser-service/internal/database"
	"shop/server"
)

func main() {
	database.ConnectDb()
	server.StartServer("8083")
	database.CloseDb()
}
