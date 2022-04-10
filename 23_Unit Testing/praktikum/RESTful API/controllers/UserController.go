package controllers

import (
	"crud_go/config"
	"crud_go/middleware"
	M "crud_go/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// GetUsersController get all users
func GetUsersController(c echo.Context) error {
	var users []M.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

// GetUserController get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	var user M.User

	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
	//panic("")
}

// CreateUserController create new user
func CreateUserController(c echo.Context) error {
	user := M.User{}

	err := c.Bind(&user)
	if err != nil {
		return err
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  user,
	})
}

// DeleteUserController delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	var user M.User
	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	if err := config.DB.Delete(&user, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  user,
	})
}

// UpdateUserController update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	var user M.User

	id, _ := strconv.Atoi(c.Param("id"))
	name := c.FormValue("name")

	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}
	user.Name = name
	config.DB.Save(&user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  user,
	})
}

func LoginUserController(c echo.Context) error {
	user := M.User{}

	err := c.Bind(&user)
	if err != nil {
		return err
	}

	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "failed to login",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(int(user.ID), user.Name)

	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "anjay failed to login",
			"error":   err.Error(),
		})
	}

	userResponse := M.UsersResponse{Name: user.Name, Email: user.Email, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user":    userResponse,
	})
}
