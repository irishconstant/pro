package tech

import (
	"core/contract"
	"core/ref"
)

//SupplyDistrict район теплоснабжения [dbo].[Supply_Districts]
type SupplyDistrict struct {
	Key       int
	Name      string
	Location  ref.Region
	pipelines []Edge
}

// Source теплоисточник
type Source struct {
	Key int

	Object Object
	nodes  []SourceNode // Элементы схемы энергоснабжения, непосредственно относящиеся к теплоисточнику. Например, коллекторы отпуска тепловой энергии

	SeasonMode ref.SeasonMode // Сезонность котельной: •	Круглогодичное; •	Сезонное.
	FuelType   ref.FuelType   // Вид топлива

	NormSupplyValue float32 // Нормативная подпитка тепловых сетей (м3)

	SupplierElectricity contract.LegalEntity // Организация-поставщик электрической энергии на котельную
	VoltageNominal      ref.VoltageNominal   // Уровень напряжения по договору (ВН, СН1, СН2, НН)

	TransportGas         contract.LegalEntity // Организация-транспортировщик природного газа на котельную
	SupplierGas          contract.LegalEntity // Организация-поставщик природного газа на котельную
	SupplierTechWater    contract.LegalEntity // Организация-поставщик воды на технологические нужды котельной
	SupplierHotWater     contract.LegalEntity // Организация-поставщик воды на ГВС
	SupplierCanalisation contract.LegalEntity // Организация, оказывающая услугу водоотведения на котельной
	SupplierHeat         contract.LegalEntity // Организация-поставщик покупного тепла на котельную (ЦТП)

	params []SourceParam // Утвержденные параметры с разбивкой по месяцам

	SourceFact []SourceFact // Фактические параметры работы котельной по расчётным периодам
}

// SourceParam месячный параметр котельной
type SourceParam struct {
	Month      int
	Losses     float32 // Потери  на собственные нужды, %	Для каждого месяца
	Efficiency float32 // К.П.Д. брутто котельной (%)	Для каждого месяца
}

// SourceFact параметры котельной в расчётном периоде
type SourceFact struct {
	ref.CalcPeriod
	WorkDuration           int     // Продолжительность работы источника (в часах)
	TempWater              float32 // t°х.воды
	TempAir                float32 // t°возд
	HeatDuration           int     // Отопление, час
	TempHeat               float32 // Отопление, град
	FuelConsumption        float32 // Расход натурального топлива, тыс.м3, тн
	ElectricityConsumption float32 // Эл.энергия, тыс. кВт*час
	TechWaterConstumption  float32 // Вода на технологические нужды, тыс. м3
	HotWaterConsumption    float32 // Вода на ГВС, тыс. м3
	Canalisation           float32 // Канализирование, тыс. м3
	PaidHeat               float32 // Покупное тепло, Гкал
}

// SourceNode элемент схемы энергоснабжения, задействованный в выработке (генерации)
type SourceNode struct {
	Node Node

	Load      float32 // Подключенная нагрузка, Гкал/час:
	HeatLoad  float32 //отопление
	VentLoad  float32 //вентиляция
	HWSLoad   float32 //ГВС
	SteamLoad float32 //пар

	TempGraphHP  TempGraph //Температурный график теплоснабжения в ОП
	TempGraphNHP TempGraph //Температурный график теплоснабжения в МОП

	isDevice bool // Наличие коммерческого узла учета тепловой энергии
}
