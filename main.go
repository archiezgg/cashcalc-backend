package main

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/IstvanN/cashcalc-backend/controller"
)

var port = ":8080"

func main() {
	router := httprouter.New()

	controller.Startup(router)

	log.Println("CashCalc 2020 is up and running on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}