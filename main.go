package main

import (
	"context"
	"log"
	"net/http"

	"github.com/IstvanN/cashcalc-backend/controller"
	"github.com/IstvanN/cashcalc-backend/database"
)

var port = ":8080"

func main() {
	router := controller.StartupRouter()
	db := database.Startup()
	defer db.Disconnect(context.TODO())

	log.Println("CashCalc 2020 is up and running on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
