package main

import (
	"fmt"
	"repo"
)

func main() {

	connection := repo.GetConnectionParams("config.ini")
	db := connection.ConnectToDatabase()
	defer repo.CloseConnect(db)

	// Получаем пользователей
	users := repo.GetUsers(db, connection.Database)
	for key, value := range users { // Order not specified
		fmt.Println(key, *value)
	}
}
