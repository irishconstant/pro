package model

import (
	"encoding/gob"

	"github.com/gorilla/sessions"
)

// User хранит информацию о Пользователях
type User struct {
	Key           string
	Password      string
	Name          string
	FamilyName    string
	Authenticated bool
	Roles         map[int]*Role
}

// Role хранит информацию о конкретной Роли
type Role struct {
	Key  int
	Name string
}

// Store хранит данные о сессиях
var Store *sessions.CookieStore

func init() {
	// Временно закомментировал. Надоело куки чистить. Для продакшн версии - раскомментировать!
	authKeyOne := []byte("1234546789012345678901234567890121234546789012345678901234567890") // securecookie.GenerateRandomKey(64)
	encryptionKeyOne := []byte("12345467890123456789012345678901")                           //securecookie.GenerateRandomKey(32)
	Store =
		sessions.NewCookieStore(
			authKeyOne,
			encryptionKeyOne,
		)

	Store.Options = &sessions.Options{
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

	gob.Register(User{})
}

//GetUser получает пользователя текущей сессии и проверяет авторизован он или нет
func GetUser(s *sessions.Session) User {
	val := s.Values["user"]
	var user = User{}
	user, ok := val.(User)
	if !ok {
		return User{Authenticated: false}
	}

	return user
}
