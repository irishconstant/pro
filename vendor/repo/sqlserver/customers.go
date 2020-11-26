package sqlserver

import (
	"domain"
	"fmt"
	"strconv"
)

//GetUserCustomersPagination возвращает всех Потребителей конкретного Пользователя для страницы
func (s *SQLServer) GetUserCustomersPagination(u domain.User, currentPage int, pageSize int) (map[int]*domain.Customer, error) {
	customers := make(map[int]*domain.Customer)
	rows, err := s.db.Query(selectWithPagination(s.dbname, "Customers", "ID", "F_Users", u.Key, pageSize, currentPage))

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
	rows, err := s.db.Query(selectWithPagination(s.dbname, "Customers", "ID", "F_Users", u.Key, 0, 0))

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

	return nil
}

//GetCustomer возвращает пользователя по первичному ключу
func (s SQLServer) GetCustomer(id int) (*domain.Customer, error) {
	rows, err := s.db.Query(selectWithPagination(s.dbname, "Customers", "ID", "ID", strconv.Itoa(id), 0, 0))

	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return nil, err
	}
	var customer domain.Customer
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
		user, err := s.GetUser(e)
		if err != nil {
			fmt.Println("Ошибка c получением Пользователя: ", err)
			return nil, err
		}

		customer = domain.Customer{
			Key:            a,
			Name:           b,
			PatronymicName: c,
			FamilyName:     d,
			User:           *user}

	}

	return &customer, nil
}

//UpdateCustomer обновляет данные Потребителя
func (s SQLServer) UpdateCustomer(customer *domain.Customer) error {
	_, err := s.db.Query(fmt.Sprintf("UPDATE %s.dbo.Customers SET C_Family_Name = '%s', C_Name = '%s', C_Patronymic_Name = '%s', F_Users = '%s' WHERE ID =  %s",
		s.dbname, customer.FamilyName, customer.Name, customer.PatronymicName, customer.User.Key, strconv.Itoa(customer.Key)))
	if err != nil {
		fmt.Println("Ошибка при обновлени Пользователя", err)
	}
	return err
}
