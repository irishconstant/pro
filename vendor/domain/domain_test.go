package domain

import (
	"reflect"
	"testing"
)

func TestMakePages(t *testing.T) {

	// Организация
	var expectedResult []Page
	firstPage := Page{Number: 1, FirstPage: true}
	lastPage := Page{Number: 100, LastPage: true}
	nextPage := Page{Number: 6, NextPage: true}
	previousPage := Page{Number: 4, PreviousPage: true}
	currentPage := Page{Number: 5, CurrentPage: true}
	expectedResult = append(expectedResult, firstPage)
	expectedResult = append(expectedResult, lastPage)
	expectedResult = append(expectedResult, nextPage)
	expectedResult = append(expectedResult, previousPage)
	expectedResult = append(expectedResult, currentPage)

	// Действие
	result := MakePages(1, 100, 5)

	// Утверждение
	if reflect.DeepEqual(expectedResult, result) {
		t.Errorf("Некорректно создаются страницы для пагинации")
	}
}
