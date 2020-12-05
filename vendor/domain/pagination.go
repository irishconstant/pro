package domain

import "math"

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
