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

	"github.com/IstvanN/cashcalc-backend/security"

	"github.com/gorilla/mux"
)

func registerLoginRoutes(router *mux.Router) {
	router.HandleFunc(properties.LoginEndpoint, loginHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc(properties.RefreshEndpoint, refreshHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc(properties.LogoutEndpoint, logoutHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc(properties.IsAuthorizedEndpoint, isAuthorizedHandler).Methods(http.MethodGet, http.MethodOptions)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var userToAuth models.User
	if err := json.NewDecoder(r.Body).Decode(&userToAuth); err != nil ||
		userToAuth.Password == "" || userToAuth.Username == "" {
		security.LogErrorAndSendHTTPError(w, err, http.StatusUnprocessableEntity)
		return
	}

	user, err := security.AuthenticateNewUser(w, userToAuth)
	if err != nil {
		return
	}
	msg := fmt.Sprintf("{\"message\": \"%s\",\"role\": \"%v\"}", "Logged in successfully", user.Role)
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

	if _, err := security.RefreshTokenAndSetTokensAsCookies(w, rb.RefreshToken); err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	writeMessage(w, "Token refreshed successfully")
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if err := security.DeleteTokensFromCookies(w, r); err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	writeMessage(w, "Logged out succesfully")
}

func isAuthorizedHandler(w http.ResponseWriter, r *http.Request) {
	roleToCompare, ok := r.URL.Query()["role"]
	if !ok {
		err := fmt.Errorf("role parameter not found in URL query")
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	log.Println(roleToCompare)
	if !security.IsTokenValidForAccessLevel(models.Role(roleToCompare[0]), w, r) {
		return
	}

	writeMessage(w, "authorized")
}
