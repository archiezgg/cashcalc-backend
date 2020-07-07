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
	if err := json.NewDecoder(r.Body).Decode(&userToAuth); err != nil ||
		userToAuth.Password == "" || userToAuth.Username == "" {
		security.LogErrorAndSendHTTPError(w, err, http.StatusUnprocessableEntity)
		return
	}

	u, err := repositories.GetUserByUsername(userToAuth.Username)
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(userToAuth.Password))
	if err != nil {
		err := fmt.Errorf("the given role-password combination is invalid: %v - %v", userToAuth.Username, userToAuth.Password)
		security.LogErrorAndSendHTTPError(w, err, http.StatusUnauthorized)
		return
	}
	if err := security.GenerateTokenPairsAndSetThemAsCookies(w, u); err != nil {
		return
	}
	log.Printf("user '%v' has successfully logged in", u.Username)
	msg := fmt.Sprintf("{\"message\": \"%s\",\"role\": \"%v\"}", "Logged in successfully", u.Role)
	w.Write([]byte(msg))
}

func refreshHandler(w http.ResponseWriter, r *http.Request) {
	type requestedBody struct {
		RefreshToken string `json:"refreshToken"`
	}

	var rb requestedBody
	if err := json.NewDecoder(r.Body).Decode(&rb); err != nil || rb.RefreshToken == "" {
		security.LogErrorAndSendHTTPError(w, err, http.StatusUnprocessableEntity)
		return
	}

	if err := security.RefreshToken(w, rb.RefreshToken); err != nil {
		return
	}

	writeMessage(w, "Token refreshed successfully")
}
