package main

import (
	"controller"
	"dependency"
	"fmt"
)

func main() {
	// Работаем с подключениями
	dbc := dependency.GetDependency()
	connectionParams := dbc.GetConnectionParams("config.ini")
	db := dbc.ConnectToDatabase(connectionParams.ConnectionString)
	defer dbc.CloseConnect(db)
	users := dbc.GetUsers(db, connectionParams.Database)
	for key, value := range users { // Order not specified
		fmt.Println(key, *value)
	}

	// Запускаем веб-сервер
	controller.WebAppMain(dbc, connectionParams.Database)

}
