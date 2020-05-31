/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/IstvanN/cashcalc-backend/properties"
	"github.com/IstvanN/cashcalc-backend/repositories"

	"github.com/IstvanN/cashcalc-backend/security"

	"github.com/IstvanN/cashcalc-backend/models"

	"github.com/gorilla/mux"
)

func registerCalcRoutes(router *mux.Router) {
	ep := properties.CalcEndpoint
	s := router.PathPrefix(ep).Subrouter()
	s.HandleFunc("", calcHandler).Methods(http.MethodPost)
	s.Use(security.AccessLevelCarrier)
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
	var inputData models.CalcInputData
	if err := json.NewDecoder(r.Body).Decode(&inputData); err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusUnprocessableEntity)
		return
	}

	resultData, err := repositories.CalcResult(inputData)
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resultData)
}
