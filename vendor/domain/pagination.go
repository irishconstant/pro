package domain

import "math"

//Page представляет для любых представлений
type Page struct {
	Number       int
	PreviousPage bool
	NextPage     bool
	CurrentPage  bool
}

//MakePages генерирует последовательности страниц
func MakePages(first, last, current int) []Page {
	var pages []Page

	for i := first; i <= last; i++ {

		if i == first || math.Abs(float64(current-i)) <= 2 || i == last {
			page := Page{Number: i}

			if current == i {
				page.CurrentPage = true
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
