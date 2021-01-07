package controller

import "core/auth"

// sessionInformation общая структура для шаблонов html
type sessionInformation struct {
	User auth.User

	// Возможные расширения для сессии
	Attribute    interface{}
	AttributeMap map[interface{}]interface{}

	Error string
}
