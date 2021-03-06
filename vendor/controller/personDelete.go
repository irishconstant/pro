package controller

import (
	"core/contract"
	"net/http"
	"strconv"
)

func (h *DecoratedHandler) personDelete(w http.ResponseWriter, r *http.Request) {
	keyPerson, err := strconv.Atoi(r.URL.Query().Get("key"))

	Person, err := h.connection.GetPerson(keyPerson)
	session, err := Store.Get(r, "cookie-name")
	check(err)
	user := GetUser(session)
	//err = h.connection.GetUserAttributes(&user)
	check(err)

	if r.Method == http.MethodGet {
		Person.PossibleUsers, err = h.connection.GetAllUsers()
		check(err)
		currentInformation := sessionInformation{User: *user, Attribute: Person}
		executeHTML("Person", "update", w, currentInformation)
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		familyName := r.FormValue("familyname")
		patronymicName := r.FormValue("patronymicname")
		userLogin := r.FormValue("user")

		user, err := h.connection.GetUser(userLogin)
		newPerson := contract.Person{
			Key:            Person.Key,
			Name:           name,
			FamilyName:     familyName,
			PatronymicName: patronymicName,
			User:           *user,
		}

		err = h.connection.UpdatePerson(&newPerson)

		if err != nil {
			executeHTML("person", "update", w, nil)
		}
		http.Redirect(w, r, "/person", http.StatusFound)
	}

}
