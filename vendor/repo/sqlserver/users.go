package sqlserver

import (
	"fmt"
	"model"
	"strconv"

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

//GetPossibleRoles возвращает роли, которые может присваивать Пользователь с определенной ролью (неавторизованный = Гость)
func (s *SQLServer) GetPossibleRoles(model.Role) map[int]model.Role {
	var roles = make(map[int]model.Role) ///!!!
	return roles                         ///!!!
}

//GetAllRoles возвращает все роли, возможные в Системе
func (s *SQLServer) GetAllRoles() map[int]*model.Role {
	var roles = make(map[int]*model.Role)

	rows, err := s.db.Query(fmt.Sprintf("SELECT ID, C_Name, B_Admin_Only FROM %s.dbo.Roles", s.dbname))
	defer rows.Close()
	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return nil
	}

	for rows.Next() {
		var (
			Key       int
			Name      string
			AdminOnly bool
		)

		rows.Scan(&Key, &Name, &AdminOnly)
		role := model.Role{
			Key:       Key,
			Name:      Name,
			AdminOnly: AdminOnly}
		if Key != 0 {
			roles[Key] = &role
		}
	}
	return roles
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

	// Создаём самого пользователя
	hashedPassword, err := HashPassword(u.Password)
	_, err = s.db.Query(fmt.Sprintf("INSERT INTO %s.dbo.Users (Login, Password, C_Family_Name, C_Name) SELECT '%s', '%s', '%s', '%s'",
		s.dbname, u.Key, hashedPassword, u.FamilyName, u.Name))
	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return false
	}

	// Создаём его роль (сделано через отдельную таблицу. Т.к. заложено на будущее. когда ролей будет много)
	_, err = s.db.Query(fmt.Sprintf("INSERT INTO %s.dbo.User_Roles (F_User, F_Roles) SELECT '%s', %s",
		s.dbname, u.Key, strconv.FormatInt(int64(u.Role.Key), 10)))
	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return false
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

	rows, err := s.db.Query(fmt.Sprintf("SELECT TOP 1 r.ID, r.C_Name FROM [%s].dbo.User_Roles AS ur INNER JOIN [%s].dbo.Roles AS r ON r.ID = ur.F_Roles  WHERE ur.F_Users = '%s'",
		s.dbname, s.dbname, user.Key))

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
		user.Role = &role
	}

	// Получаем возможности для роли Пользователя
	s.GetRoleAbilities(user.Role)

	return user, nil
}

//GetRoleByID возвращает роль и её возможности по идентификатору
func (s *SQLServer) GetRoleByID(id int) (*model.Role, error) {
	// Получаем роль из БД
	rows, err := s.db.Query(fmt.Sprintf("SELECT r.ID, r.C_Name [%s].dbo.Roles WHERE r.ID = %s",
		s.dbname, strconv.FormatInt(int64(id), 10)))

	var role model.Role

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
		role = model.Role{Key: a, Name: b}
	}

	s.GetRoleAbilities(&role)

	return &role, nil
}

// GetRoleAbilities получает данные о возможностях роли
func (s SQLServer) GetRoleAbilities(role *model.Role) (bool, error) {

	rows, err := s.db.Query(fmt.Sprintf("SELECT a.ID, a.C_Name, ar.B_Create, ar.B_Read, ar.B_Update, ar.B_Delete FROM [%s].dbo.Area_Roles AS ar INNER JOIN [%s].dbo.Areas AS a ON a.ID = ar.F_Areas WHERE ar. F_Roles = %s",
		s.dbname, s.dbname, strconv.FormatInt(int64(role.Key), 10)))

	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return false, err
	}
	defer rows.Close()

	createMap := make(map[int]*model.Area)
	readMap := make(map[int]*model.Area)
	updateMap := make(map[int]*model.Area)
	deleteMap := make(map[int]*model.Area)

	for rows.Next() {
		var (
			ID     int
			name   string
			create bool
			read   bool
			update bool
			delete bool
		)
		rows.Scan(&ID, &name, &create, &read, &update, &delete)

		area := model.Area{
			Key:  ID,
			Name: name,
		}

		// Код ниже выглядит странно, но там приходится делать, т.к. при изначальной инициализации структуры создаётся nil мапа
		if create {
			createMap[area.Key] = &area
		}
		if read {
			readMap[area.Key] = &area
		}
		if update {
			updateMap[area.Key] = &area
		}
		if delete {
			deleteMap[area.Key] = &area
		}

	}
	role.CreateAbility = createMap
	role.ReadAbility = readMap
	role.UpdateAbility = updateMap
	role.DeleteAbility = deleteMap

	return true, nil
}
