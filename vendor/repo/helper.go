package repo

import "fmt"

//CreateSelectQuery creates select all columns query
func CreateSelectQuery(database string, table string) string {

	return fmt.Sprintf("SELECT * FROM [%s].dbo.[%s]", database, table)
}

//CreateDeleteQuery creates delete all rows from table query
func CreateDeleteQuery(database string, table string) string {
	return fmt.Sprintf("DELETE FROM [%s].dbo.[%s]", database, table)
}
