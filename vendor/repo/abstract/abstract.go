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
	GetUsers() (map[int]*model.User, error)
	//Авторизация является функцией определения прав доступа к ресурсам и управления этим доступом.
	//Autorisation(user string, password string) string

	//Аутентификация (от греческого: αυθεντικός ; реальный или подлинный): подтверждение подлинности чего-либо или кого либо.
	//Например, предъявление паспорта - это подтверждение подлинности заявленного имени отчества.
	//Authentication(user string, password string) string
}
