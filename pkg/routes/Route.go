package routes

import (
	movies "github.com/farmrakpong/goimdbreal/pkg/controller/movies"
	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	// Define the route for getting all users
	e.GET("/movies", movies.GetAllMoviesHandler)

	// // Define the route for getting a user by ID
	e.GET("/movies/:id", movies.GetAllMoviesByIDHandlerByID)

	// // Define the route for creating a new user
	e.POST("/users", movies.CreateMovies)

	// // Define the route for updating a user by ID
	// e.PUT("/users/:id", updateUser)

	// // Define the route for deleting a user by ID
	// e.DELETE("/users/:id", deleteUser)
}
