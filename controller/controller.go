package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// StartupRouter creates instance of registers all the routes of the subroutes, supposed to be called in main func
func StartupRouter() (router *httprouter.Router) {
	router = httprouter.New()
	router.GET("/favicon.ico", faviconHandler)
	registerCountriesRoutes(router)
	return
}

func faviconHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "frontend/favicon.ico")
}

func setContentTypeToJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
