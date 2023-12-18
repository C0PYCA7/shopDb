package main

import (
	"shop/sale-service/internal/database"
	"shop/server"
)

func main() {
	database.ConnectDb()
	defer database.CloseDb()
	server.StartServer("8084")
}
