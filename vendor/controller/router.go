package controller

import (
	"fmt"
	"html/template"
	"log"
	"model"
	"net/http"
	"path/filepath"
	"repo/abstract"

	"github.com/gorilla/mux"
)

//Router запускае web-сервер и настраивает маршрутизацию
func Router(dbc abstract.DatabaseConnection) {
	staticDir := "/static/"
	h := Handler{connection: dbc}
	router := mux.NewRouter()
	router.PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
	router.HandleFunc("/", h.index)
	router.HandleFunc("/login", h.login)
	router.HandleFunc("/logout", h.logout)
	router.HandleFunc("/forbidden", h.forbidden)
	router.HandleFunc("/customer", h.customer)
	router.HandleFunc("/reg", h.reg)
	//router.Use(authMiddleware)

	//corsOrigins := handlers.AllowedOrigins([]string{"*"}) // Для работы с AJAX

	http.ListenAndServe(":8080", router) // handlers.CORS(corsOrigins)(router))
}

//Handler тип мне нужен для того, чтобы было общее соединение с БД
type Handler struct {
	connection abstract.DatabaseConnection
}

func executeHTML(page string, w http.ResponseWriter, param interface{}) {
	absPath, _ := filepath.Abs(fmt.Sprintf("../pro/vendor/view/%s/%s.html", page, page))
	html, err := template.ParseFiles(absPath)
	check(err)
	err = html.Execute(w, param)
	check(err)
}

type sessionInformation struct {
	User      model.User
	Attribute interface{}
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Промежуточный слой", r.RequestURI, r.Method)
		next.ServeHTTP(w, r)
	})
}
