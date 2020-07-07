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
	"strings"

	"github.com/IstvanN/cashcalc-backend/models"
	"github.com/IstvanN/cashcalc-backend/repositories"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// LogErrorAndSendHTTPError takes and error and a http status code, and formats them to
// create proper logging and formatted http respond at the same time
func LogErrorAndSendHTTPError(w http.ResponseWriter, err error, httpStatusCode int) {
	log.Println(err)
	errorMsg := fmt.Sprintf("{\"error\": \"%v\"}", err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	w.Write([]byte(errorMsg))
}

// AccessLevelCarrier serves as middleware for carrier access level
func AccessLevelCarrier(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isTokenValidForAccessLevel(models.RoleCarrier, w, r) {
			next.ServeHTTP(w, r)
		}
	})
}

// AccessLevelAdmin serves as middleware for admin access level
func AccessLevelAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isTokenValidForAccessLevel(models.RoleAdmin, w, r) {
			next.ServeHTTP(w, r)
		}
	})
}

// AccessLevelSuperuser serves as middleware for superuser access level
func AccessLevelSuperuser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isTokenValidForAccessLevel(models.RoleSuperuser, w, r) {
			next.ServeHTTP(w, r)
		}
	})
}

func isTokenValidForAccessLevel(accessLevel models.Role, w http.ResponseWriter, r *http.Request) bool {
	var token string
	var err error

	token, err = extractTokenFromCookie(w, r)
	if err != nil {
		token, err = extractTokenFromHeader(r)
		if err != nil {
			LogErrorAndSendHTTPError(w, err, http.StatusUnauthorized)
			return false
		}
	}

	role, err := decodeRoleFromAccessToken(token)
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

// AuthenticateNewUser takes a user model, and checks if the credentials are valid, returns with the user, returns error if not
func AuthenticateNewUser(w http.ResponseWriter, userToAuth models.User) (models.User, error) {
	u, err := repositories.GetUserByUsername(userToAuth.Username)
	if err != nil {
		LogErrorAndSendHTTPError(w, err, http.StatusUnauthorized)
		return models.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(userToAuth.Password))
	if err != nil {
		err := fmt.Errorf("the given role-password combination is invalid: %v - %v", userToAuth.Username, userToAuth.Password)
		LogErrorAndSendHTTPError(w, err, http.StatusUnauthorized)
		return models.User{}, err
	}
	if _, err := GenerateTokenPairsAndSetThemAsCookies(w, u); err != nil {
		return models.User{}, err
	}
	log.Printf("user '%v' has successfully logged in", u.Username)
	return u, nil
}

// GenerateTokenPairsAndSetThemAsCookies generate access- and refresh token, sets them as http headers, and returns with the access token
func GenerateTokenPairsAndSetThemAsCookies(w http.ResponseWriter, user models.User) (string, error) {
	at, rt, err := generateTokenPairs(user)
	if err != nil {
		LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return "", err
	}
	accessTokenCookie := &http.Cookie{
		Name:     "access-token",
		Value:    at,
		HttpOnly: true,
	}

	refreshTokenCookie := &http.Cookie{
		Name:     "refresh-token",
		Value:    rt,
		HttpOnly: true,
	}

	http.SetCookie(w, accessTokenCookie)
	http.SetCookie(w, refreshTokenCookie)
	return at, nil
}

func extractTokenFromCookie(w http.ResponseWriter, r *http.Request) (string, error) {
	accessTokenCookie, err := validateAccessTokenCookie(r)
	if err == nil {
		return accessTokenCookie.Value, nil
	}

	refreshTokenCookie, err := r.Cookie("refresh-token")
	if err != nil {
		return "", err
	}

	accessToken, err := RefreshTokenAndSetTokensAsCookies(w, refreshTokenCookie.Value)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func validateAccessTokenCookie(r *http.Request) (*http.Cookie, error) {
	accessTokenCookie, err := r.Cookie("access-token")
	if err != nil {
		return &http.Cookie{}, err
	}

	_, err = decodeClaimsFromToken(accessTokenCookie.Value, accessKey)
	if err != nil {
		return &http.Cookie{}, err
	}

	return accessTokenCookie, nil
}

func extractTokenFromHeader(r *http.Request) (string, error) {
	tokenStrings, ok := r.Header["Authorization"]
	if !ok {
		return "", fmt.Errorf("no token provided in header or cookie")
	}

	bearerToken := tokenStrings[0]
	sliced := strings.Split(bearerToken, " ")
	if len(sliced) != 2 {
		return "", fmt.Errorf("token format is not \"Bearer <Token>\"")
	}
	return sliced[1], nil
}

func decodeRoleFromAccessToken(tokenString string) (models.Role, error) {
	claims, err := decodeClaimsFromToken(tokenString, accessKey)
	if err != nil {
		return "", err
	}

	return claims.Role, nil
}

func decodeClaimsFromToken(tokenString string, key []byte) (CustomClaims, error) {
	var claims CustomClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil || !token.Valid {
		return CustomClaims{}, err
	}

	return claims, nil
}

func checkAccessLevel(role, accessLevel models.Role) error {
	err := fmt.Errorf("%v is trying to reach content restricted for %v", role, accessLevel)
	if accessLevel == models.RoleAdmin && role == models.RoleCarrier {
		return err
	}
	if accessLevel == models.RoleSuperuser && role != models.RoleSuperuser {
		return err
	}
	return nil
}
