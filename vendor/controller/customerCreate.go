package controller

import (
	"domain"
	"net/http"
)

// customer обработчик доступен только авторизованным пользователям, прошедшим аутентификацию. Контроллируется middleware Auth
func (h *Handler) customerCreate(w http.ResponseWriter, r *http.Request) {

	session, err := domain.Store.Get(r, "cookie-name")
	check(err)

	user := domain.GetUser(session)
	possibleRoles := h.connection.GetAllRoles()
	roleBook := domain.RoleBook{RoleCount: len(possibleRoles)}
	for _, value := range possibleRoles {
		roleBook.Roles = append(roleBook.Roles, *value)
	}

	if r.Method == http.MethodGet {

		currentInformation := sessionInformation{user, roleBook, ""}

		if auth := user.Authenticated; auth {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			executeHTML("customer", "create", w, currentInformation)
		}
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		familyName := r.FormValue("familyname")
		patronymicName := r.FormValue("patronymicname")

		newCustomer := domain.Customer{
			Name:           name,
			FamilyName:     familyName,
			PatronymicName: patronymicName,
		}

		newID := h.connection.CreateCustomer(newCustomer) // TODO: создание возвращает новый идентификатор

		if newID == 0 {
			// TODO: Обработать возможные ошибки
			executeHTML("customer", "create", w, nil)
		}

		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}

}
