package controller

import (
	"domain"
	"net/http"
)

// forbidden обрабатывает попытку получить доступ туда, куда нельзя
func (h *Handler) forbidden(w http.ResponseWriter, r *http.Request) {
	session, err := domain.Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	flashMessages := session.Flashes()
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	executeHTML("index", "forbidden", w, flashMessages)
}
