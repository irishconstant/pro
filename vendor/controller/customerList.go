package controller

import (
	"model"
	"net/http"
)

// customer обработчик доступен только авторизованным пользователям, прошедшим аутентификацию. Контроллируется middleware Auth
func (h *Handler) customer(w http.ResponseWriter, r *http.Request) { //

	session, err := model.Store.Get(r, "cookie-name")
	check(err)

	user := model.GetUser(session)
	_, err = h.connection.GetUserAttributes(&user)
	check(err)

	customers, err := h.connection.GetUserCustomers(user)
	check(err)

	customerBook := model.CustomersBook{CustomerCount: len(customers)}
	for _, value := range customers {
		customerBook.Customers = append(customerBook.Customers, *value)
	}

	currentInformation := sessionInformation{user, customerBook, ""}

	executeHTML("customer", "list", w, currentInformation)
}
