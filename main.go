package main

import (
	"fmt"
	"log"

	"github.com/farmrakpong/goimdbreal/pkg/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	db := Connect()
	// fmt.Printf("Connecting to %v\n", &db)
	fmt.Println(*db)
	e := echo.New()
	routes.UserRoute(e)
	port := "191"
	log.Fatal(e.Start(":" + port))
	// query := "SELECT * FROM `Product"

	// // execute the query
	// rows, err := db.Query(query)
	// if err != nil {
	// 	fmt.Println("rows")
	// 	// handle error
	// }
	// fmt.Println(rows)
}
