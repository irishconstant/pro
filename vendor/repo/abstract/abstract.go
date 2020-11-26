package abstract

import (
	"domain"
)

//DatabaseConnection обеспечивает интерфейс для соединения с СУБД (набор методов, который должен быть реализован для утиной типизации)
type DatabaseConnection interface {
	// Работа с соединением БД
	ConnectToDatabase() error
	CloseConnect() error
	GetConnectionParams(filePath string) error

	// Работа с Потребителями
	GetUserCustomersAll(u domain.User) (map[int]*domain.Customer, error)
	GetUserCustomersPagination(u domain.User, currentPage int, pageSize int) (map[int]*domain.Customer, error)
	CreateCustomer(u *domain.Customer) error

	// Подсистема авторизации и аутентификации
	CreateUser(user domain.User) error
	GetUser(login string) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)
	CheckPassword(a string, b string) bool
	GetUserRoles(user *domain.User) error
	GetAllRoles() (map[int]*domain.Role, error)
	GetRoleAbilities(role *domain.Role) error
	GetRoleByID(int) (*domain.Role, error)
	GetUserAttributes(user *domain.User) error
}
