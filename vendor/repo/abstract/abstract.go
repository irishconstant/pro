package abstract

import (
	"auth"
	"core/contract"
	"core/ref"
	"core/tech"
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
	GetSource(id int) (*tech.Source, error)
	GetAllSources(regime int, currentPage int, pageSize int) (map[int]*tech.Source, error)
	GetSourceQuantityFiltered(u auth.User, name string) (int, error)

	// Работа с Юридическими лицами
	GetEntity(id int) (*contract.LegalEntity, error)

	// Работа с Физическими лицами
	GetPersonsFiltered(auth.User, int, int, int, string, string, string, string) (map[int]*contract.Person, error)
	GetPersonQuantityFiltered(auth.User, string, string, string, string) (int, error)
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
