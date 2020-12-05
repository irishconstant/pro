package domain

import (
	"encoding/gob"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

// Store хранит данные о сессиях
var Store *sessions.CookieStore

// Инициализирует новый Store c данными сессий
func init() {
	// Временно закомментировал. Надоело куки чистить. Для продакшн версии - раскомментировать!
	authKeyOne := securecookie.GenerateRandomKey(64)       //[]byte("1234546789012345678901234567890121234546789012345678901234567890") // securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32) // []byte("12345467890123456789012345678901")                           //securecookie.GenerateRandomKey(32)
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
