package controller

import (
	"fmt"
	"log"
	"net/http"
)

func (h *Handler) customer(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, r.URL.String())
	// Использование параметров
	myParam := r.URL.Query().Get("param")
	if myParam != "" {
		fmt.Fprintln(w, "my Param is", myParam)
	}
	key := r.FormValue("key")
	if key != "" {
		fmt.Fprintln(w, "key is", key)
	}
}

func check(err error) {
	if err != nil {
		fmt.Println("Ошибочка", err)
		log.Fatal(err)
	}
}
