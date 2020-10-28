package sqlserver

import (
	"fmt"
	"model"
	"repo"
)

//GetUsers return all users from databases
func (s *SQLServer) GetUsers() (map[int]*model.User, error) {
	fmt.Println("Получение списка Пользователей")

	users := make(map[int]*model.User)
	rows, err := s.db.Query(repo.CreateSelectQuery(s.dbname, "Users"))

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
		user := model.User{a, b, c, d}
		if a != 0 {
			users[a] = &user
		}
	}

	return users, nil
}
