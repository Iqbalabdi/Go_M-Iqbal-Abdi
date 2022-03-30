package main

import (
	"crud_go/config"
	"crud_go/routes"
)

func init() {
	config.InitDB()
	config.InitialMigration()
}

func main() {
	e := routes.New()
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start("127.0.0.1:8081"))
}
