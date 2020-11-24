package controller

import (
	"domain"
	"fmt"
	"net/http"
)

// customer обработчик доступен только авторизованным пользователям, прошедшим аутентификацию. Контроллируется middleware Auth
func (h *Handler) customerCreate(w http.ResponseWriter, r *http.Request) {

	session, err := domain.Store.Get(r, "cookie-name")
	check(err)
	user := domain.GetUser(session)

	/*
		possibleRoles := h.connection.GetAllRoles()
		roleBook := domain.RoleBook{RoleCount: len(possibleRoles)}
		for _, value := range possibleRoles {
			roleBook.Roles = append(roleBook.Roles, *value)
		}
	*/

	var userBook domain.UserBook
	userBook.Users, err = h.connection.GetAllUsers()
	check(err)
	if r.Method == http.MethodGet {

		currentInformation := sessionInformation{user, userBook, ""}
		executeHTML("customer", "create", w, currentInformation)
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		familyName := r.FormValue("familyname")
		patronymicName := r.FormValue("patronymicname")
		user := r.FormValue("user")
		fmt.Println(user)
		newCustomer := domain.Customer{
			Name:           name,
			FamilyName:     familyName,
			PatronymicName: patronymicName,
			//User:           user,
		}

		err = h.connection.CreateCustomer(&newCustomer)
		err = session.Save(r, w)
		if err != nil {
			executeHTML("customer", "create", w, nil)
		}
		http.Redirect(w, r, "/Customer", http.StatusFound)
	}

}
