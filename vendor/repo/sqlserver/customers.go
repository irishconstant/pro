package sqlserver

import (
	"domain"
	"fmt"
)

//GetUserCustomersPagination возвращает всех Потребителей конкретного Пользователя для страницы
func (s *SQLServer) GetUserCustomersPagination(u domain.User, currentPage int, pageSize int) (map[int]*domain.Customer, error) {
	customers := make(map[int]*domain.Customer)
	rows, err := s.db.Query(selectWithPagination(s.dbname, "Customers", "ID", u.Key, pageSize, currentPage))

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
	rows, err := s.db.Query(selectWithPagination(s.dbname, "Customers", "ID", u.Key, 0, 0))

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

	fmt.Println((fmt.Sprintf("INSERT INTO Customers (C_Name, C_Family_Name, C_Patronymic_Name) SELECT '%s', '%s', '%s'", c.Name, c.FamilyName, c.PatronymicName)))
	return nil
}
