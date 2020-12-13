package tech

import (
	"core/contract"
	"core/ref"
)

//Object Объект
type Object struct {
	ref.Address
	SubObjects []SubObject
	Type       ObjectType
}

//SubObject Помещение в Объекте
type SubObject struct {
	Name   string
	Number string
}

//ObjectType Тип объекта
type ObjectType struct {
	Name     string
	is354    bool
	isEnergy bool
}

//SubObjectType Тип Подобъекта
type SubObjectType struct {
	Name   string
	isFlat bool
}

//SupplyPoint отражает Точку поставки
type SupplyPoint struct {
	Number         int
	resource       ref.EnergyResource
	registerPoints []contract.RegisterPoint
}
