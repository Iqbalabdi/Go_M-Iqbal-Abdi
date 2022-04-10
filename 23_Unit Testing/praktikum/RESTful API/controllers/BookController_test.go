package controllers

import (
	"crud_go/config"
	"github.com/go-playground/validator/v10"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	// "fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	bookSuccess = `{"title":"Pemrograman jaringan","author": "Kadarsyah","year":2022}`
	bookFail    = `{"title":"Pemrograman jaringan","author": "","year":}`

	bookUpdateSuccess = `{"title":"Pemrograman jaringan Revisi","author": "Kadarsyah","year":2025}`
	bookUpdateFail    = `{"title":"Pemrograman jaringan Revisi","author": "Kadarsyah","year":}`
)

type (
	CustomBookValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomBookValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func TestCreateBookController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	e.Validator = &CustomBookValidator{validator: validator.New()}

	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
		bodyRequest            string
	}{
		{
			name:                   "berhasil",
			path:                   "/books",
			expectedStatus:         http.StatusOK,
			expectedBodyStartsWith: "{\"books\":",
			bodyRequest:            bookSuccess,
		},
		{
			name:                   "gagal",
			path:                   "/books",
			expectedStatus:         http.StatusBadRequest,
			expectedBodyStartsWith: "{\"message\":",
			bodyRequest:            bookFail,
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(testCase.bodyRequest))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		if assert.NoError(t, CreateBookController(c)) {
			assert.Equal(t, testCase.expectedStatus, rec.Code)
			body := rec.Body.String()
			// fmt.Println(body)
			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestGetBooksController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	e.Validator = &CustomBookValidator{validator: validator.New()}

	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{
		{
			name:                   "berhasil",
			path:                   "/books",
			expectedStatus:         http.StatusOK,
			expectedBodyStartsWith: "{\"books\":",
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		if assert.NoError(t, GetBooksController(c)) {
			assert.Equal(t, testCase.expectedStatus, rec.Code)
			body := rec.Body.String()
			// fmt.Println(body)
			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}
func TestGetBookDetailControllers(t *testing.T) {

	config.InitDB()
	e := echo.New()
	e.Validator = &CustomBookValidator{validator: validator.New()}

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
			path:                   "/books/:id",
			param:                  "id",
			paramValue:             "1",
			expectedStatus:         http.StatusOK,
			expectedBodyStartsWith: "{\"books\":",
		},
		{
			name:                   "gagal",
			path:                   "/books/:id",
			param:                  "id",
			paramValue:             "dsadsa",
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

		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, testCase.expectedStatus, rec.Code)
			body := rec.Body.String()
			// fmt.Println(body)
			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}
func TestUpdateBookController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	e.Validator = &CustomBookValidator{validator: validator.New()}

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
			path:                   "/books/:id",
			param:                  "id",
			paramValue:             "1",
			expectedStatus:         http.StatusOK,
			expectedBodyStartsWith: "{\"books\":",
			bodyRequest:            bookUpdateSuccess,
		},
		{
			name:                   "gagal",
			path:                   "/books/:id",
			param:                  "id",
			paramValue:             "3",
			expectedStatus:         http.StatusBadRequest,
			expectedBodyStartsWith: "{\"message\"",
			bodyRequest:            bookUpdateFail,
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

		if assert.NoError(t, UpdateBookController(c)) {
			assert.Equal(t, testCase.expectedStatus, rec.Code)
			body := rec.Body.String()
			// fmt.Println(body)
			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestDeleteBookController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	e.Validator = &CustomBookValidator{validator: validator.New()}

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
			path:                   "/books/:id",
			param:                  "id",
			paramValue:             "1",
			expectedStatus:         http.StatusOK,
			expectedBodyStartsWith: "{\"books\":",
		},
		{
			name:                   "gagal",
			path:                   "/books/:id",
			param:                  "id",
			paramValue:             "3",
			expectedStatus:         http.StatusBadRequest,
			expectedBodyStartsWith: "{\"message\"",
		},
	}
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)
		c.SetParamNames(testCase.param)
		c.SetParamValues(testCase.paramValue)

		if assert.NoError(t, DeleteBookController(c)) {
			assert.Equal(t, testCase.expectedStatus, rec.Code)
			body := rec.Body.String()
			// fmt.Println(body)
			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}
