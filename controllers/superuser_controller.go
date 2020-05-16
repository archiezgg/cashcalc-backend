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
	s.HandleFunc("/tokens/delete", deleteTokenHandler).Methods(http.MethodDelete)
	s.HandleFunc("/tokens/deleteBulk", deleteBulkTokenHandler).Methods(http.MethodDelete)
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

func deleteTokenHandler(w http.ResponseWriter, r *http.Request) {
	type requestedBody struct {
		RefreshToken string `json:"refreshToken"`
	}

	var rb requestedBody
	if err := json.NewDecoder(r.Body).Decode(&rb); err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusUnprocessableEntity)
		return
	}

	if err := repositories.DeleteRefreshToken(rb.RefreshToken); err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	w.Write([]byte("{\"message\": \"Token deleted successfully\"}"))
}

func deleteBulkTokenHandler(w http.ResponseWriter, r *http.Request) {
	type requestedBody struct {
		RefreshTokens []string `json:"refreshTokens"`
	}

	var rb requestedBody
	if err := json.NewDecoder(r.Body).Decode(&rb); err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusUnprocessableEntity)
		return
	}

	if err := repositories.DeleteBulkRefreshToken(rb.RefreshTokens); err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	w.Write([]byte("{\"message\": \"Multiple tokens deleted successfully\"}"))
}
