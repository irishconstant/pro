package controller

import "net/http"

func (h *DecoratedHandler) sourceUpdate(w http.ResponseWriter, r *http.Request) { //

	executeHTML("source", "update", w, nil)

}
