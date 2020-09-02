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
	router.HandleFunc(properties.LoginEndpoint, loginHandler).Methods(http.MethodPost, http.MethodOptions)
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

	user, err := security.AuthenticateUser(w, userToAuth)
	if err != nil {
		return
	}
	msg := fmt.Sprintf("{\"message\": \"%s\", \"username\": \"%v\", \"role\": \"%v\"}", "Logged in successfully", user.Username, user.Role)
	w.Write([]byte(msg))
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if err := security.DeleteTokensFromCookies(w, r); err != nil {
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}
	writeMessage(w, "Logged out succesfully")
}

func isAuthorizedHandler(w http.ResponseWriter, r *http.Request) {
	roleToCompare := r.URL.Query().Get("role")
	if roleToCompare == "" {
		err := fmt.Errorf("role parameter not found in URL query")
		security.LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	if !security.IsTokenValidForAccessLevel(models.Role(roleToCompare), w, r) {
		return
	}

	writeMessage(w, "Authorized")
}
