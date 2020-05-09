/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package controllers

import (
	"net/http"
	"net/http/pprof"

	"github.com/IstvanN/cashcalc-backend/security"
	"github.com/gorilla/mux"
)

func registerDebugRoutes(router *mux.Router) {
	s := router.PathPrefix("/debug").Subrouter()
	s.HandleFunc("", pprof.Index)
	s.Use(setHTMLHeader)
	s.Use(security.AccessLevelSuperuser)
}

func setHTMLHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		next.ServeHTTP(w, r)
	})
}
