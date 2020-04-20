package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/IstvanN/cashcalc-backend/services"
	"github.com/gorilla/mux"
)

var countriesEndpoint = os.Getenv("COUNTRIES_ENDPOINT")

func registerCountriesRoutes(router *mux.Router) {
	router.HandleFunc(countriesEndpoint, allCountriesHandler).Methods(http.MethodGet)
	router.HandleFunc(countriesEndpoint+"/air", airCountriesHandler).Methods(http.MethodGet)
	router.HandleFunc(countriesEndpoint+"/road", roadCountriesHandler).Methods(http.MethodGet)
}

func allCountriesHandler(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJSON(w)
	c, err := services.GetCountries()
	if err != nil {
		logErrorAndSendHTTPError(w, err, 500)
		return
	}
	json.NewEncoder(w).Encode(c)
}

func airCountriesHandler(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJSON(w)
	ac, err := services.GetAirCountries()
	if err != nil {
		logErrorAndSendHTTPError(w, err, 500)
		return
	}
	json.NewEncoder(w).Encode(ac)
}

func roadCountriesHandler(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJSON(w)
	rc, err := services.GetRoadCountries()
	if err != nil {
		logErrorAndSendHTTPError(w, err, 500)
		return
	}
	json.NewEncoder(w).Encode(rc)
}
