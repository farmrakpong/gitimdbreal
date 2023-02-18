package Movies

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Movie struct {
	ImdbID string `json:"imdbID"`
	Title  string `json:"title"`
	Year   int    `json:"year"`
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
	// fmt.Println(id)
	for _, m := range movies {
		if m.ImdbID == id {
			return c.JSON(200, m)
		}
	}
	return c.JSON(200, map[string]string{"message": "movie" + id})
}
func CreateMovies(c echo.Context) error {
	m := &Movie{}
	req := c.Request()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	defer req.Body.Close()

	fmt.Println(string(body))
	if err := c.Bind(m); err != nil {
		return c.JSON(400, err.Error())
	}
	movies = append(movies, *m)
	return c.JSON(200, m)
}
