package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/IstvanN/cashcalc-backend/model"
	"github.com/gorilla/mux"
)

func registerPricingsRoutes(router *mux.Router) {
	router.HandleFunc("/pricings", allPricingsHandler).Methods("GET").Queries("type", "{type:[a-zA-Z]+}")
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
