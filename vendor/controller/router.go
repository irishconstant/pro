package controller

import (
	"fmt"
	"net/http"
	"repo/abstract"
)

//Router starts web-server and route between controllers
func Router(dbc abstract.DatabaseConnection) {
	h := Handler{connection: dbc}
	//http.HandleFunc("/login", h.loginController)
	//http.HandleFunc("/logout", h.logoutController)
	http.HandleFunc("/", h.mainController)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Что-то не так")
	}

}

//Handler handles something
type Handler struct {
	connection abstract.DatabaseConnection
}
