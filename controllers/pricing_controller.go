package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/IstvanN/cashcalc-backend/services"
	"github.com/gorilla/mux"
)

var pricingsEndpoint = os.Getenv("PRICINGS_ENDPOINT")

func registerPricingsRoutes(router *mux.Router) {
	router.HandleFunc(pricingsEndpoint, allPricingsHandler).Methods(http.MethodGet)
	router.HandleFunc(pricingsEndpoint+"/road", roadPricingsHandler).Methods(http.MethodGet)
	router.HandleFunc(pricingsEndpoint+"/air", airPricingsHandler).Methods(http.MethodGet)
	router.HandleFunc(pricingsEndpoint+"/road/fares/{zn:[1-5]}", roadPricingFaresByZoneNumberHandler).Methods(http.MethodGet)
	router.HandleFunc(pricingsEndpoint+"/air/fares/{zn:[0-9]}", airPricingFaresByZoneNumberHandler).Methods(http.MethodGet)
	router.HandleFunc(pricingsEndpoint+"/air/docfares/{zn:[5-9]}", airPricingDocFaresByZoneNumberHandler).Methods(http.MethodGet)
}

func allPricingsHandler(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJSON(w)
	p, err := services.GetPricings()
	if err != nil {
		logErrorAndSendHTTPError(w, err, 500)
		return
	}
	json.NewEncoder(w).Encode(p)
}

func roadPricingsHandler(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJSON(w)
	rp, err := services.GetRoadPricings()
	if err != nil {
		logErrorAndSendHTTPError(w, err, 500)
		return
	}
	json.NewEncoder(w).Encode(rp)
}

func airPricingsHandler(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJSON(w)
	ap, err := services.GetAirPricings()
	if err != nil {
		logErrorAndSendHTTPError(w, err, 500)
		return
	}
	json.NewEncoder(w).Encode(ap)
}

func roadPricingFaresByZoneNumberHandler(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJSON(w)
	zn, _ := strconv.Atoi(mux.Vars(r)["zn"])
	weightAsString, queryIsPresent := r.URL.Query()["weight"]

	if queryIsPresent {
		weight, err := strconv.ParseFloat(weightAsString[0], 64)
		if err != nil {
			logErrorAndSendHTTPError(w, err, 500)
			return
		}

		rp, err := services.GetRoadFaresByZoneNumberAndWeight(zn, weight)
		if err != nil {
			logErrorAndSendHTTPError(w, err, 500)
			return
		}
		json.NewEncoder(w).Encode(rp)
		return
	}

	rp, err := services.GetRoadFaresByZoneNumber(zn)
	if err != nil {
		logErrorAndSendHTTPError(w, err, 500)
		return
	}
	json.NewEncoder(w).Encode(rp)
}

func airPricingFaresByZoneNumberHandler(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJSON(w)
	zn, _ := strconv.Atoi(mux.Vars(r)["zn"])
	weightAsString, queryIsPresent := r.URL.Query()["weight"]

	if queryIsPresent {
		weight, err := strconv.ParseFloat(weightAsString[0], 64)
		if err != nil {
			logErrorAndSendHTTPError(w, err, 500)
			return
		}

		ap, err := services.GetAirFaresByZoneNumberAndWeight(zn, weight)
		if err != nil {
			logErrorAndSendHTTPError(w, err, 500)
			return
		}
		json.NewEncoder(w).Encode(ap)
		return
	}

	ap, err := services.GetAirFaresByZoneNumber(zn)
	if err != nil {
		logErrorAndSendHTTPError(w, err, 500)
		return
	}
	json.NewEncoder(w).Encode(ap)
}

func airPricingDocFaresByZoneNumberHandler(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJSON(w)
	zn, _ := strconv.Atoi(mux.Vars(r)["zn"])
	weightAsString, queryIsPresent := r.URL.Query()["weight"]

	if queryIsPresent {
		weight, err := strconv.ParseFloat(weightAsString[0], 64)
		if err != nil {
			logErrorAndSendHTTPError(w, err, 500)
			return
		}

		ap, err := services.GetAirDocFaresByZoneNumberAndWeight(zn, weight)
		if err != nil {
			logErrorAndSendHTTPError(w, err, 500)
			return
		}
		json.NewEncoder(w).Encode(ap)
		return
	}

	ap, err := services.GetAirDocFaresByZoneNumber(zn)
	if err != nil {
		logErrorAndSendHTTPError(w, err, 500)
		return
	}
	json.NewEncoder(w).Encode(ap)

}
