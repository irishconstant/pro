package controller

import (
	"fmt"
	"html/template"
	"log"
	"model"
	"net/http"
	"path/filepath"
)

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	users, err := h.connection.GetUsers()
	check(err)
	userBook := model.Userbook{UserCount: len(users)}
	for _, value := range users { // Порядок вывода случайный
		//	fmt.Println(key, *value)
		userBook.Users = append(userBook.Users, fmt.Sprintf("Имя: %s, Фамилия: %s, Отчество: %s", value.Name, value.FamilyName, value.PatronymicName))
	}
	absPath, _ := filepath.Abs("../pro/vendor/view/user/users.html")
	html, err := template.ParseFiles(absPath)
	check(err)
	err = html.Execute(w, userBook)
	check(err)

	fmt.Fprintln(w, "Name: ", h.name, "URL: ", r.URL.String())

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
		log.Fatal(err)
	}
}
