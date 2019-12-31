package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"

)

// Startup registers all the routes of the subroutes, supposed to be called in main func
func Startup(router *httprouter.Router) {
	router.GET("/favicon.ico", faviconHandler)
	registerCountriesRoutes(router)
}

func faviconHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "frontend/favicon.ico")
}