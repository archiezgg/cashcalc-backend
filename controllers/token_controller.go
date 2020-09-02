/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/IstvanN/cashcalc-backend/repositories"
	"github.com/IstvanN/cashcalc-backend/security"

	"github.com/IstvanN/cashcalc-backend/properties"
	"github.com/gorilla/mux"
)

func registerTokenRoutes(router *mux.Router) {
	ep := properties.TokensEndpoint
	s := router.PathPrefix(ep).Subrouter()
	s.HandleFunc("", tokensHandler).Methods(http.MethodGet, http.MethodOptions)
	s.HandleFunc("/logged-in-users", loggedInUsersHandler).Methods(http.MethodGet, http.MethodOptions)
	s.HandleFunc("/revoke", revokeTokensHandler).Methods(http.MethodDelete, http.MethodOptions)
	s.Use(security.AccessLevelSuperuser)
}

func tokensHandler(w http.ResponseWriter, r *http.Request) {
	tokens, err := repositories.GetAllRefreshTokens()
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tokens)
}

func loggedInUsersHandler(w http.ResponseWriter, r *http.Request) {
	loggedInUsers, err := repositories.GetAllLoggedInUsers()
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	loggedInUserDTOs := repositories.CreateUserDTOsFromUsers(loggedInUsers)
	json.NewEncoder(w).Encode(loggedInUserDTOs)
}

func revokeTokensHandler(w http.ResponseWriter, r *http.Request) {
	idAsString := r.URL.Query().Get("id")
	if idAsString == "" {
		err := fmt.Errorf("user id parameter is not defined")
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	id, err := strconv.Atoi(idAsString)
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
	}

	if err := repositories.DeleteAllRefreshTokensForUser(uint(id)); err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	w.Write([]byte("{\"message\": \"Token revoked successfully\"}"))
}
