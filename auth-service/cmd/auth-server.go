package main

import (
	"shop/auth-service/internal/database"
	"shop/server"
)

func main() {
	database.ConnectDb()
	defer database.CloseDb()
	server.StartServer("8080")
}
