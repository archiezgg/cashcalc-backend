package security

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/IstvanN/cashcalc-backend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var signingKey = []byte("rekettye")

// Claims is the struct for the Token Claims including role
// and standard JWT claims
type Claims struct {
	Role models.Role `json:"role"`
	jwt.StandardClaims
}

// LogErrorAndSendHTTPError takes and error and a http status code, and formats them to
// create proper logging and formatted http respond at the same time
func LogErrorAndSendHTTPError(w http.ResponseWriter, err error, httpStatusCode int) {
	log.Println(err)
	errorMsg := fmt.Sprintf("{\"error\": \"%v\"}", http.StatusText(httpStatusCode))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	w.Write([]byte(errorMsg))
}

// CreateToken takes a Role as param and creates a signed token
func CreateToken(role models.Role) (string, error) {
	claims := Claims{
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// AuthCarrierLevel serves as middleware for carrier access level
func AuthCarrierLevel(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isTokenValidForRole(models.Carrier, w, r) {
			next.ServeHTTP(w, r)
		}
	})
}

// AuthAdminLevel serves as middleware for admin access level
func AuthAdminLevel(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isTokenValidForRole(models.Admin, w, r) {
			next.ServeHTTP(w, r)
		}
	})
}

// AuthSuperuserLevel serves as middleware for superuser access level
func AuthSuperuserLevel(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isTokenValidForRole(models.Superuser, w, r) {
			next.ServeHTTP(w, r)
		}
	})
}

func isTokenValidForRole(role models.Role, w http.ResponseWriter, r *http.Request) bool {
	tokenStrings, ok := r.Header["Token"]
	if !ok {
		LogErrorAndSendHTTPError(w, fmt.Errorf("no token in header"), http.StatusUnauthorized)
		return false
	}

	var claims Claims
	token, err := jwt.ParseWithClaims(tokenStrings[0], &claims, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil || !token.Valid {
		LogErrorAndSendHTTPError(w, err, http.StatusUnauthorized)
		return false
	}

	if role == models.Admin && claims.Role == models.Carrier {
		LogErrorAndSendHTTPError(w, fmt.Errorf("%v is trying to reach content restricted for %v", claims.Role, role), http.StatusForbidden)
		return false
	}

	if role == models.Superuser && claims.Role != models.Superuser {
		LogErrorAndSendHTTPError(w, fmt.Errorf("%v is trying to reach content restricted for %v", claims.Role, role), http.StatusForbidden)
		return false
	}

	return true
}
