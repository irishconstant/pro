package auth

import "domain/sys"

// User хранит информацию о Пользователях
type User struct {
	Key           string
	Password      string
	Name          string
	FamilyName    string
	Authenticated bool
	Role          *Role
}

// UserBook хранит информацию наборе Пользователей
type UserBook struct {
	UserCount   int
	CurrentPage int
	Pages       []sys.Page
	Users       []User
}

// Role хранит информацию о конкретной Роли
type Role struct {
	Key           int
	Name          string
	AdminOnly     bool
	CreateAbility map[int]*sys.Area // Создание нового в системном разделе и запуск операции для создания чего-либо
	ReadAbility   map[int]*sys.Area // Чтение и запуск операции, не производящей изменения
	UpdateAbility map[int]*sys.Area // Изменение данных и запуск операции, производящей изменения (но без удаления)
	DeleteAbility map[int]*sys.Area // Удаление данных и запуск операции, производящей удаление
}

// RoleBook набор Ролей для интерфейса
type RoleBook struct {
	RoleCount int
	Roles     []Role
}
