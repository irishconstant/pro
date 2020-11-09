package model

import (
	"encoding/gob"
	"text/template"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

// User holds a users account information
type User struct {
	Username      string
	Authenticated bool
}

// Store will hold all session data
var Store *sessions.CookieStore

// Tpl holds all parsed templates
var Tpl *template.Template

func init() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	Store = sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	Store.Options = &sessions.Options{
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

	gob.Register(User{})

	Tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

//GetUser does
func GetUser(s *sessions.Session) User {
	val := s.Values["user"]
	var user = User{}
	user, ok := val.(User)
	if !ok {
		return User{Authenticated: false}
	}
	return user
}
