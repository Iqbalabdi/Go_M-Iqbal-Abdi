package routes

import (
	c "crud_go/app/Controllers"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.GET("/users", c.GetUsersController)
	e.GET("/users/:id", c.GetUserController)
	e.POST("/users", c.CreateUserController)
	e.DELETE("/users/:id", c.DeleteUserController)
	e.PUT("/users/:id", c.UpdateUserController)
	return e
}
