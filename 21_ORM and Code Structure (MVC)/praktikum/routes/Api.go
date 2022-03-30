package routes

import (
	"crud_go/app/Controllers"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.GET("/users", Controllers.GetUsersController)
	e.GET("/users/:id", Controllers.GetUserController)
	e.POST("/users", Controllers.CreateUserController)
	e.DELETE("/users/:id", Controllers.DeleteUserController)
	e.PUT("/users/:id", Controllers.UpdateUserController)
	return e
}
