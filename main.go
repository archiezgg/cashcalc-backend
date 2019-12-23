package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

const apiPathPrefix = "/api/v1"

var (
	port = ":8080"
)

func main() {
	router := mux.NewRouter()

	http.HandleFunc("/favicon.ico", faviconHandler)
	
	api := router.PathPrefix(apiPathPrefix).Subrouter()
	api.HandleFunc("/countries", countriesHandler).Methods("GET")

	buildHandler := http.FileServer(http.Dir("frontend/"))
	router.PathPrefix("/").Handler(buildHandler)

	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("frontend/static")))
	router.PathPrefix("/static/").Handler(staticHandler)
	
	log.Println("CashCalc 2020 is up and running on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/favicon.ico")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
		
}

func countriesHandler(w http.ResponseWriter, r *http.Request) {
	countries := getAirCountriesFromJSON()
	json.NewEncoder(w).Encode(countries)
}