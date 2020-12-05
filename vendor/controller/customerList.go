package controller

import (
	"domain"
	"fmt"
	"math"
	"net/http"
	"strconv"
)

// customer обработчик доступен только авторизованным пользователям, прошедшим аутентификацию. Контроллируется middleware Auth
func (h *DecoratedHandler) customer(w http.ResponseWriter, r *http.Request) { //

	if r.Method == http.MethodPost {
		params := make(map[string]string)
		params["name"] = r.FormValue("name")
		params["familyname"] = r.FormValue("familyname")
		params["patrname"] = r.FormValue("patrname")
		params["sex"] = r.FormValue("sex")
		filteredAddress := makeURLWithAttributes("customer", params)
		http.Redirect(w, r, filteredAddress, http.StatusFound)
	}

	// Получаем страницу из параметров
	//fmt.Println(r.URL)
	key := r.URL.Query().Get("page")
	var page int
	if key != "" {
		page, _ = strconv.Atoi(key)
	} else {
		page = 1
	}

	// Получаем параметры фильтрации
	name := r.URL.Query().Get("name")
	familyName := r.URL.Query().Get("familyname")
	patrName := r.URL.Query().Get("patrname")
	sex := r.URL.Query().Get("sex")

	session, err := domain.Store.Get(r, "cookie-name")
	check(err)

	user := domain.GetUser(session)
	err = h.connection.GetUserAttributes(&user)
	check(err)

	customers, err := h.connection.GetUserFiltredCustomersPagination(user, 0, page, h.pageSize, name, familyName, patrName, sex)
	check(err)
	customerBook := domain.CustomersBook{CustomerCount: len(customers)}

	// Если необходима пагинация
	if customerBook.CustomerCount > h.pageSize {
		customersPerPage, err := h.connection.GetUserFiltredCustomersPagination(user, 1, page, h.pageSize, name, familyName, patrName, sex) //h.connection.GetUserCustomersPagination(user, page, h.pageSize)
		check(err)
		for _, value := range customersPerPage {
			customerBook.Customers = append(customerBook.Customers, *value)
		}
		customerBook.CurrentPage = page

		// Создаем страницы для показа (1, одна слева от текущей, одна справа от текущей, последняя)
		if name != "" {
			name = "&name=" + name
		}
		if familyName != "" {
			familyName = "&familyname=" + familyName
		}
		if patrName != "" {
			patrName = "&patrname=" + patrName
		}
		if sex != "" {
			sex = "&sex=" + sex
		}
		customerBook.Pages = domain.MakePages(1, int(math.Ceil(float64(customerBook.CustomerCount)/float64(h.pageSize))), page)
		for key := range customerBook.Pages {
			customerBook.Pages[key].URL = fmt.Sprintf("/customer?%s%s%s%s", name, familyName, patrName, sex)
		}
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
