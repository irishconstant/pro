package sqlserver

import (
	"fmt"
	"model"

	"golang.org/x/crypto/bcrypt"
)

//CheckPassword проверят пароль
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

	if CheckPasswordHash(password, passwordDB) {
		fmt.Println(fmt.Sprintf("Успешная аутентификация пользователя %s", login))
		return true
	}
	fmt.Println(fmt.Sprintf("Провалена аутентификация пользователя %s", login))
	return false
}

//CreateUser создает нового Пользователя
func (s *SQLServer) CreateUser(login string, password string) bool {

	hashedPassword, err := HashPassword(password)
	_, err = s.db.Query(fmt.Sprintf("INSERT INTO %s.dbo.Users (Login, Password) SELECT '%s', '%s'", s.dbname, login, hashedPassword)) // Создаём самого пользователя
	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return false
	}

	return true
}

//CreateUserWithRoles создает нового Пользователя с ролью
func (s *SQLServer) CreateUserWithRoles(u model.User) bool {

	hashedPassword, err := HashPassword(u.Password)
	_, err = s.db.Query(fmt.Sprintf("INSERT INTO %s.dbo.Users (Login, Password, C_Family_Name, C_Name) SELECT '%s', '%s', '%s', '%s'",
		s.dbname, u.Key, hashedPassword, u.FamilyName, u.Name)) // Создаём самого пользователя
	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return false
	}

	for _, value := range u.Roles {
		_, err = s.db.Query(fmt.Sprintf("INSERT INTO %s.dbo.User_Roles (F_Users, F_Roles) SELECT '%s', '%s'",
			s.dbname, fmt.Sprint(value.Key), value.Name))
	}
	return true
}

//HashPassword хэширует пароль
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash проверят соответствие пароля и хэша
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GetUserRoles возвращает все роли Пользователя из БД
func (s *SQLServer) GetUserRoles(user *model.User) (*model.User, error) {

	roles := make(map[int]*model.Role)
	rows, err := s.db.Query(fmt.Sprintf("SELECT Id, Name FROM [%s].dbo.Users WHERE F_Users = '%s'", s.dbname, user.Key))

	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			a int
			b string
		)
		rows.Scan(&a, &b)
		role := model.Role{Key: a, Name: b}
		if a != 0 {
			roles[a] = &role
		}
	}
	user.Roles = roles
	return user, nil
}
