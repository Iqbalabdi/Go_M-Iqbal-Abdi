package http

import (
	"belajar-go-echo/model"
	"belajar-go-echo/user/controller/http/middleware"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type UserHandler struct {
	UserUC model.UserUseCase
}

func NewUserHandler(e *echo.Echo, uc model.UserUseCase) {
	handler := &UserHandler{
		UserUC: uc,
	}

	e.GET("/users", handler.GetAllUsers, middleware.JWTMiddleware())
	e.POST("/users", handler.CreateUser)
	e.POST("/login", handler.Login)
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

func (u *UserHandler) Login(c echo.Context) error {
	var request model.LoginRequest

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if request.Username != "admin" && request.Password != "admin" {
		return c.JSON(http.StatusUnauthorized, "username atau password salah")
	}

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, jwtString)
}
