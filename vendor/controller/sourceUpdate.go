package controller

import (
	"core/ref"
	"fmt"
	"net/http"
	"strconv"
)

func (h *DecoratedHandler) sourceUpdate(w http.ResponseWriter, r *http.Request) { //

	// Получаем текущий теплоисточник из параметров
	keySource, err := strconv.Atoi(r.URL.Query().Get("key"))

	// Получаем текущий период из параметров
	var calcPeriod *ref.CalcPeriod
	period := r.URL.Query().Get("period")
	if period == "" {
		calcPeriod, err = h.connection.GetCurrentPeriod()
	} else {
		calcPeriodID, err := strconv.Atoi(period)
		calcPeriod, err = h.connection.GetCalcPeriod(calcPeriodID)
		if err != nil {
			fmt.Println("Передано ошибочное значение расчётного периода")
		}
	}

	Source, err := h.connection.GetSource(keySource, calcPeriod)

	session, err := Store.Get(r, "cookie-name")
	check(err)
	user := GetUser(session)

	if r.Method == http.MethodGet {
		check(err)
		currentInformation := sessionInformation{User: *user, Attribute: Source}
		executeHTML("source", "update", w, currentInformation)
	}

	if r.Method == http.MethodPost {
		/*
			name := r.FormValue("name")
			familyName := r.FormValue("familyname")
			patronymicName := r.FormValue("patronymicname")
			dateBirth := r.FormValue("datebirth")
			dateDeath := r.FormValue("datedeath")
			sex, _ := strconv.ParseBool(r.FormValue("sex"))
			userLogin := r.FormValue("user")

			dateBirthG, _ := time.Parse("2006-01-02", dateBirth)
			dateDeathG, _ := time.Parse("2006-01-02", dateDeath)

			newUser, err := h.connection.GetUser(userLogin)
			newPerson := contract.Person{
				Key:            Person.Key,
				Name:           name,
				FamilyName:     familyName,
				PatronymicName: patronymicName,
				Sex:            sex,
				DateBirth:      dateBirthG,
				DateDeath:      dateDeathG,
				User:           *newUser,
			}

			err = h.connection.UpdatePerson(&newPerson)

			if err != nil {
				executeHTML("person", "update", w, nil)
			}
			http.Redirect(w, r, "/person", http.StatusFound)
		*/
	}

}
