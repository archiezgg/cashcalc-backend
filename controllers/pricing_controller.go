package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/IstvanN/cashcalc-backend/repositories"
	"github.com/gorilla/mux"
)

var (
	pricingsEndpoint = os.Getenv("PRICINGS_ENDPOINT")
	faresEndpoint    = os.Getenv("FARES_ENDPOINT")
	docFaresEndpoint = os.Getenv("DOCFARES_ENDPOINT")
)

func registerPricingsRoutes(router *mux.Router) {
	router.HandleFunc(pricingsEndpoint, allPricingsHandler).
		Methods("GET").
		Queries("type", "{type:[a-zA-Z]+}")
	router.HandleFunc(faresEndpoint+"/{zn:[0-9]}", pricingFaresByZoneNumberHandler).
		Methods("GET").
		Queries("type", "{type:[a-zA-Z]+}")
	router.HandleFunc(docFaresEndpoint+"/{zn:[5-9]}", pricingDocFaresByZoneNumberHandler).
		Methods("GET")
}

func allPricingsHandler(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJSON(w)
	switch t := mux.Vars(r)["type"]; t {
	case "air":
		airPricings, err := repositories.GetAirPricingsFromDB()
		if err != nil {
			log.Println(err)
		}
		json.NewEncoder(w).Encode(airPricings)
	case "road":
		roadPricings, err := repositories.GetRoadPricingsFromDB()
		if err != nil {
			log.Println(err)
		}
		json.NewEncoder(w).Encode(roadPricings)
	default:
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}
}

func pricingFaresByZoneNumberHandler(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJSON(w)

	zn, _ := strconv.Atoi(mux.Vars(r)["zn"])
	switch t := mux.Vars(r)["type"]; t {
	case "air":
		airFares, err := repositories.GetAirPricingFaresByZoneNumber(zn)
		if err != nil {
			log.Println(err)
		}
		json.NewEncoder(w).Encode(airFares)
	case "road":
		roadFares, err := repositories.GetRoadPricingFaresByZoneNumber(zn)
		if err != nil {
			log.Println(err)
		}
		json.NewEncoder(w).Encode(roadFares)
	default:
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}
}

func pricingDocFaresByZoneNumberHandler(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJSON(w)

	zn, _ := strconv.Atoi(mux.Vars(r)["zn"])
	airDocFares, err := repositories.GetAirPricingDocFaresByZoneNumber(zn)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(airDocFares)
}
