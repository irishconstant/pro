package sqlserver

import (
	"fmt"
	"model"
)

//GetUserCustomers возвращает всех Потребителей конкретного Пользователя
func (s *SQLServer) GetUserCustomers(u model.User) (map[int]*model.Customer, error) {
	customers := make(map[int]*model.Customer)
	rows, err := s.db.Query(fmt.Sprintf("SELECT ID, C_Name, C_Patronymic_Name, C_Family_Name, F_Users FROM [%s].dbo.Customers WHERE F_Users = '%s'", s.dbname, u.Key))

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
		customer := model.Customer{
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
