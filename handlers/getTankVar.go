package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/jamesthomasw/assignment_3/helpers"
)

type TankVar struct {
	Water int64 `json:"water"`
	StatusWater string `json:"status_water"`
	Wind int64	`json:"wind"`
	StatusWind string `json:"status_wind"`
}

func CheckStatusWater(tv TankVar) string {
	if tv.Water <= 5 {
		result := "Aman"
		return result
	} else if tv.Water > 5 && tv.Water < 9 {
		result := "Siaga"
		return result
	} else {
		result := "Bahaya"
		return result
	}
}

func CheckStatusWind(tv TankVar) string {
	if tv.Wind <= 6 {
		result := "Aman"
		return result
	} else if tv.Wind > 6 && tv.Wind < 16 {
		result := "Siaga"
		return result
	} else {
		result := "Bahaya"
		return result
	}
}

func (tv *TankVar) SetWater(result string) {
	tv.StatusWater = result
}

func (tv *TankVar) SetWind(result string) {
	tv.StatusWind = result
}

func GetTankVar(w http.ResponseWriter, r *http.Request) {

	newTankVar := TankVar{
		Water: helpers.GenRandNum(1, 100),
		Wind:  helpers.GenRandNum(1, 100),
	}

	statWater := CheckStatusWater(newTankVar)
	newTankVar.SetWater(statWater)

	statWind := CheckStatusWind(newTankVar)
	newTankVar.SetWind(statWind)


	t, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, newTankVar)

	u, err := json.Marshal(newTankVar)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(u))
}