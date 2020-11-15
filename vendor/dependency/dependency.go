package dependency

import (
	"repo/abstract"
	"repo/sqlserver"

	"github.com/golobby/container"
)

//GetDependency создаёт привязку между интерфейсом и реализацией (IoC)
func GetDependency() abstract.DatabaseConnection {

	// Если надо изменить реализацию на другую БД, достаточно реализовать её в repo и сослаться на новую реализацию здесь
	container.Singleton(func() abstract.DatabaseConnection {
		return &sqlserver.SQLServer{}
	})
	var dbc abstract.DatabaseConnection
	container.Make(&dbc)
	return dbc
}
