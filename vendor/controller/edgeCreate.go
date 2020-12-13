package controller

import "net/http"

func (h *DecoratedHandler) equipmentCreate(w http.ResponseWriter, r *http.Request) { //

	executeHTML("equipment", "create", w, nil)

}
