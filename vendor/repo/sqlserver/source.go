package sqlserver

import (
	"auth"
	"core/tech"
	"fmt"
	"strconv"
)

//GetAllSources возвращает все Источники
func (s SQLServer) GetAllSources(regime int, currentPage int, pageSize int) (map[int]*tech.Source, error) {
	Sources := make(map[int]*tech.Source)
	var query string
	query = fmt.Sprintf("EXEC %s.dbo.GetFilteredPaginatedSources NULL, NULL, NULL, NULL, NULL, %d, %d, %d",
		s.dbname, pageSize*currentPage-pageSize, pageSize, regime)

	rows, err := s.db.Query(query)

	if err != nil {
		fmt.Println("Ошибка с запросом в GetAllSources", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ID int
		)
		rows.Scan(&ID)
		newSource, _ := s.GetSource(ID)

		if ID != 0 {
			Sources[ID] = newSource
		}
	}
	return Sources, nil
}

//GetSource возвращает Источник по первичному ключу
func (s SQLServer) GetSource(id int) (*tech.Source, error) {

	query := creatorSelect(s.dbname, "Source", "ID", "ID", strconv.Itoa(id))

	rows, err := s.db.Query(query)

	if err != nil {
		fmt.Println("Ошибка c запросом в GetSource: ", err)
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
			name                 string
		)
		rows.Scan(
			&ID,
			&name,
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

		fuelTypeI, err := s.GetFuelType(fuelType)
		seasonModeI, err := s.GetSeasonMode(seasonMode)
		objectI, err := s.GetObject(object)

		/*


			supplierElectricityI, err := s.GetEntity(supplierElectricity)
			transportGasI, err := s.GetEntity(transportGas)
			supplierGasI, err := s.GetEntity(supplierGas)
			supplierTechWaterI, err := s.GetEntity(supplierTechWater)
			supplierHotWaterI, err := s.GetEntity(supplierHotWater)
			supplierCanalisationI, err := s.GetEntity(supplierCanalisation)
			supplierHeatI, err := s.GetEntity(supplierHeat)
		*/
		if err != nil {
			fmt.Println("Ошибка c запросом: ", err)
			return nil, err
		}

		Source := tech.Source{
			Key:             ID,
			Name:            name,
			NormSupplyValue: normSupplyValue,
			FuelType:        *fuelTypeI,
			SeasonMode:      *seasonModeI,
			Object:          *objectI,
		}

		/*
			Source := tech.Source{
				Key:                 ID,
				Name:                name,
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
		*/

		//TODO: params []SourceParam
		//TODO:	SourceFact []SourceFact

		return &Source, nil
	}

	return nil, err
}

//GetSourceQuantityFiltered возвращает КОЛИЧЕСТВО источников с учётом переданных фильтров
func (s *SQLServer) GetSourceQuantityFiltered(u auth.User, name string) (int, error) {
	var query string
	query = fmt.Sprintf("EXEC %s.dbo.GetQuantityFilteredSources  NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0", s.dbname)
	rows, err := s.db.Query(query)

	if err != nil {
		fmt.Println("Ошибка c запросом в GetSourceQuantityFiltered: ", query, err)
		return 0, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			Num int
		)
		rows.Scan(
			&Num)
		return Num, nil

	}

	return 0, err
}
