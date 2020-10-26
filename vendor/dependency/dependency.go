package dependency

import (
	"repo/abstract"
	"repo/sqlserver"

	"github.com/golobby/container"
)

func GetDependency() abstract.DatabaseConnection {
	container.Singleton(func() abstract.DatabaseConnection {
		return &sqlserver.SQLServer{}
	})
	var dbc abstract.DatabaseConnection
	container.Make(&dbc)
	return dbc
}
