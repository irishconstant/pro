package dependency

import (
	"repo/abstract"
	"repo/sqlserver"

	"github.com/golobby/container"
)

//GetDependency creates binding between interface and implementation
func GetDependency() abstract.DatabaseConnection {

	// Если надо изменить реализацию на другую БД, достаточно реализовать её в repo и сослаться на новую реализацию здесь. Логика не поменяется
	container.Singleton(func() abstract.DatabaseConnection {
		return &sqlserver.SQLServer{}
	})
	var dbc abstract.DatabaseConnection
	container.Make(&dbc)
	return dbc
}
