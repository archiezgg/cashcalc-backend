package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/IstvanN/cashcalc-backend/model"
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
		airCountries, err := model.GetAirCountriesFromDB()
		if err != nil {
			log.Println(err)
		}
		json.NewEncoder(w).Encode(airCountries)
	case "road":
		roadCountries, err := model.GetRoadCountriesFromDB()
		if err != nil {
			log.Println(err)
		}
		json.NewEncoder(w).Encode(roadCountries)
	default:
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}
}
