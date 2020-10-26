package sqlserver

import (
	"database/sql"
	"domain"
	"fmt"
)

//GetUsers return all users from databases
func (SQLServer) GetUsers(db *sql.DB, dbname string) map[int]*domain.User {
	db.Ping()
	fmt.Println("Получение списка Пользователей")
	users := make(map[int]*domain.User)
	rows, err := db.Query(CreateSelectQuery(dbname, "Users"))

	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
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
		user := domain.User{a, b, c, d}
		if a != 0 {
			users[a] = &user
		}
	}

	return users
}
