package controller

import (
	"model"
	"net/http"
)

func (h *Handler) customer(w http.ResponseWriter, r *http.Request) {
	session, err := model.Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := model.GetUser(session)

	if auth := user.Authenticated; !auth {
		session.AddFlash("Доступ запрещён (пройдите авторизацию и аутентификацию)!")
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/forbidden", http.StatusFound)
		return
	}

	customers, err := h.connection.GetCustomers()
	check(err)
	customerBook := model.CustomersBook{CustomerCount: len(customers)}
	for _, value := range customers {
		customerBook.Customers = append(customerBook.Customers, *value)
	}

	executeHTML("customer", w, customerBook)
}
