package routes

import (
	"crud_go/constants"
	c "crud_go/controllers"
	mid "crud_go/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	mid.LogMiddleware(e)
	// Route / to handler function

	// Users
	e.GET("/users", c.GetUsersController)
	e.GET("/users/:id", c.GetUserController)
	e.POST("/users", c.CreateUserController)
	e.DELETE("/users/:id", c.DeleteUserController)
	e.PUT("/users/:id", c.UpdateUserController)
	e.POST("/login", c.LoginUserController)

	// books
	e.GET("/books", c.GetBooksController)
	e.GET("/books/:id", c.GetBookController)
	e.POST("/books", c.CreateBookController)
	e.DELETE("/books/:id", c.DeleteBookController)
	e.PUT("/books/:id", c.UpdateBookController)

	eJwtAuth := e.Group("/jwt")
	eJwtAuth.Use(middleware.JWT([]byte(constants.JWT_SECRET_KEY)))
	eJwtAuth.GET("/users", c.GetUsersController)
	eJwtAuth.GET("/users/:id", c.GetUserController)
	eJwtAuth.DELETE("/users/:id", c.DeleteUserController)
	eJwtAuth.PUT("/users/:id", c.UpdateUserController)
	return e
}
