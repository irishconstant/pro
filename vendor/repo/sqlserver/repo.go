package sqlserver

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"repo/abstract"

	// импорт пакета с драйвером для дб
	_ "github.com/denisenkom/go-mssqldb"
)

type SQLServer struct {
}

//ConnectToDatabase does
func (SQLServer) ConnectToDatabase(connectionString string) *sql.DB {
	db, err := sql.Open("mssql", connectionString)
	//fmt.Println(reflect.TypeOf(db))
	if err != nil {
		fmt.Println("Не подключается к серверу баз данных: ", err.Error())
		panic(err)
	}
	fmt.Println("Соединение с сервером баз данных установлено")
	return db
}

//CloseConnect closes connection to db
func (SQLServer) CloseConnect(db *sql.DB) {
	fmt.Println("Закрыто соединение с сервером баз данных")
	db.Close()
}

//GetConnectionParams returns params from config file
func (SQLServer) GetConnectionParams(filePath string) abstract.SQLConnect {
	fmt.Println("Получен путь конфигурационного файла:", filePath)
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		err = fmt.Errorf("Не найден конфигурационный файл")
		panic(err)
	}
	byteValue, _ := ioutil.ReadAll(file)
	var result map[string]string
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		err = fmt.Errorf("Ошибка при попытке сериализовать файл из json")
		panic(err)
	}

	fmt.Println("Получен url сервера баз данных из конфигурационного файла", result["connectionString"])
	fmt.Println("Получено наименование базы данных из кофигурационного файла", result["database"])
	var connection = abstract.SQLConnect{ConnectionString: result["connectionString"], Database: result["database"]}
	return connection
}

//CreateSelectQuery creates select all columns query
func CreateSelectQuery(database string, table string) string {
	return fmt.Sprintf("SELECT * FROM [%s].dbo.[%s]", database, table)
}

//CreateDeleteQuery creates delete all rows from table query
func CreateDeleteQuery(database string, table string) string {
	return fmt.Sprintf("DELETE FROM [%s].dbo.[%s]", database, table)
}
