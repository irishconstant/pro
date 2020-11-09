package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"repo/abstract"

	"github.com/gorilla/mux"
)

//Router starts web-server and routes between controllers
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
	router.HandleFunc("/secret", h.secret)
	http.ListenAndServe(":8080", router)
}

//Handler handles something. I only need it for common attributes like databatase connection
type Handler struct {
	connection abstract.DatabaseConnection
}

type someAttribute interface {
}

func executeHTML(page string, w http.ResponseWriter, param someAttribute) {
	absPath, _ := filepath.Abs(fmt.Sprintf("../pro/vendor/view/%s/%s.html", page, page))
	html, err := template.ParseFiles(absPath)
	check(err)
	err = html.Execute(w, param)
	check(err)
}
