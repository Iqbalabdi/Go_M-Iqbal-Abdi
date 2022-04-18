package usecase

import (
	validator "gopkg.in/go-playground/validator.v9"
	"rest-clean/domain"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(u domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: u,
	}
}

func (u *userUsecase) GetAll() (res []domain.User, err error) {
	res, err = u.userRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return
}

func (u *userUsecase) Create(user domain.User) (res domain.User, err error) {
	var ok bool
	if ok, err = u.Validate(&user); !ok {
		return res, err
	}

	res, err = u.userRepo.Create(user)
	if err != nil {
		return res, err
	}
	return
}

func (u *userUsecase) Validate(user *domain.User) (bool, error) {
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *userUsecase) Login(user domain.User) (bool, error) {
	if ok, err := u.Validate(&user); !ok {
		return false, err
	}

	userAuth, err := u.userRepo.GetByQuery(user)

	if err != nil || userAuth.Password != user.Password {
		return false, err
	}

	return true, nil
}
