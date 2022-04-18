package domain

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	*gorm.Model
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserUsecase interface {
	GetAll() ([]User, error)
	Create(user User) (User, error)
	Validate(user *User) (bool, error)
	Login(user User) (bool, error)
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetByQuery(user User) (User, error)
	Create(user User) (User, error)
}
