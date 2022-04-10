package controllers

import (
	database "crud_go/lib"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetBooksController(c echo.Context) error {
	books, e := database.GetBooks()
	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  books,
	})
}

func GetBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book, e := database.GetBookById(id)
	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  book,
	})
}

func CreateBookController(c echo.Context) error {
	result, e := database.CreateBook(c)
	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  result,
	})
}

func DeleteBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, e := database.DeleteBookById(id)
	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  result,
	})
}

func UpdateBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, e := database.UpdateBookById(id, c)
	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  result,
	})
}
