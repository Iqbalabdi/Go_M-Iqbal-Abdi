package controllers

import (
	"crud_go/config"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	userSuccess = `{"name":"Alta","password": "123","email":"alta@gmail.com"}`
	userFail    = `{"name":"Alta","password": "","email":"alta@gmail.com"}`

	userLoginSuccess = `{"email":"alta@gmail.com","password": "123"}`
	userLoginFail    = `{"email":"alta@gmail.com","password": ""}`

	userUpdateSuccess = `{"name":"AltaBARU","password": "123","email":"alta@gmail.com"}`
	userUpdateFail    = `{"name":"AltaBARU","password": "123","email":""}`
)

func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

type (
	CustomUserValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomUserValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func TestCreateUserController(t *testing.T) {
	e := InitEchoTestAPI()

	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
		bodyRequest            string
	}{
		{
			name:                   "berhasil",
			path:                   "/users",
			expectedStatus:         http.StatusOK,
			expectedBodyStartsWith: "{\"status\":\"success\",\"users\":",
			bodyRequest:            userSuccess,
		},
	}
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(testCase.bodyRequest))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, testCase.expectedStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestLoginUserController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	e.Validator = &CustomUserValidator{validator: validator.New()}

	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
		bodyRequest            string
	}{
		{
			name:                   "berhasil",
			path:                   "/login",
			expectedStatus:         http.StatusOK,
			expectedBodyStartsWith: "{\"status\":\"success login\",\"user\":",
			bodyRequest:            userLoginSuccess,
		},
		{
			name:                   "gagal",
			path:                   "/login",
			expectedStatus:         http.StatusBadRequest,
			expectedBodyStartsWith: "{\"message\"",
			bodyRequest:            userLoginFail,
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(testCase.bodyRequest))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		if assert.NoError(t, LoginUserController(c)) {
			assert.Equal(t, testCase.expectedStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestGetUserController(t *testing.T) {
	e := InitEchoTestAPI()

	var testCases = []struct {
		name                   string
		path                   string
		param                  string
		paramValue             string
		expectedStatus         int
		expectedBodyStartsWith string
	}{
		{
			name:                   "berhasil",
			path:                   "/users/:id",
			param:                  "id",
			paramValue:             "1",
			expectedStatus:         http.StatusOK,
			expectedBodyStartsWith: "{\"status\":\"success\",\"user\":",
		},
		{
			name:                   "gagal",
			path:                   "/users/:id",
			param:                  "id",
			paramValue:             "4",
			expectedStatus:         http.StatusBadRequest,
			expectedBodyStartsWith: "{\"message\"",
		},
	}
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)
		c.SetParamNames(testCase.param)
		c.SetParamValues(testCase.paramValue)

		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, testCase.expectedStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}

}

func TestGetUsersController(t *testing.T) {
	e := InitEchoTestAPI()

	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{
		{
			name:                   "berhasil",
			path:                   "/users",
			expectedStatus:         http.StatusOK,
			expectedBodyStartsWith: "{\"status\":\"success\",\"users\":",
		},
	}
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		if assert.NoError(t, GetUsersController(c)) {
			assert.Equal(t, testCase.expectedStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestUpdateUserController(t *testing.T) {
	e := InitEchoTestAPI()

	var testCases = []struct {
		name                   string
		path                   string
		param                  string
		paramValue             string
		expectedStatus         int
		expectedBodyStartsWith string
		bodyRequest            string
	}{
		{
			name:                   "berhasil",
			path:                   "/users/:id",
			expectedStatus:         http.StatusOK,
			param:                  "id",
			paramValue:             "1",
			expectedBodyStartsWith: "{\"status\":\"success\",\"users\":",
			bodyRequest:            userUpdateSuccess,
		},
		{
			name:                   "gagal",
			path:                   "/users/:id",
			param:                  "id",
			paramValue:             "3",
			expectedStatus:         http.StatusBadRequest,
			expectedBodyStartsWith: "{\"message\"",
			bodyRequest:            userUpdateFail,
		},
	}
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(testCase.bodyRequest))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)
		c.SetParamNames(testCase.param)
		c.SetParamValues(testCase.paramValue)

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, testCase.expectedStatus, rec.Code)
			body := rec.Body.String()
			fmt.Println()
			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestDeleteUserController(t *testing.T) {
	e := InitEchoTestAPI()

	var testCases = []struct {
		name                   string
		path                   string
		param                  string
		paramValue             string
		expectedStatus         int
		expectedBodyStartsWith string
	}{
		{
			name:                   "berhasil",
			path:                   "/users/:id",
			param:                  "id",
			paramValue:             "1",
			expectedStatus:         http.StatusOK,
			expectedBodyStartsWith: "{\"status\":\"success\",\"users\":",
		},
		{
			name:                   "gagal",
			path:                   "/users/:id",
			param:                  "id",
			paramValue:             "3",
			expectedStatus:         http.StatusBadRequest,
			expectedBodyStartsWith: "{\"message\"",
		},
	}
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)
		c.SetParamNames(testCase.param)
		c.SetParamValues(testCase.paramValue)

		if assert.NoError(t, DeleteUserController(c)) {
			assert.Equal(t, testCase.expectedStatus, rec.Code)
			body := rec.Body.String()
			fmt.Println()
			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}
