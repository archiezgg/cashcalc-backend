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

	"github.com/IstvanN/cashcalc-backend/models"
	"github.com/IstvanN/cashcalc-backend/properties"

	"github.com/IstvanN/cashcalc-backend/repositories"
	"github.com/IstvanN/cashcalc-backend/security"

	"github.com/gorilla/mux"
)

func registerUserRoutes(router *mux.Router) {
	ep := properties.UsersEndpoint
	s := router.PathPrefix(ep).Subrouter()
	s.HandleFunc("/usernames", usernamesHandler).Methods(http.MethodGet, http.MethodOptions)
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
	superusers := s.PathPrefix("/superusers").Subrouter()
	superusers.HandleFunc("", getSuperusersHandler).Methods(http.MethodGet, http.MethodOptions)
	superusers.HandleFunc("/create", createSuperuserHandler).Methods(http.MethodPut, http.MethodOptions)
	superusers.Use(security.AccessLevelSuperuser)
}

func usernamesHandler(w http.ResponseWriter, r *http.Request) {
	usernames, err := repositories.GetAllUsernames()
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(usernames)
}

func getCarriersHandler(w http.ResponseWriter, r *http.Request) {
	carriers, err := repositories.GetUsersByRole(models.RoleCarrier)
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	carrierDTOs := repositories.CreateUserDTOsFromUsers(carriers)

	json.NewEncoder(w).Encode(carrierDTOs)
}

func getAdminsHandler(w http.ResponseWriter, r *http.Request) {
	admins, err := repositories.GetUsersByRole(models.RoleAdmin)
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	adminDTOs := repositories.CreateUserDTOsFromUsers(admins)

	json.NewEncoder(w).Encode(adminDTOs)
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

	if err := repositories.DeleteUserByIDAndRole(uint(id), models.RoleCarrier); err != nil {
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

	if err := repositories.DeleteUserByIDAndRole(uint(id), models.RoleAdmin); err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	writeMessage(w, "Admin deleted successfully")
}

func getSuperusersHandler(w http.ResponseWriter, r *http.Request) {
	superusers, err := repositories.GetUsersByRole(models.RoleSuperuser)
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	superuserDTOs := repositories.CreateUserDTOsFromUsers(superusers)

	json.NewEncoder(w).Encode(superuserDTOs)
}

func createSuperuserHandler(w http.ResponseWriter, r *http.Request) {
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

	if err := repositories.CreateUser(rb.Username, rb.Password, models.RoleSuperuser); err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	writeMessage(w, "Superuser created successfully")
}
