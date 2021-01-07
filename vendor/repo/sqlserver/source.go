package sqlserver

import (
	"core/auth"
	"core/ref"
	"core/tech"
	"fmt"
	"strconv"
)

//GetAllSources возвращает все Источники
func (s SQLServer) GetAllSources(regime int, currentPage int, pageSize int, name string, address string, seasonMode int, fuelType int, period *ref.CalcPeriod) (map[int]*tech.Source, error) {
	Sources := make(map[int]*tech.Source)
	var query string
	query = fmt.Sprintf("EXEC %s.dbo.GetFilteredPaginatedSources '%s','%s', %d, %d, NULL, NULL, %d, %d, %d",
		s.dbname, name, address, seasonMode, fuelType, pageSize*currentPage-pageSize, pageSize, regime)

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
		newSource, _ := s.GetSource(ID, period)

		if ID != 0 {
			Sources[ID] = newSource
		}
	}
	return Sources, nil
}

//GetSource возвращает Источник по первичному ключу
func (s SQLServer) GetSource(id int, period *ref.CalcPeriod) (*tech.Source, error) {

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
			divisionID           int
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
			&divisionID,
		)

		fuelTypeI, err := s.GetFuelType(fuelType)
		seasonModeI, err := s.GetSeasonMode(seasonMode)
		objectI, err := s.GetObject(object)
		division, err := s.GetDivision(divisionID)
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
			Division:        division,
			/*
				SupplierElectricity: *supplierElectricityI, // Организация-поставщик электрической энергии на котельную
				//	VoltageNominal      ref.VoltageNominal   // Уровень напряжения по договору (ВН, СН1, СН2, НН)
				TransportGas:         *transportGasI,         // Организация-транспортировщик природного газа на котельную
				SupplierGas:          *supplierGasI,          // Организация-поставщик природного газа на котельную
				SupplierTechWater:    *supplierTechWaterI,    // Организация-поставщик воды на технологические нужды котельной
				SupplierHotWater:     *supplierHotWaterI,     // Организация-поставщик воды на ГВС
				SupplierCanalisation: *supplierCanalisationI, // Организация, оказывающая услугу водоотведения на котельной
				SupplierHeat:         *supplierHeatI,         // Организация-поставщик покупного тепла на котельную (ЦТП)
			*/
		}

		if period != nil {
			var SourceFact []*tech.SourceFact
			newSourceFact, err := s.GetSourceFacts(Source.Key, period.Key)
			if err != nil {
				fmt.Println("Ошибка при получении фактических данных теплоисточников из GetSource")
			}
			SourceFact = append(SourceFact, newSourceFact)

			Source.Facts = SourceFact
		}

		//TODO: params []SourceParam
		return &Source, nil
	}
	return nil, err
}

//GetSourceQuantityFiltered возвращает КОЛИЧЕСТВО источников с учётом переданных фильтров
func (s *SQLServer) GetSourceQuantityFiltered(u auth.User, name string, address string, seasonMode int, fuelType int, period *ref.CalcPeriod) (int, error) {
	var query string
	query = fmt.Sprintf("EXEC %s.dbo.GetQuantityFilteredSources '%s', '%s', %d, %d, NULL, NULL, NULL, NULL, 0", s.dbname, name, address, seasonMode, fuelType)
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

//GetSourceFacts возвращает фактические данные по теплоисточнику за конкретный период
func (s SQLServer) GetSourceFacts(sourceID int, calcPeriodID int) (*tech.SourceFact, error) {
	var query string
	query = fmt.Sprintf("SELECT  [ID], [F_Calc_Period], [N_Work_Duration], [N_Temp_Water], [N_Temp_Air], [N_Heat_Duration], [N_Temp_Heat]"+
		", [N_Fuel_Consumption], [N_Electricity_Consumption], [N_TechWater_Constumption], [N_HotWater_Consumption], [N_Canalisation], [N_PaidHeat] "+
		"FROM %s.dbo.Source_Facts WHERE F_Source = %d AND F_Calc_Period = %d",
		s.dbname, sourceID, calcPeriodID)
	rows, err := s.db.Query(query)
	if err != nil {
		fmt.Println("Ошибка c запросом в GetSourceFacts: ", query, err)
		return nil, err
	}
	Period, err := s.GetCalcPeriod(calcPeriodID)

	defer rows.Close()
	for rows.Next() {
		var (
			ID                     int
			calcPeriodID           int     // Иденитфикатор расчётного периода
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
		)

		rows.Scan(
			&ID,
			&calcPeriodID,
			&WorkDuration,
			&TempWater,
			&TempAir,
			&HeatDuration,
			&TempHeat,
			&FuelConsumption,
			&ElectricityConsumption,
			&TechWaterConstumption,
			&HotWaterConsumption,
			&Canalisation,
			&PaidHeat)

		if err != nil {
			fmt.Println("Ошибка c запросом в GetSourceFacts при получении периода: ", query, err)
			return nil, err
		}

		SourceFact := tech.SourceFact{
			Period:                 *Period,
			WorkDuration:           WorkDuration,
			TempWater:              TempWater,
			TempAir:                TempAir,
			HeatDuration:           HeatDuration,
			TempHeat:               TempHeat,
			FuelConsumption:        FuelConsumption,
			ElectricityConsumption: ElectricityConsumption,
			TechWaterConstumption:  TechWaterConstumption,
			HotWaterConsumption:    HotWaterConsumption,
			Canalisation:           Canalisation,
			PaidHeat:               PaidHeat,
		}
		return &SourceFact, nil
	}
	SourceFact := tech.SourceFact{
		Period:                 *Period,
		WorkDuration:           744,
		TempWater:              10,
		TempAir:                25,
		HeatDuration:           744,
		TempHeat:               10,
		FuelConsumption:        0,
		ElectricityConsumption: 0,
		TechWaterConstumption:  0,
		HotWaterConsumption:    0,
		Canalisation:           0,
		PaidHeat:               0,
	}
	return &SourceFact, nil
}
