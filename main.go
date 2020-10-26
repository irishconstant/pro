package main

import (
	"dependency"
	"fmt"
)

func main() {
	dbc := dependency.GetDependency()
	connectionParams := dbc.GetConnectionParams("config.ini")
	db := dbc.ConnectToDatabase(connectionParams.ConnectionString)
	defer dbc.CloseConnect(db)

	// Получаем пользователей
	users := dbc.GetUsers(db, connectionParams.Database)
	for key, value := range users { // Order not specified
		fmt.Println(key, *value)
	}

}
