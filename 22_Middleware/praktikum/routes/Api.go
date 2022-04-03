package routes

import (
	c "crud_go/app/Controllers"
	"crud_go/constants"
	mid "crud_go/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	mid.LogMiddleware(e)
	// Route / to handler function
	e.GET("/users", c.GetUsersController)
	e.GET("/users/:id", c.GetUserController)
	e.POST("/users", c.CreateUserController)
	e.DELETE("/users/:id", c.DeleteUserController)
	e.PUT("/users/:id", c.UpdateUserController)
	e.POST("/login", c.LoginUserController)

	eJwtAuth := e.Group("/jwt")
	eJwtAuth.Use(middleware.JWT([]byte(constants.JWT_SECRET_KEY)))
	eJwtAuth.GET("/users", c.GetUsersController)
	eJwtAuth.GET("/users/:id", c.GetUserController)
	eJwtAuth.DELETE("/users/:id", c.DeleteUserController)
	eJwtAuth.PUT("/users/:id", c.UpdateUserController)
	return e
}
