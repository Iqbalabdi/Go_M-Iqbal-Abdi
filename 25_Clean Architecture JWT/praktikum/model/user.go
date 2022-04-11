package model

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Username string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
type User struct {
	*gorm.Model

	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserUseCase interface {
	GetAllUsers(c echo.Context) ([]User, error)
	CreateUser(c echo.Context, u *User) error
}

type UserRepository interface {
	GetAllUsers(c echo.Context) ([]User, error)
	CreateUser(c echo.Context, u *User) error
}
