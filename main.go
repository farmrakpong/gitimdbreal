package main

import (
	"log"

	"github.com/farmrakpong/goimdbreal/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routes.UserRoute(e)
	port := "191"
	log.Fatal(e.Start(":" + port))

}