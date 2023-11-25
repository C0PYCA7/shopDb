package server

import (
	"github.com/labstack/echo/v4"
	"log"
	route3 "shop/addproduct-service/route"
	route4 "shop/adduser-service/route"
	"shop/auth-service/route"
	route2 "shop/home-service/route"
)

func StartServer(port string) {
	e := echo.New()

	route.SetupRoutesForAuth(e)
	route2.SetupRoutesForHome(e)
	route3.SetupRoutesForAdd(e)
	route4.SetupRoutesForAddUser(e)
	err := e.Start(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
