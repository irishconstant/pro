package sqlserver

import (
	"domain"
	"fmt"
	"strconv"
	"time"
)

//GetUserFiltredCustomersPagination возвращает всех Потребителей конкретного Пользователя для страницы с учётом переданных фильтров
func (s *SQLServer) GetUserFiltredCustomersPagination(u domain.User, regime int, currentPage int, pageSize int, name string, familyname string, patrname string, sex string) (map[int]*domain.Customer, error) {
	customers := make(map[int]*domain.Customer)

	var query string
	if sex == "" {
		query =
			fmt.Sprintf("EXEC %s.dbo.GetFilteredPaginatedCustomers '%s', '%s', '%s', '%s', NULL, %d, %d, %d",
				s.dbname, u.Key, familyname, name, patrname, pageSize*currentPage-pageSize, pageSize, regime)
	} else {
		query =
			fmt.Sprintf("EXEC %s.dbo.GetFilteredPaginatedCustomers '%s', '%s', '%s', '%s', %s, %d, %d, %d",
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
		customer := domain.Customer{
			Key:            ID,
			FamilyName:     FamilyName,
			Name:           Name,
			PatronymicName: PatronymicName,
			Sex:            Sex,
			//DateBirth:      DateBirth,
			//DateDeath:      DateDeath,
			User: u}
		if ID != 0 {
			customers[ID] = &customer
		}
	}
	return customers, nil
}

//CreateCustomer создаёт нового Потребителя
func (s SQLServer) CreateCustomer(c *domain.Customer) error {
	var bSex int
	if strconv.FormatBool(c.Sex) == "true" {
		bSex = 1
	} else {
		bSex = 0
	}
	rows, err := s.db.Query(fmt.Sprintf("INSERT INTO Customers (C_Name, C_Family_Name, C_Patronymic_Name, B_Sex, F_Users) SELECT '%s', '%s', '%s', %d, '%s' SELECT SCOPE_IDENTITY()", c.Name, c.FamilyName, c.PatronymicName, bSex, c.User.Key))
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

//GetCustomer возвращает пользователя по первичному ключу
func (s SQLServer) GetCustomer(id int) (*domain.Customer, error) {

	rows, err := s.db.Query(selectWithPagination(s.dbname, "Customer", "ID", "ID", strconv.Itoa(id), 0, 0))

	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return nil, err
	}
	var customer domain.Customer
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

		DateBirthG, _ := time.Parse(time.RFC3339, DateBirth)
		DateDeathG, _ := time.Parse(time.RFC3339, DateDeath)

		customer = domain.Customer{
			Key:            ID,
			Name:           Name,
			PatronymicName: PatronymicName,
			FamilyName:     FamilyName,
			Sex:            sexBool,
			DateBirth:      DateBirthG,
			DateDeath:      DateDeathG,
			User:           *user}

	}

	err = s.GetCustomerContacts(&customer)

	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return nil, err
	}

	return &customer, nil
}

// GetCustomerContacts получает Контакты для Пользователя
func (s SQLServer) GetCustomerContacts(customer *domain.Customer) error {
	var contacts []domain.Contact
	rows, err := s.db.Query(selectWithPagination(s.dbname, "CustomerContacts", "ID", "F_Customer", strconv.Itoa(customer.Key), 0, 0))
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
	customer.Contacts = contacts
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

//GetContactType возвращает Тип контакта по его идентификатору
func (s SQLServer) GetContactType(id int) (*domain.ContactType, error) {
	rows, err := s.db.Query(selectWithPagination(s.dbname, "ContactType", "ID", "ID", strconv.Itoa(id), 0, 0))

	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return nil, err
	}
	defer rows.Close()
	var contactType domain.ContactType
	for rows.Next() {
		var (
			a int
			b string
		)
		rows.Scan(&a, &b)
		if err != nil {
			fmt.Println("Ошибка c получением Типа контакта: ", err)
			return nil, err
		}

		contactType = domain.ContactType{
			Key:  a,
			Name: b,
		}
	}

	return &contactType, nil
}

//UpdateCustomer обновляет данные Потребителя
func (s SQLServer) UpdateCustomer(customer *domain.Customer) error {
	var sex string
	if customer.Sex == true {
		sex = "1"
	} else {
		sex = "0"
	}

	_, err := s.db.Query(fmt.Sprintf("UPDATE %s.dbo.Customers"+
		" SET C_Family_Name = '%s', C_Name = '%s', C_Patronymic_Name = '%s', F_Users = '%s', D_Date_Birth = '%s', D_Date_Death = '%s', B_Sex = '%s'"+
		" WHERE ID =  %s",
		s.dbname, customer.FamilyName, customer.Name, customer.PatronymicName, customer.User.Key, ConvertDate(customer.DateBirth), ConvertDate(customer.DateDeath), sex, strconv.Itoa(customer.Key)))
	if err != nil {
		fmt.Println("Ошибка при обновлени Пользователя", err)
	}
	return err
}

//DeleteCustomer удаляет Потребителя
func (s SQLServer) DeleteCustomer(customer *domain.Customer) error {
	_, err := s.db.Query(fmt.Sprintf("DELETE FROM %s.dbo.Customers WHERE ID =  %s",
		s.dbname, strconv.Itoa(customer.Key)))
	if err != nil {
		fmt.Println("Ошибка при удалении Пользователя", err)
	}
	return err
}
