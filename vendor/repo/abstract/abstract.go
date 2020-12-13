package abstract

import (
	"auth"
	"core/contract"
	"core/ref"
)

//DatabaseConnection обеспечивает интерфейс для соединения с СУБД
type DatabaseConnection interface {
	// Работа с соединением БД
	ConnectToDatabase() error
	CloseConnect() error
	GetConnectionParams(filePath string) error

	// Работа со справочниками
	GetContactType(int) (*contract.ContactType, error)
	GetAllContactTypes() ([]*contract.ContactType, error)
	GetCitizenship(int) (*contract.Citizenship, error)
	GetAllCitizenship() ([]*contract.Citizenship, error)
	GetSeasonMode(int) (*ref.SeasonMode, error)
	GetFuelType(int) (*ref.FuelType, error)

	// Работа с Техническими данными

	// Работа с Юридическими лицами

	// Работа с Физическими лицами
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
