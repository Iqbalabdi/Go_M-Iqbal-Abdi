package mysql_test

import (
	"errors"
	"regexp"
	"testing"
	"time"
	// "database/sql/driver"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"rest-clean/domain"
	userMysqlRepo "rest-clean/user/repository/mysql"
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	gormDB, err := gorm.Open("mysql", db)

	if err != nil {
		t.Fatalf("Error '%s' terjadi ketika membuat koneksi ke stub database", err)
	}

	mockUsers := []domain.User{
		domain.User{
			Email: "a@a.com", Password: "123456",
		},
		domain.User{
			Email: "b@b.com", Password: "123456",
		},
	}
	timeIs := time.Now()
	rows := sqlmock.NewRows([]string{"id", "email", "password", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, mockUsers[0].Email, mockUsers[0].Password, timeIs, timeIs, nil).
		AddRow(2, mockUsers[1].Email, mockUsers[1].Password, timeIs, timeIs, nil)

	t.Run("success", func(t *testing.T) {
		query := regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL")

		mock.ExpectQuery(query).WillReturnRows(rows)
		userRepo := userMysqlRepo.NewMysqlUserRepository(gormDB)
		list, err := userRepo.GetAll()
		assert.Empty(t, err)
		assert.Len(t, list, 2)
	})

	t.Run("fail", func(t *testing.T) {
		query := regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL")

		mock.ExpectQuery(query).WillReturnError(errors.New("GetAll - error unit test"))
		userRepo := userMysqlRepo.NewMysqlUserRepository(gormDB)
		_, err := userRepo.GetAll()
		assert.Error(t, err)
	})
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	gormDB, err := gorm.Open("mysql", db)

	if err != nil {
		t.Fatalf("Error '%s' terjadi ketika membuat koneksi ke stub database", err)
	}

	us := &domain.User{
		Email:    "test@test.com",
		Password: "password",
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()

		mock.ExpectExec("INSERT INTO `users`").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, us.Email, us.Password).WillReturnResult(sqlmock.NewResult(0, 1))

		mock.ExpectCommit()

		userRepo := userMysqlRepo.NewMysqlUserRepository(gormDB)

		list, err := userRepo.Create(*us)
		assert.Empty(t, err)

		assert.NotEmpty(t, list)
	})

	t.Run("fail", func(t *testing.T) {
		mock.ExpectBegin()

		mock.ExpectExec("INSERT INTO `users`").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, us.Email, us.Password).WillReturnError(errors.New("Create - error unit test"))

		mock.ExpectCommit()

		userRepo := userMysqlRepo.NewMysqlUserRepository(gormDB)

		_, err := userRepo.Create(*us)
		assert.Error(t, err)
	})
}

func TestGetByQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	gormDB, err := gorm.Open("mysql", db)
	timeIs := time.Now()

	if err != nil {
		t.Fatalf("Error '%s' terjadi ketika membuat koneksi ke stub database", err)
	}

	us := &domain.User{
		Email:    "test@test.com",
		Password: "password",
	}

	rows := sqlmock.NewRows([]string{"id", "email", "password", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, us.Email, us.Password, timeIs, timeIs, nil)

	t.Run("success", func(t *testing.T) {
		query := regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND ((email = ?)) ORDER BY `users`.`id` ASC LIMIT 1")
		mock.ExpectQuery(query).WithArgs(us.Email).WillReturnRows(rows)

		userRepo := userMysqlRepo.NewMysqlUserRepository(gormDB)
		user, err := userRepo.GetByQuery(*us)
		assert.NoError(t, err)
		assert.NotEmpty(t, user)
	})

	t.Run("fail", func(t *testing.T) {
		query := regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND ((email = ?)) ORDER BY `users`.`id` ASC LIMIT 1")
		mock.ExpectQuery(query).WithArgs(us.Email).WillReturnError(errors.New("GetByQuery - error unit test"))

		userRepo := userMysqlRepo.NewMysqlUserRepository(gormDB)
		_, err := userRepo.GetByQuery(*us)
		assert.Error(t, err)
	})
}
