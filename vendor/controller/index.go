package controller

import (
	"model"
	"net/http"
)

func (h *Handler) index(w http.ResponseWriter, r *http.Request) {

	session, err := model.Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := model.GetUser(session)
	executeHTML("index", w, user)

}
