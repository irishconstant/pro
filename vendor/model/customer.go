package model

import "math"

//Customer представляет из себя типичного Потребителя
type Customer struct {
	Key            int
	Name           string
	PatronymicName string
	FamilyName     string
	User           User // Ответственный пользователь
}

//CustomersBook представляет из себя набор Потребителей определённого Пользователя
type CustomersBook struct {
	CustomerCount int
	CurrentPage   int
	Pages         []Page // Приходится хранить слайс данных для генерации из шаблона
	Customers     []Customer
}

//Page представляет для любых представлений
type Page struct {
	Number       int
	PreviousPage bool
	NextPage     bool
	CurrentPage  bool
}

//MakePages генерирует последовательности страниц
func MakePages(min, max, current int) []Page {
	pages := make([]Page, max-min+1)
	for i := range pages {
		if i == 0 || math.Abs(float64(current-min-i)) <= 2 || i == max-1 {
			pages[i].Number = min + i
		}

		if current == min+i {
			pages[i].CurrentPage = true
		}

		if current-min-i == -2 && current < max-min+1 {
			//pages[i].Number = current - 1
			pages[i].NextPage = true
		}

		if current-min-i == 2 && current > 1 {
			//pages[i].Number = current + 1
			pages[i].PreviousPage = true
		}

	}
	return pages
}
