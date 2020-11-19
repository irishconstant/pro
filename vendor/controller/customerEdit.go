package controller

import (
	"fmt"
	"net/http"
)

func (h *Handler) customerEdit(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key != "" {
		fmt.Fprintln(w, r.URL.String(), "Будет отредактирован Пользователь с ИД", key)
	}
}
