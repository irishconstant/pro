package sqlserver

import "testing"

func TestCreateSelectQuery(t *testing.T) {
	var result string
	expectedResult := "SELECT * FROM [Administratum].dbo.[Users]"
	database := "Administratum"
	table := "Users"
	result = CreateSelectQuery(database, table)
	if expectedResult != result {
		t.Errorf("CreateSelectQuery with database %s and table %s returned result \"%s\", but expected \"%s\"", database, table, result, expectedResult)
	}
}
