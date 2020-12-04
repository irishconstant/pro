package controller

import (
	"domain"
	"net/http"
)

//index обрабывате запросы к стартовой странице
func (h *Handler) index(w http.ResponseWriter, r *http.Request) {
	session, err := domain.Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := domain.GetUser(session)
	if user.Authenticated {
		h.connection.GetUserAttributes(&user)
	}

	executeHTML("index", "index", w, user)

}
