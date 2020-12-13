package tech

import (
	"core/contract"
	"core/ref"
)

//Object Объект
type Object struct {
	Key          int
	Name         string
	BuildAddress string
	ref.Address
	SubObjects []SubObject
	Type       ObjectType
}

//SubObject Помещение в Объекте
type SubObject struct {
	Key    int
	Name   string
	Number string
}

//ObjectType Тип объекта
type ObjectType struct {
	Key      int
	Name     string
	is354    bool
	isEnergy bool
}

//SubObjectType Тип Подобъекта
type SubObjectType struct {
	Key    int
	Name   string
	isFlat bool
}

//SupplyPoint отражает Точку поставки
type SupplyPoint struct {
	Key            int
	Number         int
	resource       ref.EnergyResource
	registerPoints []contract.RegisterPoint
}
