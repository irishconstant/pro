package abstract

import (
	"model"
)

//DatabaseConnection обеспечивает интерфейс для соединения с СУБД (набор методов, который должен быть реализован для утиной типизации)
type DatabaseConnection interface {
	// Работа с соединением БД
	ConnectToDatabase() error
	CloseConnect() error
	GetConnectionParams(filePath string) error

	// Работа с Потребителями
	GetUserCustomersAll(u model.User) (map[int]*model.Customer, error)
	GetUserCustomersPagination(u model.User, currentPage int, pageSize int) (map[int]*model.Customer, error)

	// Подсистема авторизации и аутентификации
	CreateUser(user model.User) bool
	CheckPassword(a string, b string) bool
	GetUserRoles(user *model.User) (*model.User, error)
	GetAllRoles() map[int]*model.Role
	GetRoleAbilities(role *model.Role) (bool, error)
	GetRoleByID(int) (*model.Role, error)
	GetUserAttributes(user *model.User) (bool, error)
}
