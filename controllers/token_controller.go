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

func registerTokenRoutes(router *mux.Router) {
	ep := properties.TokensEndpoint
	s := router.PathPrefix(ep).Subrouter()
	s.HandleFunc("", tokensHandler).Methods(http.MethodGet, http.MethodOptions)
	s.HandleFunc("/revoke", revokeTokenHandler).Methods(http.MethodDelete, http.MethodOptions)
	s.HandleFunc("/revoke-bulk", revokeBulkTokenHandler).Methods(http.MethodDelete, http.MethodOptions)
	s.HandleFunc("/revoke-all", revokeAllTokensHandler).Methods(http.MethodDelete, http.MethodOptions)
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

func revokeTokenHandler(w http.ResponseWriter, r *http.Request) {
	type requestedBody struct {
		Username string `json:"username"`
	}

	var rb requestedBody
	if err := json.NewDecoder(r.Body).Decode(&rb); err != nil || rb.Username == "" {
		security.LogErrorAndSendHTTPError(w, err, http.StatusUnprocessableEntity)
		return
	}

	if err := repositories.DeleteRefreshToken(rb.Username); err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	w.Write([]byte("{\"message\": \"Token revoked successfully\"}"))
}

func revokeBulkTokenHandler(w http.ResponseWriter, r *http.Request) {
	type requestedBody struct {
		Usernames []string `json:"usernames"`
	}

	var rb requestedBody
	if err := json.NewDecoder(r.Body).Decode(&rb); err != nil || rb.Usernames == nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusUnprocessableEntity)
		return
	}

	if err := repositories.DeleteBulkRefreshToken(rb.Usernames); err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	w.Write([]byte("{\"message\": \"Multiple tokens revoked successfully\"}"))
}

func revokeAllTokensHandler(w http.ResponseWriter, r *http.Request) {
	if err := repositories.DeleteAllTokens(); err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	w.Write([]byte("{\"message\": \"All tokens revoked successfully\"}"))
}
