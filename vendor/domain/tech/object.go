package tech

import "domain/ref"

//Object Объект
type Object struct {
	ref.Address
	SubObjects []SubObject
	Type       ObjectType
}

//SubObject Помещение в Объекте
type SubObject struct {
}

//ObjectType Тип объекта
type ObjectType struct {
	Name  string
	is354 bool
}

//SubObjectType Тип Подобъекта
type SubObjectType struct {
	Name   string
	isFlat bool
}

//SupplyPoint отражает Точку поставки
type SupplyPoint struct {
	Number   int
	resource ref.EnergyResource
}
