/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// StartupRouter creates instance of registers all the routes of the subroutes, supposed to be called in main func
func StartupRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/", welcomeHandler).Methods("GET")
	registerLoginRoutes(router)
	registerCountriesRoutes(router)
	registerPricingsRoutes(router)
	registerPricingVarsRoutes(router)
	router.Use(setJSONHeaderMiddleWare)
	return
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "Welcome to CashCalc 2020"}`))
}

// setJSONHeaderMiddleWare sets the header to application/json for a given handler
func setJSONHeaderMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
