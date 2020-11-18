package controller

import (
	"fmt"
	"html/template"
	"model"
	"net/http"
	"path/filepath"
	"repo/abstract"

	"github.com/gorilla/mux"
)

//Router запускает web-сервер и настраивает маршрутизацию
func Router(dbc abstract.DatabaseConnection) {
	staticDir := "/static/"
	h := Handler{connection: dbc}
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
	api.HandleFunc("/customer", h.customer)
	api.Path("/customer").Handler(http.HandlerFunc(h.customer))

	http.ListenAndServe(":8080", router)
	//corsOrigins := handlers.AllowedOrigins([]string{"*"}) // TODO: для работы с AJAX
	// handlers.CORS(corsOrigins)(router))
}

//Handler тип мне нужен для того, чтобы было что-то общее у всех обработчиков
type Handler struct {
	connection abstract.DatabaseConnection
}

// executeHTML инкапсулирует работу с шаблонами и генерацию html
func executeHTML(folder string, page string, w http.ResponseWriter, param interface{}) {
	absPath, _ := filepath.Abs(fmt.Sprintf("../pro/vendor/view/%s/%s.html", folder, page))
	html, err := template.ParseFiles(absPath)
	check(err)
	err = html.Execute(w, param)
	check(err)
}

// sessionInformation общая структура для шаблонов html
type sessionInformation struct {
	User      model.User
	Attribute interface{}
	Error     string
}

// authMiddleware выполняется для проверки аутентифицирован ли пользователь. TODO: сделать доступ к определенным разделам по ролям
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := model.Store.Get(r, "cookie-name")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user := model.GetUser(session)

		if user.Authenticated == false {
			session.AddFlash("Доступ запрещён (пройдите авторизацию и аутентификацию)!")
			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				http.Redirect(w, r, "/forbidden", http.StatusFound)
				return
			}
			http.Redirect(w, r, "/forbidden", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}
