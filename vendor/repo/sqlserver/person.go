package sqlserver

import (
	"auth"
	"core/contract"
	"fmt"
	"strconv"
	"time"
)

//GetPersonQuantityFiltered возвращает КОЛИЧЕСТВО Потребителей конкретного Пользователя с учётом переданных фильтров
func (s *SQLServer) GetPersonQuantityFiltered(u auth.User, name string, familyname string, patrname string, sex string) (int, error) {
	var query string
	if sex == "" {
		query =
			fmt.Sprintf("EXEC %s.dbo.GetQuantityFilteredPersons  '%s', '%s', '%s', '%s', NULL",
				s.dbname, u.Key, familyname, name, patrname)
	} else {
		query =
			fmt.Sprintf("EXEC %s.dbo.GetQuantityFilteredPersons  '%s', '%s', '%s', '%s', %s",
				s.dbname, u.Key, familyname, name, patrname, sex)
	}

	rows, err := s.db.Query(query)

	if err != nil {
		fmt.Println("Ошибка c запросом: ", query, err)
		return 0, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			Num int
		)
		rows.Scan(
			&Num)
		return Num, nil

	}

	return 0, err
}

//GetPersonsFiltered возвращает всех Потребителей конкретного Пользователя для страницы с учётом переданных фильтров
func (s *SQLServer) GetPersonsFiltered(u auth.User, regime int, currentPage int, pageSize int, name string, familyname string, patrname string, sex string) (map[int]*contract.Person, error) {
	Persons := make(map[int]*contract.Person)

	var query string
	if sex == "" {
		query =
			fmt.Sprintf("EXEC %s.dbo.GetFilteredPaginatedPersons '%s', '%s', '%s', '%s', NULL, %d, %d, %d",
				s.dbname, u.Key, familyname, name, patrname, pageSize*currentPage-pageSize, pageSize, regime)
	} else {
		query =
			fmt.Sprintf("EXEC %s.dbo.GetFilteredPaginatedPersons '%s', '%s', '%s', '%s', %s, %d, %d, %d",
				s.dbname, u.Key, familyname, name, patrname, sex, pageSize*currentPage-pageSize, pageSize, regime)

	}

	rows, err := s.db.Query(query)

	if err != nil {
		fmt.Println("Ошибка c запросом: ", query, err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			ID             int
			FamilyName     string
			Name           string
			PatronymicName string
			UserKey        string
			CitizenshipKey int
			Sex            bool
			DateBirth      string
			DateDeath      string
		)
		rows.Scan(
			&ID,
			&FamilyName,
			&Name,
			&PatronymicName,
			&UserKey,
			&CitizenshipKey,
			&Sex,
			&DateBirth,
			&DateDeath)

		newPerson, _ := s.GetPerson(ID)

		if ID != 0 {
			Persons[ID] = newPerson
		}
	}
	return Persons, nil
}

//CreatePerson создаёт нового Потребителя
func (s SQLServer) CreatePerson(c *contract.Person) error {
	var bSex int
	if strconv.FormatBool(c.Sex) == "true" {
		bSex = 1
	} else {
		bSex = 0
	}
	// TODO: Добавить дату рождения и прочую шнягу
	rows, err := s.db.Query(fmt.Sprintf("INSERT INTO Persons (C_Name, C_Family_Name, C_Patronymic_Name, B_Sex, F_Users) SELECT '%s', '%s', '%s', %d, '%s' SELECT SCOPE_IDENTITY()", c.Name, c.FamilyName, c.PatronymicName, bSex, c.User.Key))
	defer rows.Close()
	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return err
	}

	for rows.Next() {
		rows.Scan(&c.Key)
	}
	fmt.Println("Создан Потребитель с идентификатором:", c.Key)
	return nil
}

