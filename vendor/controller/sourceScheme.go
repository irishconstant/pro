package controller

import "net/http"

func (h *DecoratedHandler) sourceScheme(w http.ResponseWriter, r *http.Request) { //

	executeHTML("source", "scheme", w, nil)

}
