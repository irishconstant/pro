package controller

import (
	"core/contract"
	"net/http"
	"strconv"
	"time"
)

func (h *DecoratedHandler) personUpdate(w http.ResponseWriter, r *http.Request) {
	keyPerson, err := strconv.Atoi(r.URL.Query().Get("key"))

	Person, err := h.connection.GetPerson(keyPerson)

	session, err := Store.Get(r, "cookie-name")
	check(err)
	user := GetUser(session)

	if r.Method == http.MethodGet {
		Person.PossibleUsers, err = h.connection.GetAllUsers()
		check(err)
		currentInformation := sessionInformation{User: *user, Attribute: Person}
		executeHTML("person", "update", w, currentInformation)
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		familyName := r.FormValue("familyname")
		patronymicName := r.FormValue("patronymicname")
		dateBirth := r.FormValue("datebirth")
		dateDeath := r.FormValue("datedeath")
		sex, _ := strconv.ParseBool(r.FormValue("sex"))
		userLogin := r.FormValue("user")

		dateBirthG, _ := time.Parse("2006-01-02", dateBirth)
		dateDeathG, _ := time.Parse("2006-01-02", dateDeath)

		newUser, err := h.connection.GetUser(userLogin)
		newPerson := contract.Person{
			Key:            Person.Key,
			Name:           name,
			FamilyName:     familyName,
			PatronymicName: patronymicName,
			Sex:            sex,
			DateBirth:      dateBirthG,
			DateDeath:      dateDeathG,
			User:           *newUser,
		}

		err = h.connection.UpdatePerson(&newPerson)

		if err != nil {
			executeHTML("person", "update", w, nil)
		}
		http.Redirect(w, r, "/person", http.StatusFound)
	}

}
