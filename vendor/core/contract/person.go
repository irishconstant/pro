package contract

import (
	"auth"
	"core/ref"
	"time"
)

//Person Физическое лицо
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
