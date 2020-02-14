package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
)

// StartupRouter creates instance of registers all the routes of the subroutes, supposed to be called in main func
func StartupRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/", welcomeHandler).Methods("GET")
	registerCountriesRoutes(router)
	registerPricingsRoutes(router)
	return
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to CashCalc 2020!"))
}

func setContentTypeToJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
