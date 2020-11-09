package model

//Customer respresents a typical person
type Customer struct {
	ID             int
	Name           string
	PatronymicName string
	FamilyName     string
}

//CustomersBook does
type CustomersBook struct {
	CustomerCount int
	Customers     []Customer
}
