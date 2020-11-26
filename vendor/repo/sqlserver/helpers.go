package sqlserver

import "fmt"

func selectWithPagination(databaseName string, tableName string, orderParam string, whereParam string, whereValue string, pageSize int, currentPage int) string {
	switch tableName {
	case "Customers":
		if currentPage <= 0 {
			return fmt.Sprintf("SELECT ID, C_Name, C_Patronymic_Name, C_Family_Name, F_Users FROM [%s].dbo.Customers WHERE %s = '%s' ORDER BY %s",
				databaseName, whereParam, whereValue, orderParam)
		}
		return fmt.Sprintf("SELECT ID, C_Name, C_Patronymic_Name, C_Family_Name, F_Users FROM [%s].dbo.Customers WHERE %s = '%s' ORDER BY %s"+
			" OFFSET %d ROWS  FETCH NEXT %d ROWS ONLY", databaseName, whereParam, whereValue, orderParam, pageSize*currentPage-pageSize, pageSize)
	case "Users":
		return fmt.Sprintf("SELECT u.C_Name, u.C_Family_Name FROM [%s].dbo.Users AS u WHERE u.Login = '%s'",
			databaseName, whereValue)
	}
	return ""
}



