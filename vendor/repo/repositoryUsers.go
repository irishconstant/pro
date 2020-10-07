package repo

import (
	"database/sql"
	"domain"
	"fmt"
)

//GetUsers return all users from databases
func GetUsers(db *sql.DB, dbname string) map[int]*domain.User {
	db.Ping()
	fmt.Println("Получение списка Пользователей")
	users := make(map[int]*domain.User)
	query := fmt.Sprintf("SELECT Id, C_Name, C_Patronymic_Name, C_Family_Name AS name FROM [%s].dbo.Users", dbname)

	rows, err := db.Query(query)

	if err != nil {
		fmt.Println("Ошибка c запросом", err)
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

func AddUser() {

}

func DeleteUser() {

}

func EditUser() {

}

func AddContactToUser() {

}

func DeletContactToUser() {

}
