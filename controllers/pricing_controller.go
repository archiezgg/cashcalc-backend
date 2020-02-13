package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/IstvanN/cashcalc-backend/model"
	"github.com/IstvanN/cashcalc-backend/service"
	"github.com/gorilla/mux"
)

func registerPricingsRoutes(router *mux.Router) {
	router.HandleFunc("/pricings", allPricingsHandler).
		Methods("GET").
		Queries("type", "{type:[a-zA-Z]+}")
	router.HandleFunc("/pricings/fares/{zn:[0-9]}", pricingFaresByZoneNumberHandler).
		Methods("GET").
		Queries("type", "{type:[a-zA-Z]+}")
	router.HandleFunc("/pricings/docfares/{zn:[5-9]}", pricingDocFaresByZoneNumberHandler).
		Methods("GET")
}

func allPricingsHandler(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJSON(w)
	switch t := mux.Vars(r)["type"]; t {
	case "air":
		airPricings, err := model.GetAirPricingsFromDB()
		if err != nil {
			log.Println(err)
		}
		json.NewEncoder(w).Encode(airPricings)
	case "road":
		roadPricings, err := model.GetRoadPricingsFromDB()
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
		airFares, err := service.GetAirPricingFaresByZoneNumber(zn)
		if err != nil {
			log.Println(err)
		}
		json.NewEncoder(w).Encode(airFares)
	case "road":
		roadFares, err := service.GetRoadPricingFaresByZoneNumber(zn)
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
	airDocFares, err := service.GetAirPricingDocFaresByZoneNumber(zn)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(airDocFares)
}
