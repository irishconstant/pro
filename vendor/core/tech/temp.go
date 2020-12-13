package tech

import "core/ref"

// Temperatures cправочник фактических температур
type Temperatures struct {
	Key        int
	Location   ref.Address // Если температуры заданы в привязке к адресу
	Source     Node        // Если температуры заданы в привязке к теплоисточнику
	Year       int
	tempMonths []TemperaturesMonth // Температуры помесячно
	tempDaily  []TemperaturesDaily // Температуры посуточно
}

// TemperaturesSNIP справочник температура по СНиП
type TemperaturesSNIP struct {
	Key        int
	Location   ref.Address
	tempMonths []TemperaturesMonth
}

// TemperaturesMonth Значение температур помесячно
type TemperaturesMonth struct {
	Key int

	Month int

	AirTemp    int
	GroundTemp int
	WaterTemp  int
	HeatDays   int
}

// TemperaturesDaily значения температур помесячно
type TemperaturesDaily struct {
	Key int

	Month int
	Day   int

	AirTemp    int
	GroundTemp int
	WaterTemp  int
}

//TempGraph температурный график
type TempGraph struct {
	Key    int
	Name   string
	values map[int]TempGraphValues
}

//TempGraphValues записи температурного графика
type TempGraphValues struct {
	Key      int
	AirTemp  int     // Температура наружного воздуха
	DirTemp  float32 // Температура подачи
	RevTemp  float32 // Температура обратки
	HeatTemp float32 // Температура в системе отопления
	isCut    bool    // Признак того, что здесь начинается срезка
}
