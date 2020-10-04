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
	"github.com/IstvanN/cashcalc-backend/repositories"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const (
	// AccessTokenHeaderKey is the key for access tokens in headers
	AccessTokenHeaderKey = "access-token"
	// RefreshTokenHeaderKey is the key for refresh tokens in headers
	RefreshTokenHeaderKey = "refresh-token"
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
		if IsTokenValidForAccessLevel(models.RoleCarrier, w, r) {
			next.ServeHTTP(w, r)
		}
	})
}

// AccessLevelAdmin serves as middleware for admin access level
func AccessLevelAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if IsTokenValidForAccessLevel(models.RoleAdmin, w, r) {
			next.ServeHTTP(w, r)
		}
	})
}

// AccessLevelSuperuser serves as middleware for superuser access level
func AccessLevelSuperuser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if IsTokenValidForAccessLevel(models.RoleSuperuser, w, r) {
			next.ServeHTTP(w, r)
		}
	})
}

// IsTokenValidForAccessLevel checks if the role in the token can get the resources
// for given access level
func IsTokenValidForAccessLevel(accessLevel models.Role, w http.ResponseWriter, r *http.Request) bool {
	var token string
	var err error

	token, err = extractTokenFromHeader(w, r)
	if err != nil {
		LogErrorAndSendHTTPError(w, err, http.StatusUnauthorized)
		return false
	}

	role, err := decodeRoleFromAccessToken(token)
	if err != nil {
		LogErrorAndSendHTTPError(w, err, http.StatusUnauthorized)
		return false
	}

	if accessLevel != models.RoleCarrier && accessLevel != models.RoleAdmin && accessLevel != models.RoleSuperuser {
		err := fmt.Errorf("given role can either be %v, %v or %v", models.RoleCarrier, models.RoleAdmin, models.RoleSuperuser)
		LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return false
	}

	if err := checkAccessLevel(role, accessLevel); err != nil {
		LogErrorAndSendHTTPError(w, err, http.StatusForbidden)
		return false
	}

	return true
}

// AuthenticateUser takes a user model, and checks if the credentials are valid,
// returns with the user if yes, returns error if not
func AuthenticateUser(w http.ResponseWriter, userToAuth models.User) (models.User, error) {
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
	if _, err := GenerateTokenPairsForUserAndSetThemAsHeaders(w, u); err != nil {
		return models.User{}, err
	}
	log.Printf("user '%v' has successfully logged in", u.Username)
	return u, nil
}

// LogoutUser takes the refresh token from the header and deletes it from the DB
func LogoutUser(r *http.Request) error {
	refreshTokenString := r.Header.Get(RefreshTokenHeaderKey)
	if err := repositories.DeleteRefreshTokenByTokenString(refreshTokenString); err != nil {
		return err
	}
	return nil
}

// GenerateTokenPairsForUserAndSetThemAsHeaders generate access- and refresh token,
// saves them for given user in DB,
// sets them as http headers, and returns with the access token
func GenerateTokenPairsForUserAndSetThemAsHeaders(w http.ResponseWriter, user models.User) (string, error) {
	at, rt, err := generateTokenPairs(user)
	if err != nil {
		LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return "", err
	}

	w.Header().Set(AccessTokenHeaderKey, at)
	w.Header().Set(RefreshTokenHeaderKey, rt)

	return at, nil
}

func generateAccessTokenAndSetItAsHeader(w http.ResponseWriter, user models.User) (string, error) {
	at, err := GenerateAccessToken(user)
	if err != nil {
		LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return "", err
	}

	w.Header().Set(AccessTokenHeaderKey, at)
	return at, nil
}

func extractTokenFromHeader(w http.ResponseWriter, r *http.Request) (string, error) {
	accessToken, err := validateAndGetAccessTokenFromHeader(r)
	if err == nil {
		return accessToken, nil
	}

	refreshToken := r.Header.Get(RefreshTokenHeaderKey)
	if refreshToken == "" {
		return "", fmt.Errorf("no header specified as %v", RefreshTokenHeaderKey)
	}

	newAccessToken, err := RefreshTokenAndSetTokensAsHeaders(w, refreshToken)
	if err != nil {
		return "", err
	}
	return newAccessToken, nil
}

func validateAndGetAccessTokenFromHeader(r *http.Request) (string, error) {
	accessToken := r.Header.Get(AccessTokenHeaderKey)
	if accessToken == "" {
		return "", fmt.Errorf("no header specified as %v", AccessTokenHeaderKey)
	}

	_, err := decodeClaimsFromToken(accessToken, accessKey)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

// decodeRoleFromAccessToken decodes the role from the access token JWT string
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
