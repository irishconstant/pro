package main

import (
	"controller"
	"dependency"
)

func main() {

	dbc := dependency.GetDependency()
	dbc.GetConnectionParams("config.ini")
	dbc.ConnectToDatabase()
	defer dbc.CloseConnect()
	controller.Router(dbc)

	//TODO: dep init
}
