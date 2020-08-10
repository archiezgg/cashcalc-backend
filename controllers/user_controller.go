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
	"github.com/IstvanN/cashcalc-backend/properties"

	"github.com/IstvanN/cashcalc-backend/repositories"
	"github.com/IstvanN/cashcalc-backend/security"

	"github.com/gorilla/mux"
)

func registerUserRoutes(router *mux.Router) {
	ep := properties.UsersEndpoint
	s := router.PathPrefix(ep).Subrouter()
	s.HandleFunc("/all", usernamesHandler).Methods(http.MethodGet, http.MethodOptions)
	s.Use(security.AccessLevelAdmin)
	carriers := s.PathPrefix("/carriers").Subrouter()
	carriers.HandleFunc("", getCarriersHandler).Methods(http.MethodGet, http.MethodOptions)
	carriers.HandleFunc("/create", createCarrierHandler).Methods(http.MethodPut, http.MethodOptions)
	carriers.HandleFunc("/delete", deleteCarrierHandler).Methods(http.MethodDelete, http.MethodOptions)
	carriers.Use(security.AccessLevelAdmin)
	admins := s.PathPrefix("/admins").Subrouter()
	admins.HandleFunc("", getAdminsHandler).Methods(http.MethodGet, http.MethodOptions)
	admins.HandleFunc("/create", createAdminHandler).Methods(http.MethodPut, http.MethodOptions)
	admins.HandleFunc("/delete", deleteAdminHandler).Methods(http.MethodDelete, http.MethodOptions)
	admins.Use(security.AccessLevelSuperuser)
}

func usernamesHandler(w http.ResponseWriter, r *http.Request) {
	usernames, err := repositories.GetUsernames()
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(usernames)
}

func getCarriersHandler(w http.ResponseWriter, r *http.Request) {
	carriers, err := repositories.GetUsernamesByRole(models.RoleCarrier)
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(carriers)
}

func getAdminsHandler(w http.ResponseWriter, r *http.Request) {
	admins, err := repositories.GetUsernamesByRole(models.RoleAdmin)
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(admins)
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

func createAdminHandler(w http.ResponseWriter, r *http.Request) {
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

	if err := repositories.CreateUser(rb.Username, rb.Password, models.RoleAdmin); err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	writeMessage(w, "Admin created successfully")
}

func deleteAdminHandler(w http.ResponseWriter, r *http.Request) {
	type requestedBody struct {
		Username string `json:"username"`
	}

	var rb requestedBody
	if err := json.NewDecoder(r.Body).Decode(&rb); err != nil || rb.Username == "" {
		security.LogErrorAndSendHTTPError(w, err, http.StatusUnprocessableEntity)
		return
	}

	if err := repositories.DeleteUserByUsernameAndRole(rb.Username, models.RoleAdmin); err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	writeMessage(w, "Admin deleted successfully")
}
