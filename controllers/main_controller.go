/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var isProcessOngoing bool

// StartupRouter creates instance of registers all the routes of the subroutes, supposed to be called in main func
func StartupRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/", welcomeHandler).Methods(http.MethodGet, http.MethodOptions)
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
	if isProcessOngoing {
		time.Sleep(time.Second * 1)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-PINGOTHER")

		if os.Getenv("ENVIRONMENT") == "DEV" {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "https://cashcalc.web.app")
		}

		if r.Method == "OPTIONS" {
			return
		}

		isProcessOngoing = true
		next.ServeHTTP(w, r)
		isProcessOngoing = false
	})
}

func writeMessage(w http.ResponseWriter, msg string) {
	finalMessage := fmt.Sprintf("{\"message\": \"%s\"}", msg)
	w.Write([]byte(finalMessage))
}
