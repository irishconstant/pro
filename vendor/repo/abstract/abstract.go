package abstract

import (
	"database/sql"
	"domain"
)

//SQLConnect parameters
type SQLConnect struct {
	//comment
	ConnectionString string
	Database         string
}

//DatabaseConnection provides an interface for database connection
type DatabaseConnection interface {
	ConnectToDatabase(connectionString string) *sql.DB
	CloseConnect(db *sql.DB)
	GetConnectionParams(filePath string) SQLConnect
	GetUsers(db *sql.DB, dbname string) map[int]*domain.User
	//Авторизация является функцией определения прав доступа к ресурсам и управления этим доступом.
	//Autorisation(user string, password string) string

	//Аутентификация (от греческого: αυθεντικός ; реальный или подлинный): подтверждение подлинности чего-либо или кого либо.
	//Например, предъявление паспорта - это подтверждение подлинности заявленного имени отчества.
	//Authentication(user string, password string) string
}
