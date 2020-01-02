package controller

import (
	"encoding/json"
	"github.com/IstvanN/cashcalc-backend/model"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func registerCountriesRoutes(router *httprouter.Router) {
	router.GET("/countries", allCountriesHandler)
}

func allCountriesHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	countries := model.GetAirCountriesFromDB()
	setContentTypeToJSON(w)
	json.NewEncoder(w).Encode(countries)
}
