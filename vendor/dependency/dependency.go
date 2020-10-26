package dependency

import (
	"repo/abstract"
	"repo/sqlserver"

	"github.com/golobby/container"
)

func GetDependency() {
	container.Singleton(func() abstract.DatabaseConnection {
		return &sqlserver.SQLServer{}
	})
}
