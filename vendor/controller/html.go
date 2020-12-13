package controller

import (
	"auth"
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

func makeURLWithAttributes(origin string, params map[string]string) string {

	var paramPart string

	for key, value := range params {
		if value != "" {
			paramPart = paramPart + key + "=" + value + "&"
		}
	}
	result := "/" + origin + "?" + paramPart
	return result
}

func check(err error) {
	if err != nil {
		fmt.Println("Ошибочка", err)
		log.Fatal(err)
	}
}

// sessionInformation общая структура для шаблонов html
type sessionInformation struct {
	User      auth.User
	Attribute interface{}
	Error     string
}
