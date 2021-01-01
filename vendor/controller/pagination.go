package controller

import (
	"auth"
	"core/contract"
	"core/tech"
	"math"
)

// UserBook хранит информацию наборе Пользователей
type UserBook struct {
	UserCount   int
	CurrentPage int
	Pages       []Page
	Users       []auth.User
}

// RoleBook набор Ролей для интерфейса
type RoleBook struct {
	RoleCount int
	Roles     []auth.Role
}

// ContactBook представляет из себя набор Контактов определённого Потребителя или Пользователя
type ContactBook struct {
	ContactCount int
	Contacts     []contract.Contact
}

// DocBook представляет из себя набор Документов определённого Потребителя или Пользователя
type DocBook struct {
	DocCount int
	Docs     []contract.Doc
}

//PersonBook представляет из себя набор Потребителей определённого Пользователя
type PersonBook struct {
	PersonCount int // Сколько Потребителей всего в книге
	CurrentPage int // Текущая страница

	Pages   []Page
	Persons []contract.Person
}

//SourceBook представляет из себя набор тепло-электро-водоисточников
type SourceBook struct {
	Count       int
	CurrentPage int

	Pages   []Page
	Sources []tech.Source
}

//Page представляет страницу любых представлений
type Page struct {
	Number       int
	FirstPage    bool
	LastPage     bool
	PreviousPage bool
	NextPage     bool
	CurrentPage  bool
	URL          string
}

//MakePages генерирует последовательности страниц для отображения
func MakePages(first, last, current int) []Page {
	var pages []Page

	for i := first; i <= last; i++ {

		if i == first || math.Abs(float64(current-i)) <= 2 || i == last {
			page := Page{Number: i}

			switch i {
			case current:
				page.CurrentPage = true
			case first:
				page.FirstPage = true
			case last:
				page.LastPage = true
			}

			if current-i == -1 && current < last {
				page.NextPage = true
			}

			if current-i == 1 && current > first {
				page.PreviousPage = true
			}
			pages = append(pages, page)
		}
	}
	return pages
}
