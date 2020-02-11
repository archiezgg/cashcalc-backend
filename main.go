package main

import (
	"log"
	"net/http"
	"os"

	"github.com/IstvanN/cashcalc-backend/controller"
	"github.com/IstvanN/cashcalc-backend/database"
)

var port = ":" + os.Getenv("PORT")

func main() {
	db := database.Startup()
	defer db.Close()

	router := controller.StartupRouter()

	log.Println("CashCalc 2020 is up and running on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
