package tech

import (
	"domain/ref"
	"sync"
	"time"
)

//SupplyDistrict район теплоснабжения
type SupplyDistrict struct {
	Key    int
	Name   string
	Region ref.Region
}

//Source источник энергоснабжения (теплоснабжения, электроснабжения, водоснабжения). По сути набор вершин, соединенных гранями.
type Source struct {
	Key   int
	nodes []*Node
	lock  sync.RWMutex
}

//Node элементы схемы энергоснабжения. По сути - вершина (node) графа
type Node struct {
	Key    int
	Name   string
	Object Object
}

//NodeType Тип элемента схемы энергоснабжения
type NodeType struct {
	Key  int
	Name string
}

//Edge линия энергоснабжения. По сути - грань (edge) графа.
type Edge struct {
	Key int
}

//EdgeOperation переключения. Включения/отключения/ограничения линии
type EdgeOperation struct {
	Key   int
	date  time.Time
	value int // Процент активности линии
}

//HeatGraph температурный график
type HeatGraph struct {
	Name   string
	values map[int]heatGraphValues
}

//heatGraphValues записи температурного графика
type heatGraphValues struct {
	AirTemp  int     // Температура наружного воздуха
	DirTemp  float32 // Температура подачи
	RevTemp  float32 // Температура обратки
	HeatTemp float32 // Температура в системе отопления
}
