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

	"github.com/IstvanN/cashcalc-backend/models"
	"github.com/IstvanN/cashcalc-backend/properties"

	"github.com/IstvanN/cashcalc-backend/security"

	"github.com/gorilla/mux"
)

func registerLoginRoutes(router *mux.Router) {
	router.HandleFunc(properties.LoginEndpoint, loginHandler).Methods(http.MethodPost)
	router.HandleFunc(properties.RefreshEndpoint, refreshHandler).Methods(http.MethodPost)
	router.HandleFunc("/logout", logoutHandler).Methods(http.MethodPost)
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
	accessTokenCookie, err := r.Cookie(security.AccessTokenCookieKey)
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
	}
	refreshTokenCookie, err := r.Cookie(security.RefreshTokenCookieKey)
	if err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
	}

	accessTokenCookie.MaxAge = -1
	refreshTokenCookie.MaxAge = -1
	http.SetCookie(w, accessTokenCookie)
	http.SetCookie(w, refreshTokenCookie)
}
