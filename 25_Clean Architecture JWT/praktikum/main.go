package main

import (
	"belajar-go-echo/config"
	_userHttpDelivery "belajar-go-echo/user/controller/http"
	_userRepo "belajar-go-echo/user/repository/sqlite"
	_userUcase "belajar-go-echo/user/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = config.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	app := echo.New()
	userRepo := _userRepo.NewSqliteUserRepository(db)
	userUCase := _userUcase.NewUserUseCase(userRepo)

	_userHttpDelivery.NewUserHandler(app, userUCase)

	app.Logger.Fatal(app.Start("127.0.0.1:9091"))
}
