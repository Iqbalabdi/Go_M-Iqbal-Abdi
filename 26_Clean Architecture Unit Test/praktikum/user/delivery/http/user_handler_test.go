package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	// "strconv"
	"strings"
	"testing"
	// "time"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	// "github.com/stretchr/testify/require"
	// "github.com/labstack/echo/v4/middleware"

	"rest-clean/domain"
	"rest-clean/domain/mocks"
	jwtUsecase "rest-clean/middleware"
	mocksJwt "rest-clean/middleware/mocks"
	userHttp "rest-clean/user/delivery/http"
)

func TestLogin(t *testing.T) {
	mockUserUcase := new(mocks.UserUsecase)
	mockMidJwt := new(mocksJwt.JWTService)
	handler := userHttp.NewUserHandler(mockUserUcase, mockMidJwt)
	var key = "tokenKey"

	user := &domain.User{
		Email:    "test@test.com",
		Password: "sandi",
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		b, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
			return
		}

		mockUserUcase.On("Login", mock.Anything).Return(true, nil).Once()
		mockMidJwt.On("GenerateToken", mock.Anything).Return(key, nil).Once()
		req, err := http.NewRequest(echo.POST, "/login", strings.NewReader(string(b)))

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler.Login(c)
		body := rec.Body.String()
		assert.True(t, strings.HasPrefix(body, "{\"message\":\"tokenKey"))
	})

	t.Run("fail-wrong password", func(t *testing.T) {
		e := echo.New()
		b, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
			return
		}
		mockUserUcase.On("Login", mock.Anything).Return(false, nil).Once()
		mockMidJwt.On("GenerateToken", mock.Anything).Return(nil, nil).Once()
		req, err := http.NewRequest(echo.POST, "/login", strings.NewReader(string(b)))

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler.Login(c)
		code := rec.Code
		assert.Equal(t, http.StatusUnauthorized, code)
	})

	t.Run("fail-user server error", func(t *testing.T) {
		e := echo.New()
		b, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
			return
		}

		mockUserUcase.On("Login", mock.Anything).Return(false, domain.ErrInternalServerError).Once()
		mockMidJwt.On("GenerateToken", mock.Anything).Return(nil, nil).Once()
		req, err := http.NewRequest(echo.POST, "/login", strings.NewReader(string(b)))

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler.Login(c)
		code := rec.Code
		assert.Equal(t, http.StatusInternalServerError, code)
	})

	t.Run("fail-user not found", func(t *testing.T) {
		e := echo.New()
		b, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
			return
		}

		mockUserUcase.On("Login", mock.Anything).Return(false, domain.ErrNotFound).Once()
		mockMidJwt.On("GenerateToken", mock.Anything).Return(nil, nil).Once()
		req, err := http.NewRequest(echo.POST, "/login", strings.NewReader(string(b)))

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler.Login(c)
		code := rec.Code
		assert.Equal(t, code, http.StatusNotFound)
	})

	t.Run("fail-generate token", func(t *testing.T) {
		// memperbarui session
		mockUserUcase := new(mocks.UserUsecase)
		mockMidJwt := new(mocksJwt.JWTService)
		handler := userHttp.NewUserHandler(mockUserUcase, mockMidJwt)

		e := echo.New()
		b, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
			return
		}

		mockUserUcase.On("Login", mock.Anything).Return(true, nil).Once()
		mockMidJwt.On("GenerateToken", mock.Anything).Return(key, domain.ErrInternalServerError).Once()
		req, err := http.NewRequest(echo.POST, "/login", strings.NewReader(string(b)))

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler.Login(c)
		code := rec.Code
		assert.Equal(t, code, http.StatusInternalServerError)
	})
}

func TestCreate(t *testing.T) {
	mockUserUcase := new(mocks.UserUsecase)
	mockMidJwt := new(mocksJwt.JWTService)
	handler := userHttp.NewUserHandler(mockUserUcase, mockMidJwt)

	user := domain.User{
		Email:    "test@test.com",
		Password: "sandi",
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		b, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
			return
		}

		mockUserUcase.On("Create", mock.Anything).Return(user, nil).Once()
		req, err := http.NewRequest(echo.POST, "/users", strings.NewReader(string(b)))

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler.Create(c)
		body := rec.Body.String()
		assert.True(t, strings.HasPrefix(body, "{\"email\":"))
	})

	t.Run("fail", func(t *testing.T) {
		e := echo.New()
		b, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
			return
		}
		user = domain.User{}
		mockUserUcase.On("Create", mock.Anything).Return(user, domain.ErrInternalServerError).Once()
		req, err := http.NewRequest(echo.POST, "/users", strings.NewReader(string(b)))

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler.Create(c)
		code := rec.Code
		assert.Equal(t, http.StatusInternalServerError, code)
	})
}

func TestGetAll(t *testing.T) {
	mockUserUcase := new(mocks.UserUsecase)
	mockMidJwt := new(mocksJwt.JWTService)
	handler := userHttp.NewUserHandler(mockUserUcase, mockMidJwt)

	user := []domain.User{
		domain.User{
			Email:    "test@test.com",
			Password: "sandi",
		},
		domain.User{
			Email:    "asdad@tdasdasest.com",
			Password: "sadasdasdndi",
		},
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		b, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
			return
		}

		mockUserUcase.On("GetAll", mock.Anything).Return(user, nil).Once()
		req, err := http.NewRequest(echo.GET, "/users", strings.NewReader(string(b)))

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler.GetAll(c)
		code := rec.Code
		assert.Equal(t, code, http.StatusOK)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		b, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
			return
		}
		user = []domain.User{}
		mockUserUcase.On("GetAll", mock.Anything).Return(user, domain.ErrInternalServerError).Once()
		req, err := http.NewRequest(echo.GET, "/users", strings.NewReader(string(b)))

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler.GetAll(c)
		code := rec.Code
		assert.Equal(t, http.StatusInternalServerError, code)
	})
}

func TestRoute(t *testing.T) {
	mockUserUcase := new(mocks.UserUsecase)
	handlerJwt := jwtUsecase.NewJwtService("secreats")
	handler := userHttp.NewUserHandler(mockUserUcase, handlerJwt)

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		handler.Route(e)
	})
}
