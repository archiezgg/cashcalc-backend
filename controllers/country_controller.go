package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/IstvanN/cashcalc-backend/repositories"
	"github.com/gorilla/mux"
)

func registerCountriesRoutes(router *mux.Router) {
	router.HandleFunc("/countries", allCountriesHandler).
		Methods("GET").
		Queries("type", "{type:[a-zA-Z]+}")
}

func allCountriesHandler(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJSON(w)
	switch t := mux.Vars(r)["type"]; t {
	case "air":
		countriesAir, err := repositories.GetCountriesAirFromDB()
		if err != nil {
			log.Println(err)
		}
		json.NewEncoder(w).Encode(countriesAir)
	case "road":
		countriesRoad, err := repositories.GetCountriesRoadFromDB()
		if err != nil {
			log.Println(err)
		}
		json.NewEncoder(w).Encode(countriesRoad)
	default:
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}
}
