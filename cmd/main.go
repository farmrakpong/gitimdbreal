package main

import (
	"fmt"
	"log"

	_ "github.com/farmrakpong/goimdbreal/db"
	userRepository "github.com/farmrakpong/goimdbreal/pkg/repository"
	"github.com/farmrakpong/goimdbreal/pkg/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	s := userRepository.UserRepository()
	// user := userRepository.User{Name: "John", TotalProduct: "5"}
	// fmt.Printf("%v\n", user.TotalProduct)
	fmt.Print(s)
	e := echo.New()
	routes.UserRoute(e)
	port := "191"
	log.Fatal(e.Start(":" + port))
}
