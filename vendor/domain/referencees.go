package domain

// ContactType Тип контакта. Справочник
type ContactType struct {
	Key        int
	Name       string
	Validation string
}

// Citizenship Гражданство. Справочник
type Citizenship struct {
	Key  int
	Name string
}

// DocType Тип документа. Справочник
type DocType struct {
	Key            int
	Name           string
	Citizenship    Citizenship
	IsSerialNumber bool
	IsNumber       bool
	IsFromCode     bool
	IsDateBegin    bool
	IsDateEnd      bool
}

// CustomerType Тип потребителя. Справочник
type CustomerType struct {
	Key  int
	Name string
}
