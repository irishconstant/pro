package domain

//Object Объект
type Object struct {
	Address
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
