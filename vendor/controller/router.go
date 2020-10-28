package controller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"repo/abstract"
)

//Router starts web-server and route between controllers
func Router(dbc abstract.DatabaseConnection) {

	fmt.Println("Web-сервер запущен (CTRL + C для остановки)")

	rootHandler := &Handler{name: "root", connection: dbc}
	http.Handle("/", rootHandler)

	testHandler := &Handler{name: "test", connection: dbc}
	http.Handle("/users", testHandler)

	fmt.Println("starting server at: 8080")
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)

}

//Handler handles something
type Handler struct {
	name       string
	db         *sql.DB
	database   string
	connection abstract.DatabaseConnection
}
