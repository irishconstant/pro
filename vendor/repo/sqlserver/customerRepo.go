package sqlserver

import (
	"fmt"
	"model"
	"repo"
)

//GetCustomers return all users from databases
func (s *SQLServer) GetCustomers() (map[int]*model.Customer, error) {
	fmt.Println("Получение списка Пользователей")

	customers := make(map[int]*model.Customer)
	rows, err := s.db.Query(repo.CreateSelectQuery(s.dbname, "Customers"))

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
		)
		rows.Scan(&a, &b, &c, &d)
		customer := model.Customer{a, b, c, d}
		if a != 0 {
			customers[a] = &customer
		}
	}

	return customers, nil
}
