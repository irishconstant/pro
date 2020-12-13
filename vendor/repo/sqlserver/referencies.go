package sqlserver

import (
	"core/contract"
	"fmt"
)

// GetContactType возвращает Тип контакта по его ИД
func (s SQLServer) GetContactType(id int) (*contract.ContactType, error) {
	rows, err := s.db.Query(fmt.Sprintf("SELECT ID, C_Name, C_Validation_Mask, B_Address  FROM %s.dbo.Contact_Types WHERE ID = %d", s.dbname, id))
	if err != nil {
		fmt.Printf("Ошибка с получением Типа контакта")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			ID             int
			name           string
			validationMask string
			isAddress      int
		)
		rows.Scan(
			&ID,
			&name,
			&validationMask,
			&isAddress)

		contactType := contract.ContactType{
			Key:        ID,
			Name:       name,
			Validation: validationMask,
			IsAddress:  getBoolValue(isAddress)}
		return &contactType, nil
	}
	return nil, err
}

// GetAllContactTypes возвращает все возможные типы контактов
func (s SQLServer) GetAllContactTypes() ([]*contract.ContactType, error) {
	rows, err := s.db.Query(fmt.Sprintf("SELECT ID FROM %s.dbo.Contact_Types", s.dbname))
	if err != nil {
		fmt.Printf("Ошибка с получением Типов контакта")
		return nil, err
	}
	defer rows.Close()
	var contactTypes []*contract.ContactType
	for rows.Next() {
		var (
			ID int
		)
		rows.Scan(
			&ID,
		)
		newContactType, _ := s.GetContactType(ID)
		contactTypes = append(contactTypes, newContactType)
	}
	return contactTypes, err
}

// GetCitizenship возвращает гражданство по его ид
func (s SQLServer) GetCitizenship(id int) (*contract.Citizenship, error) {
	if id == 0 {
		id = 1
	}

	rows, err := s.db.Query(fmt.Sprintf("SELECT ID, C_Name FROM %s.dbo.Citizenships WHERE ID = %d", s.dbname, id))
	if err != nil {
		fmt.Printf("Ошибка с получением Гражданства")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			ID   int
			name string
		)
		rows.Scan(
			&ID,
			&name)

		citizenship := contract.Citizenship{
			Key:  ID,
			Name: name}
		return &citizenship, nil
	}
	return nil, err
}

// GetAllCitizenship возвращает все возможные варианты гражданства
func (s SQLServer) GetAllCitizenship() ([]*contract.Citizenship, error) {
	rows, err := s.db.Query(fmt.Sprintf("SELECT ID FROM %s.dbo.Citizenships", s.dbname))
	if err != nil {
		fmt.Printf("Ошибка с получением Гражданств")
		return nil, err
	}
	defer rows.Close()

	var citizenships []*contract.Citizenship
	for rows.Next() {
		var (
			ID int
		)
		rows.Scan(
			&ID)

		newCitizenship, _ := s.GetCitizenship(ID)
		citizenships = append(citizenships, newCitizenship)
	}
	return citizenships, err
}
