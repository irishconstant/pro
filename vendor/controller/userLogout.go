package controller

import (
	"auth"
	"net/http"
)

// logout обрабывает попытку разлогиниться
func (h *DecoratedHandler) logout(w http.ResponseWriter, r *http.Request) {
	session, err := auth.Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//session.Values["SystemUser"] = nil //auth.User{}
	session.Values["Authenticated"] = false
	session.Options.MaxAge = -1

	err = session.Save(r, w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)

	//fmt.Println("значение куков (логаут)", session.Values)

}
