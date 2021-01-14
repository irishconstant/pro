package contract

import (
	"core/auth"
	"core/ref"
	"time"
)

// LegalEntity юридическое лицо
type LegalEntity struct {
	Key int

	User      auth.User // Ответственный пользователь
	Name      string
	ShortName string

	INN  string //	ИНН
	KPP  string // КПП
	OGRN string // ОГРН

	Form ref.OrganistaionForm

	roles []EntityRole

	DateReg time.Time

	PossibleUsers []auth.User // Доступные пользователи для привязки
	Contacts      []Contact   // Контакты
	Docs          []Doc       // Документы
}

// EntityRole Роль юридического лица
type EntityRole struct {
	Key      int
	RoleName string
}
