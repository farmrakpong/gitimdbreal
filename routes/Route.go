package routes

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	// Define the route for getting all users
	e.GET("/movies", getAllMoviesHandler)

	// // Define the route for getting a user by ID
	e.GET("/movies/:id", getAllMoviesByIDHandlerByID)

	// // Define the route for creating a new user
	e.POST("/users", createMovies)

	// // Define the route for updating a user by ID
	// e.PUT("/users/:id", updateUser)

	// // Define the route for deleting a user by ID
	// e.DELETE("/users/:id", deleteUser)
}

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

func getAllMoviesHandler(c echo.Context) error {
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
func getAllMoviesByIDHandlerByID(c echo.Context) error {
	id := c.Param("id")
	// fmt.Println(id)
	for _, m := range movies {
		if m.ImdbID == id {
			return c.JSON(200, m)
		}
	}
	return c.JSON(200, map[string]string{"message": "movie" + id})
}
func createMovies(c echo.Context) error {
	m := &Movie{}

	if err := c.Bind(m); err != nil {
		return c.JSON(400, err.Error())
	}
	movies = append(movies, *m)
	return c.JSON(200, m)
}

// func getUser(c echo.Context) error {
// 	// Add logic for getting a user by ID
// 	return c.JSON(http.StatusOK, "User retrieved successfully")
// }

// func createUser(c echo.Context) error {
// 	// Add logic for creating a new user
// 	return c.JSON(http.StatusOK, "User created successfully")
// }

// func updateUser(c echo.Context) error {
// 	// Add logic for updating a user by ID
// 	return c.JSON(http.StatusOK, "User updated successfully")
// }

// func deleteUser(c echo.Context) error {
// 	// Add logic for deleting a user by ID
// 	return c.JSON(http.StatusOK, "User deleted successfully")
// }
