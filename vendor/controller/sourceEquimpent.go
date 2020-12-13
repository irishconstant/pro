package controller

import "net/http"

func (h *DecoratedHandler) sourceEquipment(w http.ResponseWriter, r *http.Request) { //

	executeHTML("source", "scheme", w, nil)

}
