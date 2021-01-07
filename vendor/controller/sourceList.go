package controller

import (
	"core/ref"
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

	// Работаем с текущим пользователем
	session, err := Store.Get(r, "cookie-name")
	check(err)
	user := GetUser(session)
	check(err)

	if r.Method == http.MethodPost {
		// Получаем данные фильтров из формы и формируем параметры для вызова
		params := make(map[string]string)
		params["name"] = r.FormValue("name")
		params["address"] = r.FormValue("address")
		params["seasonmode"] = r.FormValue("seasonmode")
		params["fueltype"] = r.FormValue("fueltype")
		params["period"] = r.FormValue("period")
		filteredAddress := makeURLWithAttributes("source", params)
		// Переходим на этот урл
		http.Redirect(w, r, filteredAddress, http.StatusFound)
	}

	// Получаем текущую страницу из параметров
	key := r.URL.Query().Get("page")
	var page int
	if key != "" {
		page, _ = strconv.Atoi(key)
	} else {
		page = 1
	}

	// Получаем текущий период из параметров
	var calcPeriod *ref.CalcPeriod
	period := r.URL.Query().Get("period")
	if period == "" {
		calcPeriod, err = h.connection.GetCurrentPeriod()
	} else {
		calcPeriodID, err := strconv.Atoi(period)
		calcPeriod, err = h.connection.GetCalcPeriod(calcPeriodID)
		if err != nil {
			fmt.Println("Передано ошибочное значение расчётного периода")
		}
	}

	// Получаем параметры фильтрации из урла
	name := r.URL.Query().Get("name")
	address := r.URL.Query().Get("address")
	seasonMode := r.URL.Query().Get("seasonmode")
	fuelType := r.URL.Query().Get("fueltype")

	seasonModeI, err := strconv.Atoi(seasonMode)
	fuelTypeI, err := strconv.Atoi(fuelType)

	// Справочники
	fuelTypes, err := h.connection.GetAllFuelTypes()
	seasonModes, err := h.connection.GetAllSeasonModes()
	calcPeriods, err := h.connection.GetAllCalcPeriods()
	for _, value := range calcPeriods {
		if calcPeriod.Key == value.Key {
			value.IsSelected = true
		}
	}

	refBox := map[interface{}]interface{}{
		"FuelTypes":     fuelTypes,
		"SeasonModes":   seasonModes,
		"CalcPeriods":   calcPeriods,
		"CurrentPeriod": calcPeriod.Key,
	}
	check(err)

	/*-------------------------------------------
	 Работаем с теплоисточниками
	--------------------------------------------*/
	// Получаем количество теплоисточников
	quantity, err := h.connection.GetSourceQuantityFiltered(*user, name, address, seasonModeI, fuelTypeI, calcPeriod)
	check(err)
	sourceBook := SourceBook{Count: quantity}

	// Если необходима пагинация
	if sourceBook.Count > h.pageSize {
		sourcePerPage, err := h.connection.GetAllSources(1, page, h.pageSize, name, address, seasonModeI, fuelTypeI, calcPeriod)
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
		if address != "" {
			address = "&address=" + address
		}
		if fuelType != "" {
			fuelType = "&fueltype=" + fuelType
		}
		if seasonMode != "" {
			seasonMode = "&seasonmode=" + seasonMode
		}
		if period != "" {
			period = "&period=" + period
		}

		sourceBook.Pages = MakePages(1, int(math.Ceil(float64(sourceBook.Count)/float64(h.pageSize))), page)
		for key := range sourceBook.Pages {
			sourceBook.Pages[key].URL = fmt.Sprintf("/source?%s%s%s%s%s", name, address, fuelType, seasonMode, period)
		}
		currentInformation := sessionInformation{User: *user, Attribute: sourceBook, AttributeMap: refBox}
		executeHTML("source", "list", w, currentInformation)

	} else {
		sourcePerPage, err := h.connection.GetAllSources(0, page, h.pageSize, name, address, seasonModeI, fuelTypeI, calcPeriod)
		check(err)

		for _, value := range sourcePerPage {
			sourceBook.Sources = append(sourceBook.Sources, *value)
		}

		currentInformation := sessionInformation{User: *user, Attribute: sourceBook, AttributeMap: refBox}
		executeHTML("source", "list", w, currentInformation)
	}
}
