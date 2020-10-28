package controller

import (
	"fmt"
	"html/template"
	"log"
	"model"
	"net/http"
	"path/filepath"
	"time"
)

func (h *Handler) mainController(w http.ResponseWriter, r *http.Request) {
	// Проверяем сессию в куках
	//fmt.Println(r.Cookies)
	//	session, err := r.Cookie("session_id")
	// fmt.Println(session)
	// check(err)
	// Если не устарела, то открываем основную форму с указанием логина и роли
	// Если устарела, то открываем форму для логина

	/*
		loggedIn := (err != http.ErrNoCookie)
		if loggedIn {
			fmt.Fprintln(w, `<a href="/logout">logout</a>`)
			fmt.Fprintln(w, "Welcome, "+session.Value)
		} else {
			fmt.Fprintln(w, `<a href="/login">login</a>`)
			fmt.Fprintln(w, "You need to login")
		}
	*/
	if r.Method != http.MethodPost {
		absPath, _ := filepath.Abs("../pro/vendor/view/login/login.html")
		html, err := template.ParseFiles(absPath)
		check(err)
		err = html.Execute(w, nil)
		return
	}
	inputLogin := r.FormValue("login")
	fmt.Fprintln(w, "you enter: ", inputLogin)
}

func (h *Handler) loginController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		absPath, _ := filepath.Abs("../pro/vendor/view/login/login.html")
		html, err := template.ParseFiles(absPath)
		check(err)
		err = html.Execute(w, nil)
		return
	}
	/*
		expiration := time.Now().Add(10 * time.Hour)
		cookie := http.Cookie{
			Name:    "session_id",
			Value:   "rvasily",
			Expires: expiration,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	*/
	inputLogin := r.FormValue("login")
	fmt.Fprintln(w, "you enter: ", inputLogin)
}

func (h *Handler) logoutController(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	users, err := h.connection.GetUsers()
	check(err)
	userBook := model.Userbook{UserCount: len(users)}
	for _, value := range users { // Порядок вывода случайный
		//	fmt.Println(key, *value)
		userBook.Users = append(userBook.Users, fmt.Sprintf("Имя: %s, Фамилия: %s, Отчество: %s", value.Name, value.FamilyName, value.PatronymicName))
	}
	absPath, _ := filepath.Abs("../pro/vendor/view/user/users.html")
	html, err := template.ParseFiles(absPath)
	check(err)
	err = html.Execute(w, userBook)
	check(err)

	fmt.Fprintln(w, r.URL.String())

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
	fmt.Println("Ошибочка")
	if err != nil {
		log.Fatal(err)
	}
}
