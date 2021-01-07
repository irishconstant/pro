package controller

import (
	"core/auth"
	"net/http"
)

// login обрабатывает попытку залогиниться
func (h *DecoratedHandler) login(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		login := r.FormValue("login")
		password := r.FormValue("password")
		if h.connection.CheckPassword(login, password) == false {
			if password == "" {
				session.AddFlash("Необходимо ввести пароль")
			}
			session.AddFlash("Неправильное имя пользователя или пароль")
			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/forbidden", http.StatusFound)
			return
		}

		user := &auth.User{
			Key:           login,
			Authenticated: true,
			Comment:       "Создан на этапе логина и не изменялся",
		}
		h.connection.GetUserRoles(user)
		err = h.connection.GetUserAttributes(user)

		session.Values["SystemUser"] = user

		err = session.Save(r, w)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/source", http.StatusFound)
	}
}
