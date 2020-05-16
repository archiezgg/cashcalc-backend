/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/IstvanN/cashcalc-backend/models"
	"github.com/IstvanN/cashcalc-backend/properties"
	"github.com/IstvanN/cashcalc-backend/repositories"
	"golang.org/x/crypto/bcrypt"

	"github.com/IstvanN/cashcalc-backend/security"

	"github.com/gorilla/mux"
)

func registerLoginRoutes(router *mux.Router) {
	router.HandleFunc(properties.LoginEndpoint, loginHandler).Methods(http.MethodPost)
	router.HandleFunc(properties.RefreshEndpoint, refreshHandler).Methods(http.MethodPost)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var userToAuth models.User
	if err := json.NewDecoder(r.Body).Decode(&userToAuth); err != nil || userToAuth.Password == "" {
		security.LogErrorAndSendHTTPError(w, err, http.StatusUnprocessableEntity)
		return
	}

	u, err := repositories.GetUserByRole(userToAuth.Role)
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(userToAuth.Password))
	if err != nil {
		err := fmt.Errorf("the given role-password combination is invalid: %v - %v", userToAuth.Role, userToAuth.Password)
		security.LogErrorAndSendHTTPError(w, err, http.StatusUnauthorized)
		return
	}
	if err := generateTokenPairsAndSetThemAsHeaders(w, u.Role); err != nil {
		return
	}
	w.Write([]byte("{\"message\": \"Logged in succesfully\"}"))
	log.Printf("a user with the role '%v' has successfully logged in", userToAuth.Role)
}

func refreshHandler(w http.ResponseWriter, r *http.Request) {
	type requestedBody struct {
		RefreshToken string `json:"refreshToken"`
	}

	var rb requestedBody
	if err := json.NewDecoder(r.Body).Decode(&rb); err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusUnprocessableEntity)
		return
	}

	role, err := repositories.GetRoleFromRefreshToken(rb.RefreshToken)
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusUnauthorized)
		return
	}

	if err := generateTokenPairsAndSetThemAsHeaders(w, role); err != nil {
		return
	}

	if err := repositories.DeleteRefreshToken(rb.RefreshToken); err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	w.Write([]byte("{\"message\": \"Token refreshed successfully\"}"))
}

func generateTokenPairsAndSetThemAsHeaders(w http.ResponseWriter, role models.Role) error {
	at, err := security.CreateAccessToken(role)
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return err
	}

	rt, err := security.CreateRefreshToken(role)
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Access-Token", at)
	w.Header().Set("Refresh-Token", rt)
	return nil
}
