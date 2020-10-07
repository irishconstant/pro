package repo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	// импорт пакета с драйвером для дб
	_ "github.com/denisenkom/go-mssqldb"
)

//ConnectToDatabase openes connection to database with connection string from configuration files
type ConnectToDatabase interface {
	ConnectToDatabase(filePath string) error

	//Авторизация является функцией определения прав доступа к ресурсам и управления этим доступом.
	Autorisation(user string, password string) string

	//Аутентификация (от греческого: αυθεντικός ; реальный или подлинный): подтверждение подлинности чего-либо или кого либо.
	//Например, предъявление паспорта - это подтверждение подлинности заявленного имени отчества.
	Authentication(user string, password string) string
}

//SQLConnect parameters
type SQLConnect struct {
	//comment
	ConnectionString string
	Database         string
}

//ConnectToDatabase does
func (s SQLConnect) ConnectToDatabase() *sql.DB {
	db, err := sql.Open("mssql", s.ConnectionString)
	//fmt.Println(reflect.TypeOf(db))
	if err != nil {
		fmt.Println("Не подключается к серверу баз данных: ", err.Error())
		panic(err)
	}
	fmt.Println("Соединение с сервером баз данных установлено")
	return db
}

//CloseConnect closes connection to db
func CloseConnect(db *sql.DB) {
	fmt.Println("Закрыто соединение с сервером баз данных")
	db.Close()
}

//GetConnectionParams returns params from config file
func GetConnectionParams(filePath string) SQLConnect {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		err = fmt.Errorf("Не найден конфигурационный файл")
		panic(err)
	}
	byteValue, _ := ioutil.ReadAll(file)
	var result map[string]string
	json.Unmarshal([]byte(byteValue), &result)
	fmt.Println("Получен url сервера баз данных из конфигурационного файла", result["connectionString"])
	fmt.Println("Получено наименование базы данных из кофигурационного файла", result["database"])
	var connection = SQLConnect{ConnectionString: result["connectionString"], Database: result["database"]}
	return connection
}
