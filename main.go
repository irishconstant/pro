package main

import (
	"controller"
	"dependency"
)

func main() {

	/*
		Всё началось тогда, когда были выкованы мега-кольца
		Три первых кольца задарили бесссмертным эльфам - чисто для проверки, не передохнут ли
		Семь - коротышкам из подземных канализаций
		Ну а девять расе людей. И (как показала практика) напрасно...
	*/

	dbc := dependency.GetDependency()
	dbc.GetConnectionParams("config.ini")
	dbc.ConnectToDatabase()
	defer dbc.CloseConnect()
	controller.Router(dbc)

	//TODO: dep init
}
