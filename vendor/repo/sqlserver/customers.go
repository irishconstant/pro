package sqlserver

import (
	"fmt"
	"domain"
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

func selectWithPagination(databaseName string, tableName string, orderParam string, whereParam string, pageSize int, currentPage int) string {
	switch tableName {
	case "Customers":
		if currentPage <= 0 {
			return fmt.Sprintf("SELECT ID, C_Name, C_Patronymic_Name, C_Family_Name, F_Users FROM [%s].dbo.Customers WHERE F_Users = '%s' ORDER BY %s",
				databaseName, whereParam, orderParam)
		}
		return fmt.Sprintf("SELECT ID, C_Name, C_Patronymic_Name, C_Family_Name, F_Users FROM [%s].dbo.Customers WHERE F_Users = '%s' ORDER BY %s"+
			" OFFSET %d ROWS  FETCH NEXT %d ROWS ONLY", databaseName, whereParam, orderParam, pageSize*currentPage-pageSize, pageSize)
	case "Users":
		return fmt.Sprintf("SELECT u.C_Name, u.C_Family_Name FROM [%s].dbo.Users AS u WHERE u.Login = '%s'",
			databaseName, whereParam)
	}
	return ""
}

//CreateCustomer создаёт нового Потребителя
func (s SQLServer) CreateCustomer(u domain.Customer) int {
	return 1
}
