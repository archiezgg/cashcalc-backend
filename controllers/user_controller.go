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

	"github.com/gorilla/mux"
)

func registerUserRoutes(router *mux.Router) {
	s := router.PathPrefix("/users").Subrouter()
	s.HandleFunc("", usernamesHandler)
	s.Use(security.AccessLevelAdmin)
	// s.HandleFunc("/create/carrier", createCarrierHandler)
}

func usernamesHandler(w http.ResponseWriter, r *http.Request) {
	usernames, err := repositories.GetUsernames()
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(usernames)
}

// func createCarrierHandler(w http.ResponseWriter, r *http.Request) {
// 	type requestedBody struct {
// 		Username string `json:"username"`
// 		Password string `json:"password"`
// 	}

// 	var rb requestedBody
// 	if err := decodeIntoRequestedBody(w, r, requestedBody)
// }
