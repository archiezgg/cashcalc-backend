package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/IstvanN/cashcalc-backend/security"

	"github.com/IstvanN/cashcalc-backend/properties"
	"github.com/IstvanN/cashcalc-backend/repositories"
	"github.com/gorilla/mux"
)

func registerCountriesRoutes(router *mux.Router) {
	ep := properties.CountriesEndpoint
	s := router.PathPrefix(ep).Subrouter()
	s.HandleFunc("", allCountriesHandler).Methods(http.MethodGet)
	s.HandleFunc("/air", airCountriesHandler).Methods(http.MethodGet)
	s.HandleFunc("/road", roadCountriesHandler).Methods(http.MethodGet)
	s.Use(security.AuthCarrierLevel)
}

func allCountriesHandler(w http.ResponseWriter, r *http.Request) {
	c, err := repositories.GetCountries()
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(c)
}

func airCountriesHandler(w http.ResponseWriter, r *http.Request) {
	ac, err := repositories.GetAirCountries()
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(ac)
}

func roadCountriesHandler(w http.ResponseWriter, r *http.Request) {
	rc, err := repositories.GetRoadCountries()
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(rc)
}
