package controller

import (
	"auth"
	"net/http"
	"strings"
)

// authMiddleware выполняется для проверки аутентифицирован ли пользователь. TODO: сделать доступ к определенным разделам по ролям
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := auth.Store.Get(r, "cookie-name")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user := auth.GetUser(session)

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

func caselessMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.ToLower(r.URL.Path)
		//	log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
