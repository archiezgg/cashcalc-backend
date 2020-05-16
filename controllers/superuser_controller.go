/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/IstvanN/cashcalc-backend/repositories"
	"github.com/IstvanN/cashcalc-backend/security"

	"github.com/IstvanN/cashcalc-backend/properties"
	"github.com/gorilla/mux"
)

func registerSuperuserRoutes(router *mux.Router) {
	ep := properties.SuperuserEndpoint
	s := router.PathPrefix(ep).Subrouter()
	s.HandleFunc("/tokens", tokensHandler).Methods(http.MethodGet)
	s.Use(security.AccessLevelSuperuser)
}

func tokensHandler(w http.ResponseWriter, r *http.Request) {
	tokens, err := repositories.GetAllTokens()
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tokens)
}
