package mysql

import (
	"github.com/jinzhu/gorm"
	"rest-clean/domain"
)

type mysqlUserRepository struct {
	connection *gorm.DB
}

func NewMysqlUserRepository(db *gorm.DB) domain.UserRepository {
	return &mysqlUserRepository{
		connection: db,
	}
}

func (m *mysqlUserRepository) GetAll() (res []domain.User, err error) {
	if err = m.connection.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (m *mysqlUserRepository) Create(user domain.User) (res domain.User, err error) {
	if err = m.connection.Save(&user).Error; err != nil {
		return res, err
	}
	return user, nil
}

func (m *mysqlUserRepository) GetByQuery(user domain.User) (res domain.User, err error) {
	var userAuth domain.User
	if err = m.connection.Where("email = ?", user.Email).First(&userAuth).Error; err != nil {
		return
	}
	return userAuth, nil
}
