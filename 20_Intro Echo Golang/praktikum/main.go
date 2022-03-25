package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

func (u *User) Get(val User) {
	u.Id = val.Id
	u.Name = val.Name
	u.Email = val.Email
	u.Password = val.Password
}

func (u *User) Set(name, email, password string) {
	u.Name = name
	u.Email = email
	u.Password = password
}

// -------------------- controller --------------------

// GetUsersController get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// GetUserController get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	user := User{}
	id, _ := strconv.Atoi(c.Param("id"))
	for _, val := range users {
		if val.Id == id {
			user = val
			break
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get users",
		"users":    user,
	})
}

// DeleteUserController delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	for idx, val := range users {
		if val.Id == id {
			users = append(users[:idx], users[idx+1:]...)
			break
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user",
	})
}

// UpdateUserController update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	id, _ := strconv.Atoi(c.Param("id"))
	for idx, val := range users {
		if val.Id == id {
			users[idx].Set(user.Name, user.Email, user.Password)
			break
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update name",
	})
}

// CreateUserController create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start("127.0.0.1:8080"))
}
