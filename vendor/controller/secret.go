package controller

import (
	"fmt"
	"model"
	"net/http"
)

// secret displays the secret message for authorized users
func (h *Handler) secret(w http.ResponseWriter, r *http.Request) {
	session, err := model.Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := model.GetUser(session)

	if auth := user.Authenticated; !auth {
		session.AddFlash("Доступ запрещён!")
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
	for _, value := range customers { // Порядок вывода случайный
		//	fmt.Println(key, *value)
		customerBook.Customers = append(customerBook.Customers, fmt.Sprintf("Имя: %s, Фамилия: %s, Отчество: %s", value.Name, value.FamilyName, value.PatronymicName))
	}

	executeHTML("user", w, customerBook)
	//model.Tpl.ExecuteTemplate(w, "secret.gohtml", user.Username)
}
