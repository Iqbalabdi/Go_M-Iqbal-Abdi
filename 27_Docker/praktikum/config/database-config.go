package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"rest-clean/domain"
)

func SetupDatabaseConnection() *gorm.DB {
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "",
		"DB_Host":     "3306",
		"DB_Port":     "localhost",
		"DB_Name":     "alta_section_24",
	}

	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config["DB_Username"],
			config["DB_Password"],
			config["DB_Port"],
			config["DB_Host"],
			config["DB_Name"],
		)

	DB, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&domain.User{})

	return DB
}

func CloseDatabaseConnection(db *gorm.DB) {
	err := db.Close()
	if err != nil {
		panic("failed to close database connection")
	}
}
