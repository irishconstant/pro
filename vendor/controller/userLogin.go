package controller

import (
	"auth"
	"net/http"
)

// login обрабатывает попытку залогиниться
func (h *DecoratedHandler) login(w http.ResponseWriter, r *http.Request) {
	session, err := auth.Store.Get(r, "cookie-name")
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
		}
		h.connection.GetUserRoles(user)

		session.Values["SystemUser"] = user

		//	fmt.Println("значение куков (логин)", session.Values)
		//		session.Values["authenticated"] = true
		err = session.Save(r, w)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/source", http.StatusFound)
	}
}
