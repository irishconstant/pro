package controller

import (
	"fmt"
	"html/template"
	"log"
	"model"
	"net/http"
	"path/filepath"
)

func (h *Handler) customer(w http.ResponseWriter, r *http.Request) {

	customers, err := h.connection.GetCustomers()
	check(err)
	customerBook := model.CustomersBook{CustomerCount: len(customers)}
	for _, value := range customers { // Порядок вывода случайный
		//	fmt.Println(key, *value)
		customerBook.Customers = append(customerBook.Customers, fmt.Sprintf("Имя: %s, Фамилия: %s, Отчество: %s", value.Name, value.FamilyName, value.PatronymicName))
	}
	absPath, _ := filepath.Abs("../pro/vendor/view/customer/customers.html")
	html, err := template.ParseFiles(absPath)
	check(err)
	err = html.Execute(w, customerBook)
	check(err)

	fmt.Fprintln(w, r.URL.String())
	// Использование параметров
	myParam := r.URL.Query().Get("param")
	if myParam != "" {
		fmt.Fprintln(w, "my Param is", myParam)
	}
	key := r.FormValue("key")
	if key != "" {
		fmt.Fprintln(w, "key is", key)
	}
}

func check(err error) {
	if err != nil {
		fmt.Println("Ошибочка", err)
		log.Fatal(err)
	}
}
