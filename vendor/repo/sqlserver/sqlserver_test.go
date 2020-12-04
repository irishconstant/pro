package sqlserver

import "testing"

func TestCreateSelectQueryWithPagination(t *testing.T) {
	var result string
	expectedResult := "SELECT ID, C_Name, C_Patronymic_Name, C_Family_Name, F_Users FROM [Administratum].dbo.Customers WHERE F_Users = 'rode_orm' ORDER BY ID OFFSET 3 ROWS  FETCH NEXT 3 ROWS ONLY"
	result = selectWithPagination("Administratum", "UserCustomers", "ID", "F_Users", "rode_orm", 3, 2)
	if expectedResult != result {
		t.Errorf("Должно быть: %s, а получилось: %s", expectedResult, result)
	}
}

func TestCreateSelectCustomerContacts(t *testing.T) {
	var result string
	expectedResult := "SELECT ID, F_Contact_Type, C_Value, F_Customer, B_Primary FROM Administratum.dbo.Contacts" +
		" WHERE  F_Customer = 1 ORDER BY ID"
	result = selectWithPagination("Administratum", "CustomerContacts", "ID", "F_Customer", "1", 1, 1)
	if expectedResult != result {
		t.Errorf("Должно быть: %s, а получилось: %s", expectedResult, result)
	}
}

func TestCreateSelectContactType(t *testing.T) {
	var result string
	expectedResult := "SELECT ID, C_Name FROM [Administratum].dbo.Contact_Types WHERE ID = 1"
	//(databaseName string, sourceName string, orderParam string, whereParam string, whereValue string, pageSize int, currentPage int)
	result = selectWithPagination("Administratum", "ContactType", "ID", "ID", "1", 1, 1)
	if expectedResult != result {
		t.Errorf("Должно быть: %s, а получилось: %s", expectedResult, result)
	}
}

func TestCreateSelectQueryAll(t *testing.T) {
	var result string
	expectedResult := "SELECT ID, C_Name, C_Patronymic_Name, C_Family_Name, F_Users FROM [Administratum].dbo.Customers WHERE F_Users = 'rode_orm' ORDER BY ID"
	result = selectWithPagination("Administratum", "UserCustomers", "ID", "F_Users", "rode_orm", 0, 0)
	if expectedResult != result {
		t.Errorf("Должно быть: %s, а получилось: %s", expectedResult, result)
	}
}
