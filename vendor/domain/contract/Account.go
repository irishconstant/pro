package contract

import "domain/tech"

//Account отражает Лицевой счёт
type Account struct {
	Number         int
	RegisterPoints []RegisterPoint
}

//RegisterPoint отражает Точку учёта
type RegisterPoint struct {
	Number int
	tech.SupplyPoint
}
