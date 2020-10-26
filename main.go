package main

import (
	"dependency"
	"fmt"
	abstract "repo/abstract"

	"github.com/golobby/container"
)

func main() {
	dependency.GetDependency()
	var dbc abstract.DatabaseConnection
	container.Make(&dbc)
	a := dbc.GetConnectionParams("config.ini")
	db := dbc.ConnectToDatabase(a.ConnectionString)
	defer dbc.CloseConnect(db)

	// Получаем пользователей
	users := dbc.GetUsers(db, a.Database)
	for key, value := range users { // Order not specified
		fmt.Println(key, *value)
	}

}