//GetPerson возвращает пользователя по первичному ключу
func (s SQLServer) GetPerson(id int) (*contract.Person, error) {

	rows, err := s.db.Query(creatorSelect(s.dbname, "Person", "ID", "ID", strconv.Itoa(id)))

	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return nil, err
	}
	var Person contract.Person
	defer rows.Close()
	for rows.Next() {
		var (
			ID             int
			FamilyName     string
			Name           string
			PatronymicName string
			UserLogin      string
			CitizenshipKey int
			Sex            int
			DateBirth      string
			DateDeath      string
		)
		rows.Scan(&ID, &FamilyName, &Name, &PatronymicName, &UserLogin, &CitizenshipKey, &Sex, &DateBirth, &DateDeath)
		user, err := s.GetUser(UserLogin)
		if err != nil {
			fmt.Println("Ошибка c получением Пользователя: ", err)
			return nil, err
		}

		citizenship, _ := s.GetCitizenship(CitizenshipKey)
		DateBirthG, _ := time.Parse(time.RFC3339, DateBirth)
		DateDeathG, _ := time.Parse(time.RFC3339, DateDeath)

		Person = contract.Person{
			Key:            ID,
			Name:           Name,
			PatronymicName: PatronymicName,
			FamilyName:     FamilyName,
			Sex:            getBoolValue(Sex),
			DateBirth:      DateBirthG,
			DateDeath:      DateDeathG,
			Citizenship:    *citizenship,
			User:           *user}

	}

	err = s.GetPersonContacts(&Person)

	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return nil, err
	}

	return &Person, nil
}

// GetPersonContacts получает Контакты для Пользователя
func (s SQLServer) GetPersonContacts(Person *contract.Person) error {
	var contacts []contract.Contact
	rows, err := s.db.Query(creatorSelect(s.dbname, "ContactList", "ID", "F_Person", strconv.Itoa(Person.Key)))
	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		rows.Scan(&id)
		contact, err := s.GetContact(id)
		contacts = append(contacts, *contact)
		if err != nil {
			fmt.Println("Ошибка c запросом: ", err)
			return err
		}
	}
	Person.Contacts = contacts
	return err
}

//GetContact возвращает Контакт по его идентификатору
func (s SQLServer) GetContact(id int) (*contract.Contact, error) {
	//[ID], [F_Contact_Type], [C_Value], [B_Primary]
	rows, err := s.db.Query(creatorSelect(s.dbname, "Contact", "ID", "ID", strconv.Itoa(id)))
	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return nil, err
	}
	defer rows.Close()
	var contact contract.Contact
	for rows.Next() {
		var (
			ID            int
			contactTypeID int
			value         string
			primary       bool
		)
		rows.Scan(&ID, &contactTypeID, &value, &primary)

		contactType, err := s.GetContactType(contactTypeID)
		if err != nil {
			fmt.Println("Ошибка c получением Контакта: ", err)
			return nil, err
		}

		contact = contract.Contact{
			Key:   ID,
			Value: value,
			Type:  *contactType,
		}

	}
	return &contact, nil
}

//UpdatePerson обновляет данные Потребителя
func (s SQLServer) UpdatePerson(Person *contract.Person) error {
	var sex string

	if Person.Sex == true {
		sex = "1"
	} else {
		sex = "0"
	}

	_, err := s.db.Query(fmt.Sprintf("UPDATE %s.dbo.Persons"+
		" SET C_Family_Name = '%s', C_Name = '%s', C_Patronymic_Name = '%s', F_Users = '%s', D_Date_Birth = TRY_CAST('%s' AS DATETIME), D_Date_Death = TRY_CAST('%s' AS DATETIME), B_Sex = %s"+
		" WHERE ID =  %s",
		s.dbname, Person.FamilyName, Person.Name, Person.PatronymicName, Person.User.Key, ConvertDate(Person.DateBirth), ConvertDate(Person.DateDeath), sex, strconv.Itoa(Person.Key)))
	if err != nil {
		fmt.Println("Ошибка при обновлени Пользователя", err)
	}
	return err
}

//DeletePerson удаляет Потребителя
func (s SQLServer) DeletePerson(Person *contract.Person) error {
	_, err := s.db.Query(fmt.Sprintf("DELETE FROM %s.dbo.Persons WHERE ID =  %s",
		s.dbname, strconv.Itoa(Person.Key)))
	if err != nil {
		fmt.Println("Ошибка при удалении Пользователя", err)
	}
	return err
}
