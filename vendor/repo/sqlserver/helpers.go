package sqlserver

import "fmt"

func selectWithPagination(databaseName string, sourceName string, orderParam string, whereParam string, whereValue string, pageSize int, currentPage int) string {
	switch sourceName {
	case "UserPersons":
		if currentPage <= 0 {
			return fmt.Sprintf("SELECT ID, C_Name, C_Patronymic_Name, C_Family_Name, F_Users FROM [%s].dbo.Persons WHERE %s = '%s' ORDER BY %s",
				databaseName, whereParam, whereValue, orderParam)
		}
		return fmt.Sprintf("SELECT ID, C_Name, C_Patronymic_Name, C_Family_Name, F_Users FROM [%s].dbo.Persons WHERE %s = '%s' ORDER BY %s"+
			" OFFSET %d ROWS  FETCH NEXT %d ROWS ONLY", databaseName, whereParam, whereValue, orderParam, pageSize*currentPage-pageSize, pageSize)
	case "Users":
		return fmt.Sprintf("SELECT [C_Name], [C_Family_Name] FROM [%s].dbo.Users AS u WHERE u.Login = '%s'",
			databaseName, whereValue)
	case "Person":
		return fmt.Sprintf("SELECT [ID], [C_Family_Name], [C_Name], [C_Patronymic_Name], [F_Users], ISNULL([F_Citizenship],0), CAST(B_Sex AS TINYINT), [D_Date_Birth], [D_Date_Death] FROM [%s].dbo.Persons WHERE %s = %s",
			databaseName, whereParam, whereValue)
	case "PersonContacts": //TODO: Сделать по аналогии этого варианта все остальные, достаточно возвращать идентификаторы
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
	case "PersonDocuments":
		return fmt.Sprintf("SELECT [ID], [C_Serial_Number], [C_Number], [C_From_Name], [C_From_Code], [D_Date_Begin], [D_Date_End] FROM [%s].dbo.Documents WHERE %s = %s ORDER BY %s",
			databaseName, whereParam, whereValue, orderParam)
	case "DocumentTypes":
		return fmt.Sprintf("SELECT [ID], [C_Serial_Number], [C_Number], [C_From_Name], [C_From_Code], [D_Date_Begin], [D_Date_End] FROM [%s].dbo.Documents WHERE %s = %s ORDER BY %s",
			databaseName, whereParam, whereValue, orderParam)
	case "Citizenships":
		return fmt.Sprintf("SELECT [ID], [C_Name] FROM [%s].[dbo].[Citizenships] WHERE %s = %s ORDER BY %s",
			databaseName, whereParam, whereValue, orderParam)
	case "Source":
		return fmt.Sprintf("SELECT [ID], [F_Object], [F_Season_Mode], [F_Fuel_Type], [N_Norm_Supply_Value], [F_Supplier_Electricity], [F_Voltage_Nominal], [F_Transport_Gas]"+
			", [F_Supplier_Gas], [F_Supplier_TechWater], [F_Supplier_HotWater], [F_Supplier_Canalisation], [F_Supplier_Heat] FROM [%s].dbo.Sources WHERE %s = %s",
			databaseName, whereParam, whereValue)
	case "Entity":
		return fmt.Sprintf("[ID], [F_User], [C_Name], [C_Short_Name], [INN], [KPP], [OGRN], [F_Entity_Type], [D_Date_Reg]"+
			" FROM [%s].dbo.Legal_Entities WHERE %s = %s",
			databaseName, whereParam, whereValue)
	case "Object":
		return fmt.Sprintf("[ID], [C_Name], [C_Address], [F_Region], [F_District], [F_City], [F_Town], [F_Street], [C_House], [C_Building], [U_FIAS]"+
			" FROM [%s].dbo.Objects WHERE %s = %s",
			databaseName, whereParam, whereValue)
	}
	return ""
}

func getBoolValue(i int) bool {
	if i == 1 {
		return true
	}
	return false
}
