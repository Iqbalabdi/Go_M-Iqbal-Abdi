package sqlite

import (
	"belajar-go-echo/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type sqliteUserRepository struct {
	db *gorm.DB
}

func NewSqliteUserRepository(db *gorm.DB) model.UserRepository {
	return &sqliteUserRepository{db}
}

func (s *sqliteUserRepository) GetAllUsers(c echo.Context) (res []model.User, err error) {

	res = make([]model.User, 0)
	err = s.db.Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *sqliteUserRepository) CreateUser(c echo.Context, m *model.User) (err error) {

	if err := c.Bind(m); err != nil {
		return err
	}

	err = s.db.Create(m).Error
	if err != nil {
		return err
	}

	return
}
