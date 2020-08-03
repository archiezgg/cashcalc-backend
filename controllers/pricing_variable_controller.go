/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/IstvanN/cashcalc-backend/models"

	"github.com/IstvanN/cashcalc-backend/security"

	"github.com/IstvanN/cashcalc-backend/properties"
	"github.com/IstvanN/cashcalc-backend/repositories"

	"github.com/gorilla/mux"
)

func registerPricingVarsRoutes(router *mux.Router) {
	ep := properties.PricingVarsEndpoint
	s := router.PathPrefix(ep).Subrouter()
	s.HandleFunc("", allPricingVarsHandler).Methods(http.MethodGet, http.MethodOptions)
	s.HandleFunc("/update", updatePricingVarsHandler).Methods(http.MethodPatch, http.MethodOptions)
	s.Use(security.AccessLevelAdmin)
}

func allPricingVarsHandler(w http.ResponseWriter, r *http.Request) {
	pv, err := repositories.GetPricingVariables()
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(pv)
}

func updatePricingVarsHandler(w http.ResponseWriter, r *http.Request) {
	var updatedPricingVars models.PricingVariables
	if err := json.NewDecoder(r.Body).Decode(&updatedPricingVars); err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusUnprocessableEntity)
		return
	}

	if err := repositories.UpdatePricingVariables(updatedPricingVars); err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	writeMessage(w, "Pricing variables updated successfully")
}
