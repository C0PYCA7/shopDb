package main

import (
	"shop/listofsale-service/internal/database"
	"shop/server"
)

func main() {
	database.ConnectDb()
	defer database.CloseDb()
	server.StartServer("8086")
}
