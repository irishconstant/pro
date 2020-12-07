package sqlserver

import (
	"domain"
	"fmt"
	"strconv"
	"time"
)

//GetUserFiltredPersonsPagination возвращает всех Потребителей конкретного Пользователя для страницы с учётом переданных фильтров
func (s *SQLServer) GetUserFiltredPersonsPagination(u domain.User, regime int, currentPage int, pageSize int, name string, familyname string, patrname string, sex string) (map[int]*domain.Person, error) {
	Persons := make(map[int]*domain.Person)

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
func (s SQLServer) CreatePerson(c *domain.Person) error {
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
	fmt.Println("Создан Пользователь с идентификатором:", c.Key)
	return nil
}

//GetPerson возвращает пользователя по первичному ключу
func (s SQLServer) GetPerson(id int) (*domain.Person, error) {

	rows, err := s.db.Query(selectWithPagination(s.dbname, "Person", "ID", "ID", strconv.Itoa(id), 0, 0))

	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return nil, err
	}
	var Person domain.Person
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
			sexBool        bool
			DateBirth      string
			DateDeath      string
		)
		rows.Scan(&ID, &FamilyName, &Name, &PatronymicName, &UserLogin, &CitizenshipKey, &Sex, &DateBirth, &DateDeath)
		user, err := s.GetUser(UserLogin)
		if err != nil {
			fmt.Println("Ошибка c получением Пользователя: ", err)
			return nil, err
		}

		// Неудобно, конечно, но не писать же целый конструктор
		if Sex == 1 {
			sexBool = true
		} else {
			sexBool = false
		}
		citizenship, _ := s.GetCitizenship(CitizenshipKey)
		DateBirthG, _ := time.Parse(time.RFC3339, DateBirth)
		DateDeathG, _ := time.Parse(time.RFC3339, DateDeath)

		Person = domain.Person{
			Key:            ID,
			Name:           Name,
			PatronymicName: PatronymicName,
			FamilyName:     FamilyName,
			Sex:            sexBool,
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
func (s SQLServer) GetPersonContacts(Person *domain.Person) error {
	var contacts []domain.Contact
	rows, err := s.db.Query(selectWithPagination(s.dbname, "PersonContacts", "ID", "F_Person", strconv.Itoa(Person.Key), 0, 0))
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
func (s SQLServer) GetContact(id int) (*domain.Contact, error) {
	//[ID], [F_Contact_Type], [C_Value], [B_Primary]
	rows, err := s.db.Query(selectWithPagination(s.dbname, "Contact", "ID", "ID", strconv.Itoa(id), 0, 0))
	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return nil, err
	}
	defer rows.Close()
	var contact domain.Contact
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

		contact = domain.Contact{
			Key:   ID,
			Value: value,
			Type:  *contactType,
		}

	}
	return &contact, nil
}

//UpdatePerson обновляет данные Потребителя
func (s SQLServer) UpdatePerson(Person *domain.Person) error {
	var sex string

	if Person.Sex == true {
		sex = "1"
	} else {
		sex = "0"
	}

	_, err := s.db.Query(fmt.Sprintf("UPDATE %s.dbo.Persons"+
		" SET C_Family_Name = '%s', C_Name = '%s', C_Patronymic_Name = '%s', F_Users = '%s', D_Date_Birth = '%s', D_Date_Death = '%s', B_Sex = %s"+
		" WHERE ID =  %s",
		s.dbname, Person.FamilyName, Person.Name, Person.PatronymicName, Person.User.Key, ConvertDate(Person.DateBirth), ConvertDate(Person.DateDeath), sex, strconv.Itoa(Person.Key)))
	if err != nil {
		fmt.Println("Ошибка при обновлени Пользователя", err)
	}
	return err
}

//DeletePerson удаляет Потребителя
func (s SQLServer) DeletePerson(Person *domain.Person) error {
	_, err := s.db.Query(fmt.Sprintf("DELETE FROM %s.dbo.Persons WHERE ID =  %s",
		s.dbname, strconv.Itoa(Person.Key)))
	if err != nil {
		fmt.Println("Ошибка при удалении Пользователя", err)
	}
	return err
}
