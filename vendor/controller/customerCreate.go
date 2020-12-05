package controller

import (
	"domain"
	"net/http"
)

// customer обработчик доступен только авторизованным пользователям, прошедшим аутентификацию. Контроллируется middleware Auth
func (h *DecoratedHandler) customerCreate(w http.ResponseWriter, r *http.Request) {
	// Работа с куками
	session, err := domain.Store.Get(r, "cookie-name")
	check(err)
	user := domain.GetUser(session)
	err = h.connection.GetUserAttributes(&user)
	check(err)

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

		userID := r.FormValue("user")

		User, err := h.connection.GetUser(userID)

		newCustomer := domain.Customer{
			Name:           name,
			FamilyName:     familyName,
			PatronymicName: patronymicName,
			User:           *User,
		}

		err = h.connection.CreateCustomer(&newCustomer)
		if err != nil {
			executeHTML("customer", "create", w, nil)
		}
		http.Redirect(w, r, "/customer", http.StatusFound)

	}

}
