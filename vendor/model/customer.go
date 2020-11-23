package model

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
	Pages         []int // Приходится хранить слайс данных для генерации из шаблона
	Customers     []Customer
}

//Для генерации последовательности страниц
func (c CustomersBook) MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
