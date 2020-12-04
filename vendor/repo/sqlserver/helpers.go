package sqlserver

import "fmt"

func selectWithPagination(databaseName string, sourceName string, orderParam string, whereParam string, whereValue string, pageSize int, currentPage int) string {
	switch sourceName {
	case "UserCustomers":
		if currentPage <= 0 {
			return fmt.Sprintf("SELECT ID, C_Name, C_Patronymic_Name, C_Family_Name, F_Users FROM [%s].dbo.Customers WHERE %s = '%s' ORDER BY %s",
				databaseName, whereParam, whereValue, orderParam)
		}
		return fmt.Sprintf("SELECT ID, C_Name, C_Patronymic_Name, C_Family_Name, F_Users FROM [%s].dbo.Customers WHERE %s = '%s' ORDER BY %s"+
			" OFFSET %d ROWS  FETCH NEXT %d ROWS ONLY", databaseName, whereParam, whereValue, orderParam, pageSize*currentPage-pageSize, pageSize)
	case "Users":
		return fmt.Sprintf("SELECT [C_Name], [C_Family_Name] FROM [%s].dbo.Users AS u WHERE u.Login = '%s'",
			databaseName, whereValue)
	case "Customer":
		return fmt.Sprintf("SELECT [ID], [C_Family_Name], [C_Name], [C_Patronymic_Name], [F_Users], [F_Citizenship], [B_Sex], [D_Date_Birth], [D_Date_Death] FROM [%s].dbo.Customers WHERE %s = %s",
			databaseName, whereParam, whereValue)
	case "CustomerContacts": //TODO: Сделать по аналогии этого варианта все остальные, достаточно возвращать идентификаторы
		return fmt.Sprintf("SELECT [ID] FROM %s.dbo.Contacts WHERE  %s = %s ORDER BY %s",
			databaseName, whereParam, whereValue, orderParam)
	case "ContactType":
		return fmt.Sprintf("SELECT [ID], [C_Name] FROM [%s].dbo.Contact_Types WHERE %s = %s",
			databaseName, whereParam, whereValue)
	case "Contact":
		return fmt.Sprintf("SELECT [ID], [F_Contact_Type], [C_Value], [B_Primary] FROM [%s].dbo.Contacts WHERE %s = %s",
			databaseName, whereParam, whereValue)
	case "ContactTypes":
		return fmt.Sprintf("SELECT [ID], [C_Name] FROM [%s].dbo.Contact_Types",
			databaseName)
	case "CustomerDocuments":
		return fmt.Sprintf("SELECT [ID], [C_Serial_Number], [C_Number], [C_From_Name], [C_From_Code], [D_Date_Begin], [D_Date_End] FROM [%s].dbo.Documents WHERE %s = %s ORDER BY %s",
			databaseName, whereParam, whereValue, orderParam)
	case "DocumentTypes":
		return fmt.Sprintf("SELECT [ID], [C_Serial_Number], [C_Number], [C_From_Name], [C_From_Code], [D_Date_Begin], [D_Date_End] FROM [%s].dbo.Documents WHERE %s = %s ORDER BY %s",
			databaseName, whereParam, whereValue, orderParam)
	case "Citizenships":
		return fmt.Sprintf("SELECT [ID], [C_Name] FROM [%s].[dbo].[Citizenships] WHERE %s = %s ORDER BY %s",
			databaseName, whereParam, whereValue, orderParam)
	}
	return ""
}
