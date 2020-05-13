/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// StartupRouter creates instance of registers all the routes of the subroutes, supposed to be called in main func
func StartupRouter() (router *mux.Router) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("cashcalc-backend"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
	)
	if err != nil {
		log.Fatalln(err)
	}

	router = mux.NewRouter()

	router.HandleFunc(newrelic.WrapHandleFunc(app, "/", welcomeHandler)).Methods("GET")
	registerLoginRoutes(router)
	registerCountriesRoutes(router)
	registerPricingsRoutes(router)
	registerPricingVarsRoutes(router)
	registerDebugRoutes(router)
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
