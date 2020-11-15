package abstract

import (
	"model"
)

//SQLConnect параметры для соединения с СУБД
type SQLConnect struct {
	//comment
	ConnectionString string
	Database         string
}

//DatabaseConnection обеспечивает интерфейс для соединения с СУБД (набор методов, который должен быть реализован для утиной типизации)
type DatabaseConnection interface {
	// Работа с соединением БД
	ConnectToDatabase() error
	CloseConnect() error
	GetConnectionParams(filePath string) error

	// Работа с Потребителями
	GetUserCustomers(u model.User) (map[int]*model.Customer, error)

	// Подсистема авторизации и аутентификации
	CreateUser(login string, password string) bool
	CheckPassword(a string, b string) bool
	GetUserRoles(user *model.User) (*model.User, error)
}
