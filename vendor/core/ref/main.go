package ref

import "time"

// CalcPeriod расчётный период
type CalcPeriod struct {
	Key        int
	Name       string
	Year       int
	Month      int
	IsCurrent  bool      // Признак того, что это текущий расчётный период
	IsSelected bool      // Признак того, что этот расчётный период выбран
	DateClose  time.Time // Дата закрытия расчётного периода (для уже закрытых)
}
