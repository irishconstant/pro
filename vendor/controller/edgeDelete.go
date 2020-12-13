package controller

import "net/http"

func (h *DecoratedHandler) equipmentDelete(w http.ResponseWriter, r *http.Request) { //

	executeHTML("equipment", "delete", w, nil)

}
