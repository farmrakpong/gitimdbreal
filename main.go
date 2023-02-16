package main

import (
	"log"

	"github.com/farmrakpong/goimdbreal/pkg/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	db := Connect()
	defer db.Close()
	e := echo.New()
	routes.UserRoute(e)
	port := "191"
	log.Fatal(e.Start(":" + port))

}
