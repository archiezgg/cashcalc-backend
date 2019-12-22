package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

var (
	port = ":8080"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler).Methods("GET")
	log.Println("CashCalc 2020 is up and running on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to CashCalc 2020!")
}