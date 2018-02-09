package api

import (
	"net/http"
	"../controller"
	"encoding/json"
	"../models"
)

func GetAllData(w http.ResponseWriter, r *http.Request)   {
	visas := controller.GetAllData()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(visas);
	if err != nil {
		println(err.Error())
	}
}

func CreateData(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var visa models.Visa
	err := decoder.Decode(&visa)
	if err != nil{
		println(err.Error())
	}
	controller.CreateData(visa)
}
