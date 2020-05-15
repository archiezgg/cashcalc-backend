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

	"github.com/IstvanN/cashcalc-backend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var accessKey = []byte(os.Getenv("ACCESS_KEY"))

// CustomClaims is the struct for the Token Claims including role
// and standard JWT claims
type CustomClaims struct {
	Role models.Role `json:"role"`
	jwt.StandardClaims
}

// CreateAccessToken takes a Role as param and creates a signed token
func CreateAccessToken(role models.Role) (string, error) {
	if string(accessKey) == "" {
		return "", fmt.Errorf("ACCESS_KEY is unset")
	}

	claims := CustomClaims{
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(accessKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
