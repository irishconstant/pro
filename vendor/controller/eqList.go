package controller

import "net/http"

func (h *DecoratedHandler) equipment(w http.ResponseWriter, r *http.Request) { //

	executeHTML("equipment", "list", w, nil)

}
