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
	ep := properties.LoginEndpoint
	router.HandleFunc(ep, loginHandler).Methods(http.MethodPost)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		logErrorAndSendHTTPError(w, err, http.StatusUnprocessableEntity)
		return
	}

	if u.Password != models.TestUser.Password {
		err := fmt.Errorf("the given password is invalid: %v", u.Password)
		logErrorAndSendHTTPError(w, err, http.StatusUnauthorized)
		return
	}

	st, err := security.CreateToken("carrier")
	if err != nil {
		logErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Token", st)
}
