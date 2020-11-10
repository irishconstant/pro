package sqlserver

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//CheckPassword checks password
func (s *SQLServer) CheckPassword(login string, password string) bool {

	rows, err := s.db.Query(fmt.Sprintf("SELECT TOP 1 Password FROM %s.dbo.Users WHERE Login = '%s'", s.dbname, login))
	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return false
	}
	defer rows.Close()
	var passwordDB string
	for rows.Next() {
		rows.Scan(&passwordDB)
	}

	//fmt.Println(fmt.Sprintf("Введенные логин: %s, пароль: %s", login, password))
	//fmt.Println(fmt.Sprintf("В базе данных логин: %s, пароль: %s", login, passwordDB))
	if CheckPasswordHash(password, passwordDB) {
		fmt.Println(fmt.Sprintf("Пользователь %s прошел процедуру аутентификации", login))
		return true
	}
	fmt.Println("Пароль не совпал")
	return false
}

//CreateUser creates new user in SQL Server
func (s *SQLServer) CreateUser(login string, password string) bool {

	hashedPassword, err := HashPassword(password)
	_, err = s.db.Query(fmt.Sprintf("INSERT INTO %s.dbo.Users (Login, Password) SELECT '%s', '%s'", s.dbname, login, hashedPassword))
	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return false
	}
	return true
}

//HashPassword hashes
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash checks
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
