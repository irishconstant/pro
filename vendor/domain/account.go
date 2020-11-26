package domain

//Account отражает Лицевой счёт
type Account struct {
	Number         int
	RegisterPoints []RegisterPoint
}

//RegisterPoint отражает Точку учёта
type RegisterPoint struct {
	Number int
	SupplyPoint
}

//SupplyPoint отражает Точку поставки
type SupplyPoint struct {
	Number int
}
