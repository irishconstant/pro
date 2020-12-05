package controller

import (
	"net/http"
	"repo/abstract"

	"github.com/gorilla/mux"
)

//Router запускает web-сервер и настраивает маршрутизацию
func Router(dbc abstract.DatabaseConnection) {

	staticDir := "/static/"
	h := DecoratedHandler{connection: dbc, pageSize: 7}
	router := mux.NewRouter()
	// Обработка статичных файлов
	router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	// Те, кто НЕ попадают под middleware проверку аутентификации
	router.Path("/forbidden").Handler(http.HandlerFunc(h.forbidden))
	router.Path("/user/reg").Handler(http.HandlerFunc(h.reg))
	router.Path("/user/edit").Handler(http.HandlerFunc(h.forbidden))
	router.Path("/login").Handler(http.HandlerFunc(h.login))
	router.Path("/logout").Handler(http.HandlerFunc(h.logout))
	router.Path("/").Handler(http.HandlerFunc(h.index))

	// Те, кто попадают под middleware проверку аутентификации
	api := router.PathPrefix("/").Subrouter()
	api.Use(authMiddleware)
	api.HandleFunc("/customer/", h.customer)
	api.HandleFunc("/customer", h.customer)
	api.HandleFunc("/customer/update", h.customerUpdate)
	api.HandleFunc("/customer/create", h.customerCreate)
	api.HandleFunc("/customer/delete", h.customerDelete)
	http.ListenAndServe(":8080", router)
	//corsOrigins := handlers.AllowedOrigins([]string{"*"}) // TODO: для работы с AJAX
	// handlers.CORS(corsOrigins)(router))

}

//DecoratedHandler реализует методы обработки (контроля) маршрутизации
type DecoratedHandler struct {
	connection abstract.DatabaseConnection
	pageSize   int //Максимальное количество записей на странице
}
