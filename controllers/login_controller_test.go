/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package controllers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	method   = http.MethodPost
	endpoint = "/login"
)

func executeRequest(body io.Reader) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, endpoint, body)
	w := httptest.NewRecorder()
	loginHandler(w, r)
	return w
}

func TestLoginWithNoBody(t *testing.T) {
	w := executeRequest(nil)

	expected := http.StatusUnprocessableEntity
	actual := w.Code
	if actual != expected {
		t.Errorf("%v endpoint failed: expected status code %v, got %v", "/login", expected, actual)
	}
}

func TestLoginWithBadJSON(t *testing.T) {
	body := strings.NewReader("{\"bad\": \"request\"}")
	w := executeRequest(body)

	expected := http.StatusUnprocessableEntity
	actual := w.Code
	if actual != expected {
		t.Errorf("%v endpoint failed: expected status code %v, got %v", "/login", expected, actual)
	}
}
