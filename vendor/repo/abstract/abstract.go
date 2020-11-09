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
	GetCustomers() (map[int]*model.Customer, error)
}
