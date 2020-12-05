package controller

import (
	"testing"
)

func TestMakeURLs(t *testing.T) {
	// Организация
	var testParams = map[string]string{
		"name":       "имя",
		"familyname": "фамилия",
		"sex":        "true",
	}
	expectedResult := "/customer?name=имя&familyname=фамилия&sex=true&"
	expectedResult2 := "/customer?sex=true&name=имя&familyname=фамилия&"
	expectedResult3 := "/customer?familyname=фамилия&name=имя&sex=true&"
	expectedResult4 := "/customer?sex=true&familyname=фамилия&name=имя&"
	expectedResult5 := "/customer?familyname=фамилия&sex=true&name=имя&"
	expectedResult6 := "/customer?name=имя&sex=true&familyname=фамилия&"

	// Действие
	url := makeURLWithAttributes("customer", testParams)
	// Утверждение
	if url != expectedResult && url != expectedResult2 && url != expectedResult3 && url != expectedResult4 && url != expectedResult5 && url != expectedResult6 {
		t.Errorf("Некорректно работает создание строки из параметров")
	}
}
