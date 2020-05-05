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
	s.HandleFunc("", allPricingVarsHandler)
}

func allPricingVarsHandler(w http.ResponseWriter, r *http.Request) {
	pv, err := repositories.GetPricingVariables()
	if err != nil {
		logErrorAndSendHTTPError(w, err, 500)
		return
	}
	json.NewEncoder(w).Encode(pv)
}
