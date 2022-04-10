package database

import (
	"crud_go/config"
	"crud_go/models"
	"github.com/labstack/echo/v4"
)

func GetBooks() (interface{}, error) {
	var books []models.Book
	if e := config.DB.Find(&books).Error; e != nil {
		return nil, e
	}
	return books, nil
}

func GetBookById(id int) (interface{}, error) {
	var book models.Book
	if e := config.DB.First(&book, id).Error; e != nil {
		return nil, e
	}
	return book, nil
}

func CreateBook(bookData echo.Context) (interface{}, error) {
	var book models.Book
	err := bookData.Bind(&book)
	if err != nil {
		return nil, err
	}
	if e := bookData.Validate(book); e != nil {
		return nil, e
	}
	if e := config.DB.Save(&book).Error; e != nil {
		return nil, e
	}
	return book, nil
}

func UpdateBookById(id int, bookData echo.Context) (interface{}, error) {
	var book models.Book
	var bookUpdate models.Book
	err := bookData.Bind(&bookUpdate)
	if err != nil {
		return nil, err
	}
	if e := bookData.Validate(bookUpdate); e != nil {
		return nil, e
	}
	config.DB.First(&book, id)
	var e error
	if result := config.DB.Model(&book).Where("id=?", id).Updates(&bookUpdate); result == nil {
		return nil, e
	}
	return book, nil
}

func DeleteBookById(id int) (interface{}, error) {
	var book models.Book
	var e error
	if err := config.DB.First(&book).Error; err != nil {
		return nil, err
	}

	if result := config.DB.Delete(&book, id); result == nil {
		return nil, e
	}
	return "book deleted", nil
}
