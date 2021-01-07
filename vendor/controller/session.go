package controller

import (
	"core/auth"
	"encoding/gob"

	"github.com/gorilla/sessions"
)

// Store хранит данные о сессиях
var Store *sessions.CookieStore

// Инициализирует новый Store c данными сессий
func init() {
	// Временно закомментировал. Надоело куки чистить. Для продакшн версии - раскомментировать!
	authKeyOne := []byte("1234546789012345678901234567890121234546789012345678901234567890") // securecookie.GenerateRandomKey(64)
	encryptionKeyOne := []byte("12345467890123456789012345678901")                           // securecookie.GenerateRandomKey(32)

	Store =
		sessions.NewCookieStore(
			authKeyOne,
			encryptionKeyOne,
		)

	Store.Options = &sessions.Options{
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

	// Регаем, что в куках может быть сохранён указатель на Пользователя
	gob.Register(&auth.User{})
}

//GetUser возвращает указатель на пользователя из текущей сессии в куках
func GetUser(s *sessions.Session) *auth.User {
	// Получаем указатель из сессии
	val := s.Values["SystemUser"]
	// Создаем переменную с типом указатель на класс Пользователя
	var user = &auth.User{}
	// Присваиваем этой переменной приведенное значение из сессии
	user, ok := val.(*auth.User)
	// Если не получилось привести, то значит это не пользователь, а значит что аутентифицированного пользователя нет
	if !ok {
		user = &auth.User{Authenticated: false}
		return user
	}
	// Иначе возвращаем пользователя
	return user
}
