package controller

import (
	"encoding/json"
	"net/http"

	"github.com/IstvanN/cashcalc-backend/model"
	"github.com/gorilla/mux"
)

func registerCountriesRoutes(router *mux.Router) {
	router.HandleFunc("/countries", allCountriesHandler).Methods("GET").Queries("type", "{type:[a-zA-Z]+}")
}

func allCountriesHandler(w http.ResponseWriter, r *http.Request) {
	switch t := mux.Vars(r)["type"]; t {
	case "air":
		setContentTypeToJSON(w)
		countries := model.GetAirCountriesFromDB()
		json.NewEncoder(w).Encode(countries)
	case "road":
		setContentTypeToJSON(w)
		countries := model.GetRoadCountriesFromDB()
		json.NewEncoder(w).Encode(countries)
	default:
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}
}
