package main

import (
	"controller"
	"dependency"
)

func main() {

	// Работаем с подключениями
	dbc := dependency.GetDependency()

	dbc.GetConnectionParams("config.ini")
	dbc.ConnectToDatabase()
	defer dbc.CloseConnect()

	// Работаем с веб-сервером
	controller.Router(dbc)

	//dep init
}
