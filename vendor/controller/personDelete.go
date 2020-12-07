package controller

import (
	"domain"
	"net/http"
	"strconv"
)

func (h *DecoratedHandler) PersonDelete(w http.ResponseWriter, r *http.Request) {
	keyPerson, err := strconv.Atoi(r.URL.Query().Get("key"))

	Person, err := h.connection.GetPerson(keyPerson)
	session, err := domain.Store.Get(r, "cookie-name")
	check(err)
	user := domain.GetUser(session)
	err = h.connection.GetUserAttributes(&user)
	check(err)

	if r.Method == http.MethodGet {
		Person.PossibleUsers, err = h.connection.GetAllUsers()
		check(err)
		currentInformation := sessionInformation{user, Person, ""}
		executeHTML("Person", "update", w, currentInformation)
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		familyName := r.FormValue("familyname")
		patronymicName := r.FormValue("patronymicname")
		userLogin := r.FormValue("user")

		user, err := h.connection.GetUser(userLogin)
		newPerson := domain.Person{
			Key:            Person.Key,
			Name:           name,
			FamilyName:     familyName,
			PatronymicName: patronymicName,
			User:           *user,
		}

		err = h.connection.UpdatePerson(&newPerson)

		err = session.Save(r, w)
		if err != nil {
			executeHTML("Person", "update", w, nil)
		}
		http.Redirect(w, r, "/Person", http.StatusFound)
	}

}
