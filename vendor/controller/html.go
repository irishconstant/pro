package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)


// executeHTML инкапсулирует работу с шаблонами и генерацию html
func executeHTML(folder string, page string, w http.ResponseWriter, param interface{}) {
	absPath, _ := filepath.Abs(fmt.Sprintf("../pro/vendor/view/%s/%s.html", folder, page))
	html, err := template.ParseFiles(absPath)
	check(err)
	err = html.Execute(w, param)
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println("Ошибочка", err)
		log.Fatal(err)
	}
}
