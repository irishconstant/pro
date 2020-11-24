package domain

// User хранит информацию о Пользователях
type User struct {
	Key           string
	Password      string
	Name          string
	FamilyName    string
	Authenticated bool
	Role          *Role
}

// Role хранит информацию о конкретной Роли
type Role struct {
	Key           int
	Name          string
	AdminOnly     bool
	CreateAbility map[int]*Area // Создание нового в системном разделе и запуск операции для создания чего-либо
	ReadAbility   map[int]*Area // Чтение и запуск операции, не производящей изменения
	UpdateAbility map[int]*Area // Изменение данных и запуск операции, производящей изменения (но без удаления)
	DeleteAbility map[int]*Area // Удаление данных и запуск операции, производящей удаление
}

// RoleBook набор Ролей для интерфейса
type RoleBook struct {
	RoleCount int
	Roles     []Role
}

// Area представляет из себя Подсистему (инкапсулированную)
type Area struct {
	Key  int
	Name string
	Type int // 1 - Системный раздел 2 - Операция
}
