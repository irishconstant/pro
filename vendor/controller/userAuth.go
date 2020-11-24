package controller

import (
	"domain"
	"net/http"
)

// login обрабатывает попытку залогиниться
func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	session, err := domain.Store.Get(r, "cookie-name")
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
		user := &domain.User{
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
	session, err := domain.Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["user"] = domain.User{}
	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

// sessionInformation общая структура для шаблонов html
type sessionInformation struct {
	User      domain.User
	Attribute interface{}
	Error     string
}

// authMiddleware выполняется для проверки аутентифицирован ли пользователь. TODO: сделать доступ к определенным разделам по ролям
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := domain.Store.Get(r, "cookie-name")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user := domain.GetUser(session)

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
