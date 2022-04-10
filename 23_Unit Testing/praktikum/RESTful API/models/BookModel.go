package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title" form:"title"  validate:"required"`
	Author string `json:"author" form:"author"  validate:"required"`
	Year   int    `json:"year" form:"year"  validate:"required"`
	Token  string `json:"token" form:"token"`
}
