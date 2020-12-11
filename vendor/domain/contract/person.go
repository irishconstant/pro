package contract

import (
	"domain/auth"
	"domain/ref"
	"domain/sys"
	"time"
)

//Person представляет из себя типичного Потребителя
type Person struct {
	Key int

	User           auth.User // Ответственный пользователь
	Name           string
	PatronymicName string
	FamilyName     string
	Citizenship    Citizenship
	Sex            bool
	DateBirth      time.Time
	DateDeath      time.Time

	PossibleUsers []auth.User // Доступные пользователи для привязки
	Contacts      []Contact   // Контакты
	Docs          []Doc       // Документы
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
	Address ref.Address
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

//PersonsBook представляет из себя набор Потребителей определённого Пользователя
type PersonsBook struct {
	PersonCount int // Сколько Потребителей всего в книге
	CurrentPage int // Текущая страница

	Pages   []sys.Page
	Persons []Person
}
