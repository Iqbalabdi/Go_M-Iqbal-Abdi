package usecase_test

import (
	"rest-clean/domain"
	"rest-clean/domain/mocks"

	ucase "rest-clean/user/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"errors"
	"testing"
)

func TestCreate(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUserValid := domain.User{
		Email:    "terbang",
		Password: "mengudara",
	}
	mockUserInvalid := domain.User{
		Email:    "",
		Password: "mengudara",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Create", mock.Anything).Return(mockUserValid, nil).Once()
		usecase := ucase.NewUserUseCase(mockUserRepo)
		_, err := usecase.Create(mockUserValid)

		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		mockUserRepo.On("Create", mock.Anything).Return(mockUserValid, errors.New("error unit test")).Once()
		usecase := ucase.NewUserUseCase(mockUserRepo)
		_, err := usecase.Create(mockUserValid)

		assert.Error(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		mockUserRepo.On("Create", mock.Anything).Return(mockUserInvalid, errors.New("error unit test")).Once()
		usecase := ucase.NewUserUseCase(mockUserRepo)
		_, err := usecase.Create(mockUserInvalid)

		assert.Error(t, err)
	})
}

func TestLogin(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := domain.User{
		Email:    "terbang",
		Password: "telentang",
	}
	mockUserValid := domain.User{
		Email:    "terbang",
		Password: "mengudara",
	}
	mockUserInvalid := domain.User{
		Email:    "",
		Password: "mengudara",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetByQuery", mock.Anything).Return(mockUserValid, nil).Once()
		mockUserRepo.On("Login", mock.Anything).Return(true, nil).Once()
		usecase := ucase.NewUserUseCase(mockUserRepo)
		_, err := usecase.Login(mockUserValid)

		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		mockUserRepo.On("GetByQuery", mock.Anything).Return(mockUser, nil).Once()
		mockUserRepo.On("Login", mock.Anything).Return(false, nil).Once()
		usecase := ucase.NewUserUseCase(mockUserRepo)
		_, err := usecase.Login(mockUserValid)

		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		mockUserRepo.On("GetByQuery", mock.Anything).Return(mockUserValid, errors.New("error unit test")).Once()
		mockUserRepo.On("Login", mock.Anything).Return(false, errors.New("error unit test")).Once()
		usecase := ucase.NewUserUseCase(mockUserRepo)
		_, err := usecase.Login(mockUserValid)

		assert.Error(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		mockUserRepo.On("GetByQuery", mock.Anything).Return(mockUserInvalid, errors.New("error unit test")).Once()
		mockUserRepo.On("Login", mock.Anything).Return(false, errors.New("error unit test")).Once()
		usecase := ucase.NewUserUseCase(mockUserRepo)
		_, err := usecase.Login(mockUserInvalid)

		assert.Error(t, err)
	})

}

func TestGetAll(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := domain.User{
		Email:    "terbang",
		Password: "telentang",
	}

	mockListUser := make([]domain.User, 0)
	mockListUser = append(mockListUser, mockUser)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetAll", mock.Anything).Return(mockListUser, nil).Once()
		usecase := ucase.NewUserUseCase(mockUserRepo)
		_, err := usecase.GetAll()

		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		mockUserRepo.On("GetAll", mock.Anything).Return(mockListUser, errors.New("error unit test")).Once()
		usecase := ucase.NewUserUseCase(mockUserRepo)
		_, err := usecase.GetAll()

		assert.Error(t, err)
	})
}
