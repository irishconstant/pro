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
