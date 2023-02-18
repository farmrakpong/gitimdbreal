package routes

import (
	users "github.com/farmrakpong/goimdbreal/pkg/controller"
	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	// Define the route for getting all users
	e.GET("/movies", users.GetAllMoviesHandler)

	// // Define the route for getting a user by ID
	e.GET("/movies/:id", users.GetAllMoviesByIDHandlerByID)

	// // Define the route for creating a new user
	e.POST("/createusers", users.CreateUser)

	// // Define the route for updating a user by ID
	// e.PUT("/users/:id", updateUser)

	// // Define the route for deleting a user by ID
	// e.DELETE("/users/:id", deleteUser)
}
