package tech

import (
	"core/ref"
	"time"
)

//Pipeline трубопроводы линии теплоснабжения
type Pipeline struct {
	Key int

	Year        int       // Год прокладки
	DateBegin   time.Time // Дата начала работы (обычно совпадает с годом прокладки)
	DateEnd     time.Time // Дата окончания работы
	LayerType             // Тип прокладки
	TempProject int       // Температура проектирования

	NetworkType                  // Исполнение сети: однотрубная или двухтрубная
	DiameterDirect  ref.Diameter // Диаметр подающего трубопровода
	DiamtereReverse ref.Diameter // Диаметр обратного трубопровода
	LengthDirect    float32      // Длина трубопровода подающего
	LengthReverse   float32      // Длина трубопровода обратного

	CalcType // Способ расчёта темп.коэффициента (ограничить в зависимости от исполнения сети!!!)

	TempGraphHP  TempGraph // Температурный график (ОП) (ограничить в зависимости от способа расчёта темп.коэффициента)
	TempGraphNHP TempGraph // Температурный график (МОП)

	IsolationType     //Теплоизоляционный материал
	Thickness     int // Толщина теплоизоляции

	NetworkPurposes // Назначение сети
}

//LayerType Тип прокладки
type LayerType struct {
	Key  int
	Name string
}

//IsolationType Тип изоляции
type IsolationType struct {
	Key  int
	Name string
}

//CalcType способ расчёта темп. коэффициента
type CalcType struct {
	Key  int
	Name string
}

//NetworkType Исполнение сети
type NetworkType struct {
	Key  int
	Name string
}

//NetworkPurposes Назначение сети
type NetworkPurposes struct {
	Key  int
	Name string
}
