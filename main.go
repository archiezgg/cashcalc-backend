package main

import (
	"github.com/IstvanN/cashcalc-backend/controller"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

var port = ":8080"

func main() {
	controller.Startup()

	log.Println("CashCalc 2020 is up and running on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
