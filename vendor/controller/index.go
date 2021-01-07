package controller

import (
	"net/http"
)

//index обрабывате запросы к стартовой странице
func (h *DecoratedHandler) index(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := GetUser(session)
	/*
		if user.Authenticated {
			h.connection.GetUserAttributes(&user)
		}
	*/

	executeHTML("index", "index", w, *user)
}
