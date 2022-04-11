package usecase

import (
	"belajar-go-echo/model"
	"github.com/labstack/echo/v4"
)

type userUseCase struct {
	userRepo model.UserRepository
}

func NewUserUseCase(u model.UserRepository) model.UserUseCase {
	return &userUseCase{
		userRepo: u,
	}
}

func (u *userUseCase) GetAllUsers(c echo.Context) (res []model.User, err error) {
	res, err = u.userRepo.GetAllUsers(c)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *userUseCase) CreateUser(c echo.Context, m *model.User) (err error) {
	err = u.userRepo.CreateUser(c, m)
	if err != nil {
		return err
	}

	return
}
