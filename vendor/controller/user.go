package controller

import (
	"fmt"
	"model"
	"net/http"
)

func (h *Handler) reg(w http.ResponseWriter, r *http.Request) {
	session, err := model.Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodGet {

		user := model.GetUser(session)

		if auth := user.Authenticated; auth {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			executeHTML("user", "reg", w, nil)
		}
	}

	if r.Method == http.MethodPost {
		login := r.FormValue("login")
		password := r.FormValue("password")
		name := r.FormValue("name")
		roles := r.FormValue("role")
		fmt.Println(roles)
		for key, value := range roles {
			fmt.Println(name, key, value)
		}

		result := h.connection.CreateUser(login, password)
		if result != true {
			executeHTML("user", "reg", w, "Ошибка при создании пользователя")
		}

		user := &model.User{
			Key:           login,
			Authenticated: true,
		}

		session.Values["user"] = user

		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/customer", http.StatusFound)
	}
}
