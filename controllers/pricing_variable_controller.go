package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/IstvanN/cashcalc-backend/properties"
	"github.com/IstvanN/cashcalc-backend/repositories"

	"github.com/gorilla/mux"
)

func registerPricingVarsRoutes(router *mux.Router) {
	ep := properties.PricingVarsEndpoint
	s := router.PathPrefix(ep).Subrouter()
	s.HandleFunc("", allPricingVarsHandler).Methods(http.MethodGet)
}

func allPricingVarsHandler(w http.ResponseWriter, r *http.Request) {
	pv, err := repositories.GetPricingVariables()
	if err != nil {
		LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(pv)
}
