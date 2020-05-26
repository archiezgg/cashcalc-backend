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

	"github.com/IstvanN/cashcalc-backend/repositories"
	"github.com/IstvanN/cashcalc-backend/security"

	"github.com/gorilla/mux"
)

func registerUserRoutes(router *mux.Router) {
	s := router.PathPrefix("/users").Subrouter()
	s.HandleFunc("", usernamesHandler).Methods(http.MethodGet)
	s.Use(security.AccessLevelAdmin)
	carriers := s.PathPrefix("/carrier").Subrouter()
	carriers.HandleFunc("/create", createCarrierHandler).Methods(http.MethodPut)
	carriers.HandleFunc("/delete", deleteCarrierHandler).Methods(http.MethodDelete)
	carriers.Use(security.AccessLevelAdmin)
}

func usernamesHandler(w http.ResponseWriter, r *http.Request) {
	usernames, err := repositories.GetUsernames()
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(usernames)
}

func createCarrierHandler(w http.ResponseWriter, r *http.Request) {
	type requestedBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var rb requestedBody
	if err := json.NewDecoder(r.Body).Decode(&rb); err != nil ||
		rb.Username == "" || rb.Password == "" {
		security.LogErrorAndSendHTTPError(w, err, http.StatusUnprocessableEntity)
		return
	}

	if err := repositories.CreateUser(rb.Username, rb.Password, models.RoleCarrier); err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	writeMessage(w, "Carrier created successfully")
}

func deleteCarrierHandler(w http.ResponseWriter, r *http.Request) {
	type requestedBody struct {
		Username string `json:"username"`
	}

	var rb requestedBody
	if err := json.NewDecoder(r.Body).Decode(&rb); err != nil || rb.Username == "" {
		security.LogErrorAndSendHTTPError(w, err, http.StatusUnprocessableEntity)
		return
	}

	if err := repositories.DeleteUserByUsernameAndRole(rb.Username, models.RoleCarrier); err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	writeMessage(w, "Carrier deleted successfully")
}
