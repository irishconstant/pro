package domain

// Address адрес
type Address struct {
	Address  string
	Region   Region
	District District
	City     City
	Town     Town
	Street   Street
	House    string
	Building string
	FIAS     string
}

// Region Регион
type Region struct {
	FIAS  string
	Value string
}

// District район
type District struct {
	FIAS  string
	Value string
}

// City город
type City struct {
	FIAS  string
	Value string
}

// Town населенный пункт
type Town struct {
	FIAS  string
	Value string
}

//Street улица
type Street struct {
	FIAS  string
	Value string
}
