package abstract

import (
	"model"
)

//SQLConnect parameters
type SQLConnect struct {
	//comment
	ConnectionString string
	Database         string
}

//DatabaseConnection provides an interface for database connection
type DatabaseConnection interface {
	ConnectToDatabase() error
	CloseConnect() error
	GetConnectionParams(filePath string) error

	// Работа с Потребителями
	GetCustomers() (map[int]*model.Customer, error)

	// Подсистема авторизации и аутентификации
	CreateUser(login string, password string) bool
	CheckPassword(a string, b string) bool
}
