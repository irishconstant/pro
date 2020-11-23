package model

import "testing"

// Пока бесполезный тест, просто чтобы не забыть, что теоретически у нас TDD
func TestModel(t *testing.T) {
	var result string
	expectedResult := "OK"
	result = "OK"
	if expectedResult != result {
		t.Errorf("Здесь текст ошибки %s, %s", expectedResult, result)
	}
}
