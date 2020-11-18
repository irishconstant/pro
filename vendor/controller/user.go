package controller

import (
	"fmt"
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
	user := model.GetUser(session)
	possibleRoles := h.connection.GetAllRoles() // TODO: Изменить в будущем на возможность присваивать определенные роли в зависимости от роли авторизованного пользователя
	roleBook := model.RoleBook{RoleCount: len(possibleRoles)}
	for _, value := range possibleRoles {
		roleBook.Roles = append(roleBook.Roles, *value)
	}

	if r.Method == http.MethodGet {

		currentInformation := sessionInformation{user, roleBook, ""}

		if auth := user.Authenticated; auth {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			executeHTML("user", "reg", w, currentInformation)
		}
	}

	if r.Method == http.MethodPost {
		login := r.FormValue("login")
		fmt.Println(login)
		password := r.FormValue("password")
		name := r.FormValue("name")
		familyName := r.FormValue("familyname")
		// Из интерфейса приходит идентификатор (value)
		roleID, err := strconv.Atoi(r.FormValue("role"))
		role, err := h.connection.GetRoleByID(roleID)

		newUser := model.User{
			Key:           login,
			Name:          name,
			FamilyName:    familyName,
			Password:      password,
			Role:          role,
			Authenticated: false,
		}
		//result := h.connection.CreateUser(login, password)
		result := h.connection.CreateUserWithRoles(newUser)

		if result != true {
			errorUser := model.User{Key: login, Password: password, Name: "", FamilyName: "", Authenticated: false, Role: role}
			currentInformation := sessionInformation{errorUser, roleBook, "Ошибка при создании пользователя"}
			executeHTML("user", "reg", w, currentInformation)
		}

		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
