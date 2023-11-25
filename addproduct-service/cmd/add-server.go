package main

import (
	"shop/addproduct-service/internal/database"
	"shop/server"
)

func main() {
	database.ConnectDb()
	defer database.CloseDb()
	server.StartServer("8082")
}
