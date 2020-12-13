package controller

import "net/http"

func (h *DecoratedHandler) sourceDelete(w http.ResponseWriter, r *http.Request) { //

	executeHTML("source", "edit", w, nil)

}
