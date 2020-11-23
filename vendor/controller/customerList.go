package controller

import (
	"math"
	"model"
	"net/http"
	"strconv"
)

// customer обработчик доступен только авторизованным пользователям, прошедшим аутентификацию. Контроллируется middleware Auth
func (h *Handler) customer(w http.ResponseWriter, r *http.Request) { //

	// Получаем страницы из параметров
	key := r.URL.Query().Get("page")
	var page int
	if key != "" {
		page, _ = strconv.Atoi(key)
	} else {
		page = 1
	}

	session, err := model.Store.Get(r, "cookie-name")
	check(err)

	user := model.GetUser(session)
	_, err = h.connection.GetUserAttributes(&user)
	check(err)

	customers, err := h.connection.GetUserCustomersAll(user)
	check(err)

	customerBook := model.CustomersBook{CustomerCount: len(customers)}
	// Если необходима пагинация
	if customerBook.CustomerCount > h.pageSize {
		customersPerPage, err := h.connection.GetUserCustomersPagination(user, page)
		check(err)
		for _, value := range customersPerPage {
			customerBook.Customers = append(customerBook.Customers, *value)
		}
		customerBook.Pages = customerBook.MakeRange(1, int(math.Ceil(float64(customerBook.CustomerCount)/float64(h.pageSize)))) // Округляем число страниц в большую сторону
		currentInformation := sessionInformation{user, customerBook, ""}
		executeHTML("customer", "list", w, currentInformation)

	} else {
		for _, value := range customers {
			customerBook.Customers = append(customerBook.Customers, *value)
		}

		currentInformation := sessionInformation{user, customerBook, ""}

		executeHTML("customer", "list", w, currentInformation)
	}
}
