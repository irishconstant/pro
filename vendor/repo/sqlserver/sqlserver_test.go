package sqlserver

import "testing"

// Пока бесполезный тест, просто чтобы не забыть, что теоретически у нас TDD
func TestCreateSelectQueryWithPagination(t *testing.T) {
	var result string
	expectedResult := "SELECT ID, C_Name, C_Patronymic_Name, C_Family_Name, F_Users FROM [Administratum].dbo.Customers WHERE F_Users = 'rode_orm' ORDER BY ID OFFSET 3 ROWS  FETCH NEXT 3 ROWS ONLY"
	result = selectWithPagination("Administratum", "Customers", "ID", "F_Users", "rode_orm", 3, 2)
	if expectedResult != result {
		t.Errorf("Должно быть: %s, а получилось: %s", expectedResult, result)
	}
}

func TestCreateSelectQueryAll(t *testing.T) {
	var result string
	expectedResult := "SELECT ID, C_Name, C_Patronymic_Name, C_Family_Name, F_Users FROM [Administratum].dbo.Customers WHERE F_Users = 'rode_orm' ORDER BY ID"
	result = selectWithPagination("Administratum", "Customers", "ID", "F_Users", "rode_orm", 0, 0)
	if expectedResult != result {
		t.Errorf("Должно быть: %s, а получилось: %s", expectedResult, result)
	}
}
