package controller

import (
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
			executeHTML("reg", w, nil)
		}
	}

	if r.Method == http.MethodPost {
		login := r.FormValue("login")
		password := r.FormValue("password")
		result := h.connection.CreateUser(login, password)
		if result != true {
			executeHTML("reg", w, "Ошибка при создании пользователя")
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
