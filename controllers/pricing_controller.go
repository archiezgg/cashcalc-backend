package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/IstvanN/cashcalc-backend/properties"
	"github.com/IstvanN/cashcalc-backend/repositories"
	"github.com/gorilla/mux"
)

func registerPricingsRoutes(router *mux.Router) {
	ep := properties.PricingsEndpoint
	s := router.PathPrefix(ep).Subrouter()
	s.HandleFunc("", allPricingsHandler).Methods(http.MethodGet)
	s.HandleFunc("/road", roadPricingsHandler).Methods(http.MethodGet)
	s.HandleFunc("/air", airPricingsHandler).Methods(http.MethodGet)
	s.HandleFunc("/road/fares/{zn:[1-5]}", roadFaresByZoneNumberHandler).
		Methods(http.MethodGet)
	s.HandleFunc("/air/fares/{zn:[0-9]}", airFaresByZoneNumberHandler).
		Methods(http.MethodGet)
	s.HandleFunc("/air/docfares/{zn:[5-9]}", airDocFaresByZoneNumberHandler).
		Methods(http.MethodGet)
}

func allPricingsHandler(w http.ResponseWriter, r *http.Request) {
	p, err := repositories.GetPricings()
	if err != nil {
		logErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(p)
}

func roadPricingsHandler(w http.ResponseWriter, r *http.Request) {
	rp, err := repositories.GetRoadPricings()
	if err != nil {
		logErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(rp)
}

func airPricingsHandler(w http.ResponseWriter, r *http.Request) {
	ap, err := repositories.GetAirPricings()
	if err != nil {
		logErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(ap)
}

func roadFaresByZoneNumberHandler(w http.ResponseWriter, r *http.Request) {
	zn, _ := strconv.Atoi(mux.Vars(r)["zn"])
	weightAsString, queryIsPresent := r.URL.Query()["weight"]

	if queryIsPresent {
		weight, err := strconv.ParseFloat(weightAsString[0], 64)
		if err != nil {
			logErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
			return
		}

		rp, err := repositories.GetRoadFaresByZoneNumberAndWeight(zn, weight)
		if err != nil {
			logErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(rp)
		return
	}

	rp, err := repositories.GetRoadFaresByZoneNumber(zn)
	if err != nil {
		logErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(rp)
}

func airFaresByZoneNumberHandler(w http.ResponseWriter, r *http.Request) {
	zn, _ := strconv.Atoi(mux.Vars(r)["zn"])
	weightAsString, queryIsPresent := r.URL.Query()["weight"]

	if queryIsPresent {
		weight, err := strconv.ParseFloat(weightAsString[0], 64)
		if err != nil {
			logErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
			return
		}

		ap, err := repositories.GetAirFaresByZoneNumberAndWeight(zn, weight)
		if err != nil {
			logErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(ap)
		return
	}

	ap, err := repositories.GetAirFaresByZoneNumber(zn)
	if err != nil {
		logErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(ap)
}

func airDocFaresByZoneNumberHandler(w http.ResponseWriter, r *http.Request) {
	zn, _ := strconv.Atoi(mux.Vars(r)["zn"])
	weightAsString, queryIsPresent := r.URL.Query()["weight"]

	if queryIsPresent {
		weight, err := strconv.ParseFloat(weightAsString[0], 64)
		if err != nil {
			logErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
			return
		}

		ap, err := repositories.GetAirDocFaresByZoneNumberAndWeight(zn, weight)
		if err != nil {
			logErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(ap)
		return
	}

	ap, err := repositories.GetAirDocFaresByZoneNumber(zn)
	if err != nil {
		logErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(ap)

}
