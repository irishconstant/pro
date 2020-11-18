package controller

import (
	"model"
	"net/http"
)

// login обрабатывает попытку залогиниться
func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	session, err := model.Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {

		if h.connection.CheckPassword(r.FormValue("login"), r.FormValue("password")) == false {
			if r.FormValue("password") == "" {
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

		login := r.FormValue("login")
		/*
			roles := h.connection.GetAllRoles()
			for key, value := range roles {
				h.connection.GetRoleAbilities(value)
			}
		*/
		user := &model.User{
			Key:           login,
			Authenticated: true,
		}

		h.connection.GetUserRoles(user)
		//fmt.Println("Роль пользователя (вызов из контроллера login)", user.Role.CreateAbility)

		session.Values["user"] = user

		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/customer", http.StatusFound)
	}
}

// logout обрабывает попытку разлогиниться
func (h *Handler) logout(w http.ResponseWriter, r *http.Request) {
	session, err := model.Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["user"] = model.User{}
	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
