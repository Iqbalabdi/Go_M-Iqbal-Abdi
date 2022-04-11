package http

import (
	"belajar-go-echo/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	UserUC model.UserUseCase
}

func NewUserHandler(e *echo.Echo, uc model.UserUseCase) {
	handler := &UserHandler{
		UserUC: uc,
	}

	e.GET("/users", handler.GetAllUsers)
	e.POST("/users", handler.CreateUser)
}

func (u *UserHandler) GetAllUsers(c echo.Context) error {

	res, err := u.UserUC.GetAllUsers(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   res,
	})
}

func (u *UserHandler) CreateUser(c echo.Context) error {

	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = u.UserUC.CreateUser(c, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "anjay error",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"users":  user,
	})
}
