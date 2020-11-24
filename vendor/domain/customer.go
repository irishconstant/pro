package domain

//Customer представляет из себя типичного Потребителя
type Customer struct {
	Key            int
	Name           string
	PatronymicName string
	FamilyName     string
	User           User // Ответственный пользователь
}

//CustomersBook представляет из себя набор Потребителей определённого Пользователя
type CustomersBook struct {
	CustomerCount int
	CurrentPage   int
	Pages         []Page
	Customers     []Customer
}
