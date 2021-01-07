package ref

//EnergyResource Коммунальный ресурс
type EnergyResource struct {
	Key  int
	Name string
}

/*
FuelType Вид топлива
•	Дизельное;
•	Мазут;
•	Природный газ;
•	Уголь;
•	Печное.
*/
type FuelType struct {
	Key         int
	Name        string
	BurningHeat float32 //Теплота сгорания топлива Qнр
}

/*
SeasonMode Сезонность котельной:
•	Круглогодичное;
•	Сезонное.
*/
type SeasonMode struct {
	Key  int
	Name string
}

// VoltageNominal номинал напряжения по договору (ВН, СН1, СН2, НН)
type VoltageNominal struct {
	Key  int
	Name string
}

//Diameter справочник диаметров
type Diameter struct {
	Key         int
	Name        string
	Value       int     // Условный диаметр
	WaterVolume float32 // Удельный объем воды
}
