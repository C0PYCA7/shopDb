package main

import (
	"shop/list-service/internal/database"
	"shop/server"
)

func main() {
	database.ConnectDb()
	defer database.CloseDb()
	server.StartServer("8085")
}
