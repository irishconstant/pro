package main

import (
	"controller"
	"repo/abstract"
	"repo/sqlserver"

	"github.com/golobby/container"
)

func main() {

	/*
		Всё началось тогда, когда были выкованы мега-кольца
		Три первых кольца задарили бесссмертным эльфам - чисто для проверки, не передохнут ли
		Семь - коротышкам из подземных канализаций
		Ну а девять колец задарили расе людей. И (как показала практика) напрасно...
	*/

	dbc := getDependency()
	dbc.GetConnectionParams("config.ini")
	dbc.ConnectToDatabase()
	defer dbc.CloseConnect()
	controller.Router(dbc)

	//TODO: dep init
}

//getDependency создаёт привязку между интерфейсом и реализацией (IoC)
func getDependency() abstract.DatabaseConnection {

	// Если надо изменить реализацию на другую БД, достаточно реализовать её в repo и сослаться на новую реализацию здесь
	container.Singleton(func() abstract.DatabaseConnection {
		return &sqlserver.SQLServer{}
	})
	var dbc abstract.DatabaseConnection
	container.Make(&dbc)
	return dbc
}
