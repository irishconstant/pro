package sqlserver

import (
	"core/tech"
	"fmt"
	"strconv"
)

/*
//GetAllSource возвращает все Источники
func (s SQLServer) GetAllSource() (map[int]*tech.Source, error) {

}
*/

//GetSource возвращает Источник по первичному ключу
func (s SQLServer) GetSource(id int) (*tech.Source, error) {

	rows, err := s.db.Query(selectWithPagination(s.dbname, "Source", "ID", "ID", strconv.Itoa(id), 0, 0))

	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ID                   int
			object               int
			seasonMode           int
			fuelType             int
			normSupplyValue      float32
			supplierElectricity  int
			voltageNominal       float32
			transportGas         int
			supplierGas          int
			supplierTechWater    int
			supplierHotWater     int
			supplierCanalisation int
			supplierHeat         int
		)
		rows.Scan(&ID,
			&object,
			&seasonMode,
			&fuelType,
			&normSupplyValue,
			&supplierElectricity,
			&voltageNominal,
			&transportGas,
			&supplierGas,
			&supplierTechWater,
			&supplierHotWater,
			&supplierCanalisation,
			&supplierHeat,
		)

		objectI, err := s.GetObject(object)
		seasonModeI, err := s.GetSeasonMode(seasonMode)
		fuelTypeI, err := s.GetFuelType(fuelType)
		supplierElectricityI, err := s.GetEntity(supplierElectricity)
		transportGasI, err := s.GetEntity(transportGas)
		supplierGasI, err := s.GetEntity(supplierGas)
		supplierTechWaterI, err := s.GetEntity(supplierTechWater)
		supplierHotWaterI, err := s.GetEntity(supplierHotWater)
		supplierCanalisationI, err := s.GetEntity(supplierCanalisation)
		supplierHeatI, err := s.GetEntity(supplierHeat)

		if err != nil {
			fmt.Println("Ошибка c запросом: ", err)
			return nil, err
		}

		Source := tech.Source{
			Key:                 ID,
			Object:              *objectI,
			FuelType:            *fuelTypeI,
			SeasonMode:          *seasonModeI,
			NormSupplyValue:     normSupplyValue,       // Нормативная подпитка тепловых сетей (м3)
			SupplierElectricity: *supplierElectricityI, // Организация-поставщик электрической энергии на котельную
			//	VoltageNominal      ref.VoltageNominal   // Уровень напряжения по договору (ВН, СН1, СН2, НН)
			TransportGas:         *transportGasI,         // Организация-транспортировщик природного газа на котельную
			SupplierGas:          *supplierGasI,          // Организация-поставщик природного газа на котельную
			SupplierTechWater:    *supplierTechWaterI,    // Организация-поставщик воды на технологические нужды котельной
			SupplierHotWater:     *supplierHotWaterI,     // Организация-поставщик воды на ГВС
			SupplierCanalisation: *supplierCanalisationI, // Организация, оказывающая услугу водоотведения на котельной
			SupplierHeat:         *supplierHeatI,         // Организация-поставщик покупного тепла на котельную (ЦТП)
		}

		//TODO: params []SourceParam // Утвержденные параметры с разбивкой по месяцам
		//TODO:	SourceFact []SourceFact

		return &Source, nil
	}

	return nil, err
}
