package user

import (
	"fmt"
	"strconv"

	db "github.com/farmrakpong/goimdbreal/db"
	userRepository "github.com/farmrakpong/goimdbreal/pkg/repository"
	"github.com/labstack/echo/v4"
	_ "gorm.io/gorm"
)

type Movie struct {
	ImdbID string `json:"imdbID"`
	Title  string `json:"title"`
	Year   int    `json:"year"`
}
type Status struct {
	Status int 
	StatusText string 
}

var movies = []Movie{
	{
		ImdbID: "tt0111161",
		Title:  "The Shawshank Redemption",
		Year:   1994,
	},
}

func GetAllMoviesHandler(c echo.Context) error {
	y := c.QueryParam("year")

	if y == "" {
		return c.JSON(200, movies)
	}
	year, err := strconv.Atoi(y)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	ms := []Movie{}
	for _, m := range movies {
		if m.Year == year {
			ms = append(ms, m)
		}
	}
	return c.JSON(200, ms)
}
func GetAllMoviesByIDHandlerByID(c echo.Context) error {
	id := c.Param("id")
	for _, m := range movies {
		if m.ImdbID == id {
			return c.JSON(200, m)
		}
	}
	return c.JSON(200, map[string]string{"message": "movie" + id})
	
}
func CreateUser(c echo.Context) error {
	conn := db.Connect()
	user := new(userRepository.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	val := conn.Create(user)
	if val.Error != nil {
		fmt.Println("<----err CreateUser----->")
		return c.JSON(400, Status{Status: 400 , StatusText:"bad Create"})
	}
	return c.JSON(200, Status{Status: 200,StatusText:"success !!!"})
}
