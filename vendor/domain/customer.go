package domain

//Customer представляет из себя типичного Потребителя
type Customer struct {
	Key int

	User           User // Ответственный пользователь
	Name           string
	PatronymicName string
	FamilyName     string
	Citizenship    Citizenship
	Sex            bool // true - мужчина false - женщина
	DateBirth      string
	DateDeath      string

	PossibleUsers []User    // Доступные пользователи для привязки
	Contacts      []Contact // Контакты
	Docs          []Doc     // Документы
}

// Doc Документ
type Doc struct {
	Key          int
	SerialNumber string // Серия
	Number       string // Номер
	FromName     string // Кем выдан
	FromCode     string // Код подразделения
	DateBegin    string // Дата выдачи
	DateEnd      string // Дата окончания срока действия
}

// Contact Контакт
type Contact struct {
	Key     int
	Type    ContactType
	Value   string
	Address Address
}

// ContactBook представляет из себя набор Контактов определённого Потребителя или Пользователя
type ContactBook struct {
	ContactCount int
	Contacts     []Contact
}

// DocBook представляет из себя набор Документов определённого Потребителя или Пользователя
type DocBook struct {
	DocCount int
	Docs     []Doc
}

//CustomersBook представляет из себя набор Потребителей определённого Пользователя
type CustomersBook struct {
	CustomerCount int    // Сколько Потребителей всего в книге
	CurrentPage   int    // Текущая страница
	Pages         []Page // Какие страницы отображаются для текущей страницы
	Customers     []Customer
}
