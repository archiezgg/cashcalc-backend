/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoginWithNoBody(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/login", nil)
	w := httptest.NewRecorder()
	loginHandler(w, r)

	expected := http.StatusUnprocessableEntity
	actual := w.Code
	if actual != expected {
		t.Errorf("%v endpoint failed: expected status code %v, got %v", "/login", expected, actual)
	}
}

func TestLoginWithBadJSON(t *testing.T) {
	body := strings.NewReader("{\"bad\": \"request\"}")
	r := httptest.NewRequest(http.MethodPost, "/login", body)
	w := httptest.NewRecorder()
	loginHandler(w, r)

	expected := http.StatusUnprocessableEntity
	actual := w.Code
	if actual != expected {
		t.Errorf("%v endpoint failed: expected status code %v, got %v", "/login", expected, actual)
	}
}
