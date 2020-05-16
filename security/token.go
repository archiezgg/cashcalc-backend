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

	"github.com/IstvanN/cashcalc-backend/properties"

	"github.com/IstvanN/cashcalc-backend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var (
	accessKey     = []byte(os.Getenv("ACCESS_KEY"))
	refreshKey    = []byte(os.Getenv("REFRESH_KEY"))
	refreshTokens map[string]models.Role // TODO implement it with Redis
)

func init() {
	refreshTokens = make(map[string]models.Role)
}

// CustomClaims is the struct for the Token Claims including role
// and standard JWT claims
type CustomClaims struct {
	Role models.Role `json:"role"`
	jwt.StandardClaims
}

// CreateAccessToken takes a Role as param and creates a signed access token
func CreateAccessToken(role models.Role) (string, error) {
	if string(accessKey) == "" {
		return "", fmt.Errorf("ACCESS_KEY is unset")
	}

	claims := CustomClaims{
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * properties.AccessTokenExp).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString(accessKey)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

// CreateRefreshToken takes a Role as param and creates a signed refresh token
func CreateRefreshToken(role models.Role) (string, error) {
	if string(refreshKey) == "" {
		return "", fmt.Errorf("REFRESH_KEY is unset")
	}

	claims := CustomClaims{
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * properties.RefreshTokenExp).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := token.SignedString(refreshKey)
	if err != nil {
		return "", err
	}
	refreshTokens[refreshToken] = role
	return refreshToken, nil
}

// GetRoleFromRefreshToken takes a token as a string and returns with the role if token is valid
func GetRoleFromRefreshToken(refreshToken string) (models.Role, error) {
	role, ok := refreshTokens[refreshToken]
	if !ok {
		return "", fmt.Errorf("the refresh token %v is not in database", refreshToken)
	}
	return role, nil
}

// DeleteRefreshTokenFromMemory deletes the refresh token from the in-memory DB
func DeleteRefreshTokenFromMemory(rt string) {
	delete(refreshTokens, rt)
}
