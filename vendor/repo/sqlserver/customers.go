package sqlserver

import (
	"domain"
	"fmt"
	"strconv"
)

//GetUserCustomersPagination возвращает всех Потребителей конкретного Пользователя для страницы
func (s *SQLServer) GetUserCustomersPagination(u domain.User, currentPage int, pageSize int) (map[int]*domain.Customer, error) {
	customers := make(map[int]*domain.Customer)
	rows, err := s.db.Query(selectWithPagination(s.dbname, "UserCustomers", "ID", "F_Users", u.Key, pageSize, currentPage))

	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			a int
			b string
			c string
			d string
			e string
		)
		rows.Scan(&a, &b, &c, &d, &e)
		customer := domain.Customer{
			Key:            a,
			Name:           b,
			PatronymicName: c,
			FamilyName:     d,
			User:           u}
		if a != 0 {
			customers[a] = &customer
		}
	}

	return customers, nil
}

//GetUserCustomersAll возвращает всех Потребителей Пользователя
func (s *SQLServer) GetUserCustomersAll(u domain.User) (map[int]*domain.Customer, error) {
	customers := make(map[int]*domain.Customer)
	rows, err := s.db.Query(selectWithPagination(s.dbname, "UserCustomers", "ID", "F_Users", u.Key, 0, 0))

	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			a int
			b string
			c string
			d string
			e string
		)
		rows.Scan(&a, &b, &c, &d, &e)
		customer := domain.Customer{
			Key:            a,
			Name:           b,
			PatronymicName: c,
			FamilyName:     d,
			User:           u}
		if a != 0 {
			customers[a] = &customer
		}
	}

	return customers, nil
}

//CreateCustomer создаёт нового Потребителя
func (s SQLServer) CreateCustomer(c *domain.Customer) error {

	rows, err := s.db.Query(fmt.Sprintf("INSERT INTO Customers (C_Name, C_Family_Name, C_Patronymic_Name, F_Users) SELECT '%s', '%s', '%s', '%s' SELECT SCOPE_IDENTITY()", c.Name, c.FamilyName, c.PatronymicName, c.User.Key))
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
			Sex            bool
			DateBirth      string
			DateDeath      string
		)
		rows.Scan(&ID, &FamilyName, &Name, &PatronymicName, &UserLogin, &CitizenshipKey, &Sex, &DateBirth, &DateDeath)
		user, err := s.GetUser(UserLogin)
		if err != nil {
			fmt.Println("Ошибка c получением Пользователя: ", err)
			return nil, err
		}

		customer = domain.Customer{
			Key:            ID,
			Name:           Name,
			PatronymicName: PatronymicName,
			FamilyName:     FamilyName,
			Sex:            Sex,
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
func (s SQLServer) UpdateCustomer(customer *domain.Customer) error { // Во все подобные темы добавить передачу юзера
	_, err := s.db.Query(fmt.Sprintf("UPDATE %s.dbo.Customers SET C_Family_Name = '%s', C_Name = '%s', C_Patronymic_Name = '%s', F_Users = '%s' WHERE ID =  %s",
		s.dbname, customer.FamilyName, customer.Name, customer.PatronymicName, customer.User.Key, strconv.Itoa(customer.Key)))
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
