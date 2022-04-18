package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"rest-clean/domain"
	"rest-clean/middleware"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UUsecase domain.UserUsecase
	UJwt     middleware.JWTService
}

func NewUserHandler(uc domain.UserUsecase, jwt middleware.JWTService) *UserHandler {
	return &UserHandler{
		UUsecase: uc,
		UJwt:     jwt,
	}
}

func (u *UserHandler) Route(e *echo.Echo) {
	e.GET("/users", u.GetAll, u.UJwt.JwtMiddleware())
	e.POST("/login", u.Login)
	e.POST("/users", u.Create, u.UJwt.JwtMiddleware())
}

func (u *UserHandler) GetAll(c echo.Context) error {
	listUs, err := u.UUsecase.GetAll()
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, listUs)
}

func (u *UserHandler) Create(c echo.Context) (err error) {
	var newUs domain.User
	err = c.Bind(&newUs)

	user, err := u.UUsecase.Create(newUs)
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, user)
}

// handler login
func (u *UserHandler) Login(c echo.Context) error {
	var user domain.User
	err := c.Bind(&user)
	var val bool
	val, err = u.UUsecase.Login(user)
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	if val == false {
		return c.JSON(http.StatusUnauthorized, ResponseError{Message: "Unauthorized"})
	}
	token, e := u.UJwt.GenerateToken()
	if e != nil {
		fmt.Println("masuk error", e)
		return c.JSON(GetStatusCode(e), ResponseError{Message: e.Error()})
	}
	return c.JSON(GetStatusCode(err), ResponseError{Message: token})
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case domain.ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
