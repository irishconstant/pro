package controller

import "net/http"

func (h *DecoratedHandler) equipmentUpdate(w http.ResponseWriter, r *http.Request) { //

	executeHTML("equipment", "update", w, nil)
}
