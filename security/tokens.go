/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package security

import (
	"fmt"
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
	Username string      `json:"username"`
	Role     models.Role `json:"role"`
	jwt.StandardClaims
}

// GenerateAccessToken takes a user as param and generates a signed access token
func GenerateAccessToken(user models.User) (string, error) {
	if string(accessKey) == "" {
		return "", fmt.Errorf("ACCESS_KEY is unset")
	}

	claims := CustomClaims{
		user.Username,
		user.Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * properties.AccessTokenExp).Unix(),
		},
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
		user.Username,
		user.Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * properties.RefreshTokenExp).Unix(),
		},
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

// DecodeUserFromRefreshToken decodes the username from JWT refresh token
func DecodeUserFromRefreshToken(refreshTokenString string) (models.User, error) {
	claims, err := decodeClaimsFromToken(refreshTokenString, refreshKey)
	if err != nil {
		return models.User{}, err
	}

	user, err := repositories.GetUserByUsername(claims.Username)
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
