package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"rest-clean/config"

	_jwtUsecase "rest-clean/middleware"

	_userHttpDelivery "rest-clean/user/delivery/http"
	_userRepo "rest-clean/user/repository/mysql"
	_userUcase "rest-clean/user/usecase"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

	// inisiasi jwtusecase
	jwtUsecase = _jwtUsecase.NewJwtService("secret")

	userRepository = _userRepo.NewMysqlUserRepository(db)
	userUsecase    = _userUcase.NewUserUseCase(userRepository)
	userHandler    = _userHttpDelivery.NewUserHandler(userUsecase, jwtUsecase)
)

func main() {
	e := echo.New()
	userHandler.Route(e)
	e.Logger.Fatal(e.Start("localhost:9090"))
}
