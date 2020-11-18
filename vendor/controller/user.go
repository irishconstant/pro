package controller

import (
	"model"
	"net/http"
	"strconv"
)

func (h *Handler) reg(w http.ResponseWriter, r *http.Request) {
	session, err := model.Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		user := model.GetUser(session)
		possibleRoles := h.connection.GetAllRoles()
		roleBook := model.RoleBook{RoleCount: len(possibleRoles)}
		for _, value := range possibleRoles {
			roleBook.Roles = append(roleBook.Roles, *value)
		}

		currentInformation := sessionInformation{user, roleBook, ""}

		//fmt.Println(possibleRoles)
		if auth := user.Authenticated; auth {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			executeHTML("user", "reg", w, currentInformation)
		}
	}

	if r.Method == http.MethodPost {
		login := r.FormValue("login")
		password := r.FormValue("password")
		name := r.FormValue("name")
		// Из интерфейса приходит идентификатор (value)
		roleID, err := strconv.Atoi(r.FormValue("role"))

		result := h.connection.CreateUser(login, password)
		role, err := h.connection.GetRoleByID(roleID)

		if result != true {
			errorUser := model.User{Key: login, Password: password, Name: "", FamilyName: "", Authenticated: false, Role: role}
			currentInformation := sessionInformation{errorUser, nil, "Ошибка при создании пользователя"}
			executeHTML("user", "reg", w, currentInformation)
		}

		user := &model.User{
			Key:           login,
			Name:          name,
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
