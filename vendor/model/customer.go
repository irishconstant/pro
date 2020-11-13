package model

//Customer respresents a typical person TODO: Переделать в лицевые счета
type Customer struct {
	Key            int
	Name           string
	PatronymicName string
	FamilyName     string
	User           User // Ответственный пользователь
}

//CustomersBook does
type CustomersBook struct {
	CustomerCount int
	Customers     []Customer
}
