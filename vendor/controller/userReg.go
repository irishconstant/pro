package controller

import (
	"core/auth"
	"fmt"
	"net/http"
	"strconv"
)

func (h *DecoratedHandler) reg(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := GetUser(session)
	possibleRoles, err := h.connection.GetAllRoles() // TODO: Изменить в будущем на возможность присваивать определенные роли в зависимости от роли авторизованного пользователя
	roleBook := RoleBook{RoleCount: len(possibleRoles)}
	for _, value := range possibleRoles {
		roleBook.Roles = append(roleBook.Roles, *value)
	}

	if r.Method == http.MethodGet {

		currentInformation := sessionInformation{User: *user, Attribute: roleBook}

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

		newUser := auth.User{
			Key:           login,
			Name:          name,
			FamilyName:    familyName,
			Password:      password,
			Role:          role,
			Authenticated: false,
		}
		err = h.connection.CreateUser(newUser)

		if err != nil {
			currentInformation := sessionInformation{Attribute: roleBook, Error: "Ошибка при создании пользователя"}
			executeHTML("user", "reg", w, currentInformation)
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
