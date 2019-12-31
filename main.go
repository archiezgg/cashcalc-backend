package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

var port = ":8080"

func main() {
	router := mux.NewRouter()

	http.HandleFunc("/favicon.ico", faviconHandler)
	
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("frontend/")))
	
	log.Println("CashCalc 2020 is up and running on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/favicon.ico")
}