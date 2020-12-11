package abstract

import (
	"domain/auth"
	"domain/contract"
)

//DatabaseConnection обеспечивает интерфейс для соединения с СУБД (набор методов, который должен быть реализован для утиной типизации)
type DatabaseConnection interface {
	// Работа с соединением БД
	ConnectToDatabase() error
	CloseConnect() error
	GetConnectionParams(filePath string) error

	// Работа со справочными данными
	GetContactType(int) (*contract.ContactType, error)
	GetAllContactTypes() ([]*contract.ContactType, error)
	GetCitizenship(int) (*contract.Citizenship, error)
	GetAllCitizenship() ([]*contract.Citizenship, error)

	// Работа с Потребителями
	GetUserFiltredPersonsPagination(auth.User, int, int, int, string, string, string, string) (map[int]*contract.Person, error)
	GetUserFiltredResultsQuantity(auth.User, int, int, int, string, string, string, string) (int, error)
	CreatePerson(*contract.Person) error
	GetPerson(int) (*contract.Person, error)
	UpdatePerson(*contract.Person) error
	DeletePerson(*contract.Person) error

	// Подсистема авторизации и аутентификации
	CreateUser(auth.User) error
	GetUser(string) (*auth.User, error)
	GetAllUsers() ([]auth.User, error)
	CheckPassword(string, b string) bool
	GetUserRoles(*auth.User) error
	GetAllRoles() (map[int]*auth.Role, error)
	GetRoleAbilities(*auth.Role) error
	GetRoleByID(int) (*auth.Role, error)
	GetUserAttributes(*auth.User) error
}
