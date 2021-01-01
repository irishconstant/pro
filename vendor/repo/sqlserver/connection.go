package sqlserver

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	// импорт пакета с драйвером для дб
	_ "github.com/denisenkom/go-mssqldb"
)

//SQLServer инкапсулирует всё, что нужно для подключения
type SQLServer struct {
	db               *sql.DB // Драйвер подключения к СУБД
	dbname           string  // Имя БД из конфиг.файла 
	connectionString string  // Строка подключения из конфиг.файла
	pageSize         int
}

//ConnectToDatabase соединяет непосредственно с экземпляром СУБД
func (s *SQLServer) ConnectToDatabase() error {
	var err error
	s.db, err = sql.Open("mssql", s.connectionString)
	if err != nil {
		fmt.Println("Не подключается к серверу баз данных: ", err.Error())
		panic(err)
	}
	fmt.Println("Соединение с сервером баз данных установлено")
	return err
}

//CloseConnect закрывает соединение (используется через defer)
func (s *SQLServer) CloseConnect() error {
	var err error
	fmt.Println("Закрыто соединение с сервером баз данных")
	err = s.db.Close()
	if err != nil {
		fmt.Println("Ошибка при попытке закрыть соединение с базой данных: ", err.Error())
	}
	return err
}

//GetConnectionParams получает параметры для Системы из конфигурационного файла
func (s *SQLServer) GetConnectionParams(filePath string) error {
	fmt.Println("Получен путь конфигурационного файла:", filePath)
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		err = fmt.Errorf("Не найден конфигурационный файл")
		// TODO: Реализовать при ошибке повторный запрос пути к конфигурационному файлу в командной строке
		return err
	}
	byteValue, _ := ioutil.ReadAll(file)
	var result map[string]string
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		err = fmt.Errorf("Ошибка при попытке сериализовать файл из json")
		return err
	}

	fmt.Println("Получен url сервера баз данных из конфигурационного файла", result["connectionString"])
	fmt.Println("Получено наименование базы данных из кофигурационного файла", result["database"])
	s.connectionString = result["connectionString"]
	s.dbname = result["database"]
	return err
}
