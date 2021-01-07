package contract

//Division представляет из себя подразделения уровня "Отделение"
type Division struct {
	Key  int
	Name string

	// Юр.лицо
	Entity LegalEntity

	/* Идея с расчётными периодами до уровня подразделений - плохая. РП - один на организацию - правильная
	CurrentPeriod *ref.CalcPeriod
	LastPeriod    *ref.CalcPeriod
	*/
}

//Account отражает Лицевой счёт
type Account struct {
	Number         int
	RegisterPoints []*RegisterPoint
}

//RegisterPoint отражает Точку учёта
type RegisterPoint struct {
	Number int
}

/* Если бы нужно было создать только один экземпляр подразделения
var instanceDivision *Division
var once sync.Once

// GetInstance возвращает экзепляр Подразделения
func GetInstance() *Division {
	once.Do(func() {
		instanceDivision = &Division{}
	})
	return instanceDivision
}
*/
