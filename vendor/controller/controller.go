package controller

import (
	"fmt"
	"log"
	"net/http"
	"repo"
)

func ViewHandler(writer http.ResponseWriter, request *http.Request) {

	connection := repo.GetConnectionParams("config.ini")
	db := connection.ConnectToDatabase()
	defer repo.CloseConnect(db)

	// Получаем пользователей
	users := repo.GetUsers(db, connection.Database)
	for key, value := range users { // Order not specified
		fmt.Println(key, *value)
	}

	for user := range users {
		message := []byte("Hello, web!")
	}
	message := []byte("Hello, web!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}
