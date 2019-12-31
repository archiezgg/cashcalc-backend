package controller

import (
	"net/http"
	"encoding/json"
	"github.com/IstvanN/cashcalc-backend/model"
	"github.com/julienschmidt/httprouter"
)

func registerCountriesRoutes(router *httprouter.Router) {
	router.GET("/countries", allCountriesHandler)
}

func allCountriesHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	countries := model.GetAirCountriesFromJSON()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countries)
}