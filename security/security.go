/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package security

import (
	"fmt"
	"log"
	"net/http"

	"github.com/IstvanN/cashcalc-backend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// LogErrorAndSendHTTPError takes and error and a http status code, and formats them to
// create proper logging and formatted http respond at the same time
func LogErrorAndSendHTTPError(w http.ResponseWriter, err error, httpStatusCode int) {
	log.Println(err)
	errorMsg := fmt.Sprintf("{\"error\": \"%v\"}", http.StatusText(httpStatusCode))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	w.Write([]byte(errorMsg))
}

// AccessLevelCarrier serves as middleware for carrier access level
func AccessLevelCarrier(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isTokenValidForAccessLevel(models.Carrier, w, r) {
			next.ServeHTTP(w, r)
		}
	})
}

// AccessLevelAdmin serves as middleware for admin access level
func AccessLevelAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isTokenValidForAccessLevel(models.Admin, w, r) {
			next.ServeHTTP(w, r)
		}
	})
}

// AccessLevelSuperuser serves as middleware for superuser access level
func AccessLevelSuperuser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isTokenValidForAccessLevel(models.Superuser, w, r) {
			next.ServeHTTP(w, r)
		}
	})
}

func isTokenValidForAccessLevel(accessLevel models.Role, w http.ResponseWriter, r *http.Request) bool {
	tokenStrings, ok := r.Header["Access-Token"]
	if !ok {
		LogErrorAndSendHTTPError(w, fmt.Errorf("no token in header"), http.StatusUnauthorized)
		return false
	}

	role, err := getRoleFromToken(tokenStrings[0], accessKey)
	if err != nil {
		LogErrorAndSendHTTPError(w, err, http.StatusUnauthorized)
		return false
	}

	if err := checkAccessLevel(role, accessLevel); err != nil {
		LogErrorAndSendHTTPError(w, err, http.StatusForbidden)
		return false
	}

	return true
}

func getRoleFromToken(tokenString string, key []byte) (models.Role, error) {
	var claims CustomClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil || !token.Valid {
		return "", err
	}

	return claims.Role, nil
}

func checkAccessLevel(role, accessLevel models.Role) error {
	err := fmt.Errorf("%v is trying to reach content restricted for %v", role, accessLevel)
	if accessLevel == models.Admin && role == models.Carrier {
		return err
	}
	if accessLevel == models.Superuser && role != models.Superuser {
		return err
	}
	return nil
}
