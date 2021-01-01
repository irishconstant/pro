package controller

import (
	"auth"
	"fmt"
	"math"
	"net/http"
	"strconv"
)

/* Фактические параметры работы котельных заполняются для каждого расчётного периода.
При этом для каждого расчётного периода определяется: интерфейс из двух частей.
Первая часть «Шапка», вторая – детальные данные по конкретным теплоисточникам.
*/
func (h *DecoratedHandler) source(w http.ResponseWriter, r *http.Request) { //

	// Получаем текущую страницу из параметров
	key := r.URL.Query().Get("page")
	var page int
	if key != "" {
		page, _ = strconv.Atoi(key)
	} else {
		page = 1
	}

	// Получаем параметры фильтрации из урла
	name := r.URL.Query().Get("name")

	// Работаем с текущим пользователем
	session, err := auth.Store.Get(r, "cookie-name")
	check(err)
	user := auth.GetUser(session)
	check(err)
	//err = h.connection.GetUserAttributes(&user)
	//check(err)

	// Работаем с источниками
	quantity, err := h.connection.GetSourceQuantityFiltered(*user, "")
	check(err)
	sourceBook := SourceBook{Count: quantity}

	// Если необходима пагинация
	if sourceBook.Count > h.pageSize {
		sourcePerPage, err := h.connection.GetAllSources(1, page, h.pageSize)
		check(err)
		for _, value := range sourcePerPage {
			sourceBook.Sources = append(sourceBook.Sources, *value)
		}
		sourceBook.CurrentPage = page

		// Создаем страницы для показа (1, одна слева от текущей, одна справа от текущей, последняя)
		// Инициализируем фильтры для кнопок пагинации, которые к нам ранее пришли в POST запросе
		if name != "" {
			name = "&name=" + name
		}
		sourceBook.Pages = MakePages(1, int(math.Ceil(float64(sourceBook.Count)/float64(h.pageSize))), page)
		for key := range sourceBook.Pages {
			sourceBook.Pages[key].URL = fmt.Sprintf("/source?%s%s%s%s", name, "", "", "")
		}
		currentInformation := sessionInformation{User: *user, Attribute: sourceBook}
		executeHTML("source", "list", w, currentInformation)

	} else {
		sourcePerPage, err := h.connection.GetAllSources(0, 1, h.pageSize)
		check(err)

		for _, value := range sourcePerPage {
			sourceBook.Sources = append(sourceBook.Sources, *value)
		}

		currentInformation := sessionInformation{User: *user, Attribute: sourceBook}

		executeHTML("source", "list", w, currentInformation)
	}
}
