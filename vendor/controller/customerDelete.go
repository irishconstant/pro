package controller

import (
	"domain"
	"net/http"
	"strconv"
)

func (h *DecoratedHandler) customerDelete(w http.ResponseWriter, r *http.Request) {
	keyCustomer, err := strconv.Atoi(r.URL.Query().Get("key"))

	customer, err := h.connection.GetCustomer(keyCustomer)
	session, err := domain.Store.Get(r, "cookie-name")
	check(err)
	user := domain.GetUser(session)
	err = h.connection.GetUserAttributes(&user)
	check(err)

	if r.Method == http.MethodGet {
		customer.PossibleUsers, err = h.connection.GetAllUsers()
		check(err)
		currentInformation := sessionInformation{user, customer, ""}
		executeHTML("customer", "update", w, currentInformation)
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		familyName := r.FormValue("familyname")
		patronymicName := r.FormValue("patronymicname")
		userLogin := r.FormValue("user")

		user, err := h.connection.GetUser(userLogin)
		newCustomer := domain.Customer{
			Key:            customer.Key,
			Name:           name,
			FamilyName:     familyName,
			PatronymicName: patronymicName,
			User:           *user,
		}

		err = h.connection.UpdateCustomer(&newCustomer)

		err = session.Save(r, w)
		if err != nil {
			executeHTML("customer", "update", w, nil)
		}
		http.Redirect(w, r, "/customer", http.StatusFound)
	}

}
