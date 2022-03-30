package Controllers

import (
	"crud_go/app/Models"
	"crud_go/config"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// GetUsersController get all users
func GetUsersController(c echo.Context) error {
	var users []Models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// GetUserController get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	var user Models.User

	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "get user by id",
		"user":    user,
	})
	//panic("")
}

// CreateUserController create new user
func CreateUserController(c echo.Context) error {
	user := Models.User{}

	err := c.Bind(&user)
	if err != nil {
		return err
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// DeleteUserController delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	var user Models.User
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Delete(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user with id : " + string(rune(id)),
	})
}

// UpdateUserController update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	var user Models.User

	id, _ := strconv.Atoi(c.Param("id"))
	name := c.FormValue("name")

	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user.Name = name
	config.DB.Save(&user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user with id : " + string(rune(id)),
	})
}
