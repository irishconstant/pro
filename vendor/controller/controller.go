package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"repo/abstract"
)

//WebAppMain starts web-app
func WebAppMain(db abstract.DatabaseConnection, dbname string) {
	fmt.Println("Web-сервер запущен (CTRL + C для остановки)")
	http.HandleFunc("/hello", helloHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	// Получаем представление
	// TODO: Написать отдельную функцию, выдающую абсолютный путь чисто по наименованию (без всякой шняги типа vendor и т.п.)
	absPath, _ := filepath.Abs("../pro/vendor/view/hello.html")
	html, err := template.ParseFiles(absPath)
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
