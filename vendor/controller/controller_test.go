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
	expectedResult := "/Person?name=имя&familyname=фамилия&sex=true&"
	expectedResult2 := "/Person?sex=true&name=имя&familyname=фамилия&"
	expectedResult3 := "/Person?familyname=фамилия&name=имя&sex=true&"
	expectedResult4 := "/Person?sex=true&familyname=фамилия&name=имя&"
	expectedResult5 := "/Person?familyname=фамилия&sex=true&name=имя&"
	expectedResult6 := "/Person?name=имя&sex=true&familyname=фамилия&"

	// Действие
	url := makeURLWithAttributes("Person", testParams)
	// Утверждение
	if url != expectedResult && url != expectedResult2 && url != expectedResult3 && url != expectedResult4 && url != expectedResult5 && url != expectedResult6 {
		t.Errorf("Некорректно работает создание строки из параметров")
	}
}
