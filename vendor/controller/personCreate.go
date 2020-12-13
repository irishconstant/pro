package controller

import (
	"auth"
	"core/contract"
	"net/http"
	"strconv"
)

// PersonCreate обработчик доступен только авторизованным пользователям, прошедшим аутентификацию. Контроллируется middleware Auth
func (h *DecoratedHandler) personCreate(w http.ResponseWriter, r *http.Request) {
	// Работа с куками
	session, err := auth.Store.Get(r, "cookie-name")
	check(err)
	user := auth.GetUser(session)
	err = h.connection.GetUserAttributes(&user)
	check(err)

	var userBook UserBook
	userBook.Users, err = h.connection.GetAllUsers()
	check(err)

	if r.Method == http.MethodGet {
		currentInformation := sessionInformation{user, userBook, ""}
		executeHTML("person", "create", w, currentInformation)
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		familyName := r.FormValue("familyname")
		patronymicName := r.FormValue("patronymicname")
		sex, err := strconv.ParseBool(r.FormValue("sex"))

		userID := r.FormValue("user")

		User, err := h.connection.GetUser(userID)

		newPerson := contract.Person{
			Name:           name,
			FamilyName:     familyName,
			PatronymicName: patronymicName,
			Sex:            sex,
			User:           *User,
		}

		err = h.connection.CreatePerson(&newPerson)
		if err != nil {
			executeHTML("person", "create", w, nil)
		}
		http.Redirect(w, r, "/person", http.StatusFound)

	}

}
