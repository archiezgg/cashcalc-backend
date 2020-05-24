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
	Username string
	Role     models.Role `json:"role"`
	jwt.StandardClaims
}

// CreateAccessToken takes a Role as param and creates a signed access token
func CreateAccessToken(user models.User) (string, error) {
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

// CreateRefreshToken takes a Role as param and creates a signed refresh token
func CreateRefreshToken(user models.User) (models.RefreshToken, error) {
	if string(refreshKey) == "" {
		return models.RefreshToken{}, fmt.Errorf("REFRESH_KEY is unset")
	}

	tokenExpiration := time.Now().Add(time.Minute * properties.RefreshTokenExp).Unix()

	claims := CustomClaims{
		user.Username,
		user.Role,
		jwt.StandardClaims{
			ExpiresAt: tokenExpiration,
		},
	}

	refreshTokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(refreshKey)
	if err != nil {
		return models.RefreshToken{}, err
	}

	rt := models.RefreshToken{
		Username:    user.Username,
		Role:        user.Role,
		TokenString: refreshTokenString,
		ExpiresAt:   tokenExpiration,
	}

	if err := repositories.SaveRefreshToken(rt); err != nil {
		return models.RefreshToken{}, err
	}
	return rt, nil
}
