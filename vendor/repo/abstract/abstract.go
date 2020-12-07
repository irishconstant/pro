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

	// Работа со справочными данными
	GetContactType(int) (*domain.ContactType, error)
	GetAllContactTypes() ([]*domain.ContactType, error)
	GetCitizenship(int) (*domain.Citizenship, error)
	GetAllCitizenship() ([]*domain.Citizenship, error)

	// Работа с Потребителями
	GetUserFiltredPersonsPagination(domain.User, int, int, int, string, string, string, string) (map[int]*domain.Person, error)
	CreatePerson(*domain.Person) error
	GetPerson(int) (*domain.Person, error)
	UpdatePerson(*domain.Person) error
	DeletePerson(*domain.Person) error

	// Подсистема авторизации и аутентификации
	CreateUser(domain.User) error
	GetUser(string) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)
	CheckPassword(string, b string) bool
	GetUserRoles(*domain.User) error
	GetAllRoles() (map[int]*domain.Role, error)
	GetRoleAbilities(*domain.Role) error
	GetRoleByID(int) (*domain.Role, error)
	GetUserAttributes(*domain.User) error
}
