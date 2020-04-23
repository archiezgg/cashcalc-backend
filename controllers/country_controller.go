package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/IstvanN/cashcalc-backend/properties"
	"github.com/IstvanN/cashcalc-backend/repositories"
	"github.com/gorilla/mux"
)

func registerCountriesRoutes(router *mux.Router) {
	ep := properties.CountriesEndpoint
	router.HandleFunc(ep, allCountriesHandler).Methods(http.MethodGet)
	router.HandleFunc(ep+"/air", airCountriesHandler).Methods(http.MethodGet)
	router.HandleFunc(ep+"/road", roadCountriesHandler).Methods(http.MethodGet)
}

func allCountriesHandler(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJSON(w)
	c, err := repositories.GetCountries()
	if err != nil {
		logErrorAndSendHTTPError(w, err, 500)
		return
	}
	json.NewEncoder(w).Encode(c)
}

func airCountriesHandler(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJSON(w)
	ac, err := repositories.GetAirCountries()
	if err != nil {
		logErrorAndSendHTTPError(w, err, 500)
		return
	}
	json.NewEncoder(w).Encode(ac)
}

func roadCountriesHandler(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJSON(w)
	rc, err := repositories.GetRoadCountries()
	if err != nil {
		logErrorAndSendHTTPError(w, err, 500)
		return
	}
	json.NewEncoder(w).Encode(rc)
}
