package contract

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

// EntityType Тип ОПФ. Справочник
type EntityType struct {
	Key  int
	Name string
}

// ContactType Тип контакта. Справочник
type ContactType struct {
	Key        int
	Name       string
	Validation string
	IsAddress  bool
}

