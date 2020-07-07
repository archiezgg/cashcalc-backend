/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package security

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/IstvanN/cashcalc-backend/repositories"

	"github.com/IstvanN/cashcalc-backend/properties"

	"github.com/IstvanN/cashcalc-backend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var (
	accessKey  = []byte(os.Getenv("ACCESS_KEY"))
	refreshKey = []byte(os.Getenv("REFRESH_KEY"))
)

// CustomClaims is the struct for the Token Claims including role
// and standard JWT claims
type CustomClaims struct {
	jwt.StandardClaims
	Role models.Role `json:"role"`
}

// GenerateAccessToken takes a user as param and generates a signed access token
func GenerateAccessToken(user models.User) (string, error) {
	if string(accessKey) == "" {
		return "", fmt.Errorf("ACCESS_KEY is unset")
	}

	claims := CustomClaims{
		jwt.StandardClaims{
			Issuer:    user.Username,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * properties.AccessTokenExp).Unix(),
		},
		user.Role,
	}

	accessTokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(accessKey)
	if err != nil {
		return "", err
	}
	return accessTokenString, nil
}

// GenerateRefreshToken takes a user as param and generates a signed refresh token
func GenerateRefreshToken(user models.User) (string, error) {
	if string(refreshKey) == "" {
		return "", fmt.Errorf("REFRESH_KEY is unset")
	}

	claims := CustomClaims{
		jwt.StandardClaims{
			Issuer:    user.Username,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * properties.RefreshTokenExp).Unix(),
		},
		user.Role,
	}

	refreshTokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(refreshKey)
	if err != nil {
		return "", err
	}

	if err := repositories.SaveRefreshToken(user, refreshTokenString); err != nil {
		return "", err
	}
	return refreshTokenString, nil
}

// RefreshTokenAndSetTokensAsCookies takes a refresh token and a writer, refreshes the user's token and sends back as cookie
func RefreshTokenAndSetTokensAsCookies(w http.ResponseWriter, refreshToken string) (string, error) {
	user, err := DecodeUserFromRefreshToken(refreshToken)
	if err != nil {
		LogErrorAndSendHTTPError(w, err, http.StatusUnauthorized)
		return "", err
	}

	if err := repositories.DeleteRefreshToken(refreshToken); err != nil {
		LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
		return "", err
	}

	accessToken, err := GenerateTokenPairsAndSetThemAsCookies(w, user)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func generateTokenPairs(user models.User) (accessToken string, newRefreshToken string, err error) {
	at, err := GenerateAccessToken(user)
	if err != nil {
		return "", "", err
	}

	rt, err := GenerateRefreshToken(user)
	if err != nil {
		return "", "", err
	}

	return at, rt, nil
}

// DecodeUserFromRefreshToken decodes the username from JWT refresh token
func DecodeUserFromRefreshToken(refreshTokenString string) (models.User, error) {
	claims, err := decodeClaimsFromToken(refreshTokenString, refreshKey)
	if err != nil {
		return models.User{}, err
	}

	user, err := repositories.GetUserByUsername(claims.Issuer)
	if err != nil {
		return models.User{}, err
	}

	if err := checkIfRefreshTokenIsInDB(user.Username, refreshTokenString); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func checkIfRefreshTokenIsInDB(username string, refreshToken string) error {
	tokenInDB, err := repositories.GetRefreshTokenByUsername(username)
	if err != nil {
		return err
	}

	if tokenInDB.TokenString != refreshToken {
		return fmt.Errorf("refresh token %v for user %v is not in db", refreshToken, username)
	}
	return nil
}
