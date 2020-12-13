package controller

import (
	"reflect"
	"testing"
)

func TestMakeURLs(t *testing.T) {
	// Организация
	var testParams = map[string]string{
		"name":       "имя",
		"familyname": "фамилия",
		"sex":        "true",
	}
	expectedResult := "/person?name=имя&familyname=фамилия&sex=true&"
	expectedResult2 := "/person?sex=true&name=имя&familyname=фамилия&"
	expectedResult3 := "/person?familyname=фамилия&name=имя&sex=true&"
	expectedResult4 := "/person?sex=true&familyname=фамилия&name=имя&"
	expectedResult5 := "/person?familyname=фамилия&sex=true&name=имя&"
	expectedResult6 := "/person?name=имя&sex=true&familyname=фамилия&"

	// Действие
	url := makeURLWithAttributes("person", testParams)
	// Утверждение
	if url != expectedResult && url != expectedResult2 && url != expectedResult3 && url != expectedResult4 && url != expectedResult5 && url != expectedResult6 {
		t.Errorf("Некорректно работает создание строки из параметров")
	}
}

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
