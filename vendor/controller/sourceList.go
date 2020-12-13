package controller

import "net/http"

/* Фактические параметры работы котельных заполняются для каждого расчётного периода.
При этом для каждого расчётного периода определяется: интерфейс из двух частей.
Первая часть «Шапка», вторая – детальные данные по конкретным теплоисточникам.
*/
func (h *DecoratedHandler) source(w http.ResponseWriter, r *http.Request) { //

	executeHTML("source", "list", w, nil)

}
