package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IstvanN/cashcalc-backend/models"
	"github.com/IstvanN/cashcalc-backend/properties"
	"github.com/IstvanN/cashcalc-backend/repositories"
	"golang.org/x/crypto/bcrypt"

	"github.com/IstvanN/cashcalc-backend/security"

	"github.com/gorilla/mux"
)

func registerLoginRoutes(router *mux.Router) {
	ep := properties.LoginEndpoint
	router.HandleFunc(ep, loginHandler).Methods(http.MethodPost)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var userToAuth models.User
	if err := json.NewDecoder(r.Body).Decode(&userToAuth); err != nil {
		LogErrorAndSendHTTPError(w, err, http.StatusUnprocessableEntity)
		return
	}

	u, err := repositories.GetUserByRole(userToAuth.Role)
	if err != nil {
		LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(userToAuth.Password))
	if err != nil {
		err := fmt.Errorf("the given password is invalid: %v", u.Password)
		LogErrorAndSendHTTPError(w, err, http.StatusUnauthorized)
		return
	}

	st, err := security.CreateToken(u.Role)
	if err != nil {
		LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Token", st)
}
