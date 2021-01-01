package controller

import (
	"auth"
	"fmt"
	"net/http"
)

// logout обрабывает попытку разлогиниться
func (h *DecoratedHandler) logout(w http.ResponseWriter, r *http.Request) {
	session, err := auth.Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options.MaxAge = -1
	err = session.Save(r, w)

	fmt.Println("Данные в куках до логаута", session)
	user := auth.GetUser(session)
	user.Authenticated = false
	user.Comment = "Изменил состояние авторизации в логауте"

	fmt.Println("Пользователь в логауте", user)

	fmt.Println("Данные в куках после логаута", session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)

	//fmt.Println("значение куков (логаут)", session.Values)

}
