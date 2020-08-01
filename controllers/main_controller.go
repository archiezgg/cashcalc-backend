/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// StartupRouter creates instance of registers all the routes of the subroutes, supposed to be called in main func
func StartupRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/", welcomeHandler).Methods(http.MethodGet)
	registerLoginRoutes(router)
	registerCountriesRoutes(router)
	registerPricingsRoutes(router)
	registerPricingVarsRoutes(router)
	registerTokenRoutes(router)
	registerUserRoutes(router)
	registerCalcRoutes(router)
	router.Use(setHeaderMiddleWare)
	return
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	writeMessage(w, "Welcome to CashCalc!")
}

// setHeaderMiddleWare sets the header with some pre-made CORS-enabling options
func setHeaderMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		// w.Header().Set("Access-Control-Allow-Origin", "https://cashcalc.web.app")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-PINGOTHER")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func writeMessage(w http.ResponseWriter, msg string) {
	finalMessage := fmt.Sprintf("{\"message\": \"%s\"}", msg)
	w.Write([]byte(finalMessage))
}
