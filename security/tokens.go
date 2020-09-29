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

	"github.com/IstvanN/cashcalc-backend/properties"
	"github.com/IstvanN/cashcalc-backend/repositories"

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

func createCustomClaims(user models.User) CustomClaims {
	return CustomClaims{
		jwt.StandardClaims{
			Issuer:   user.Username,
			IssuedAt: time.Now().Unix(),
		},
		user.Role,
	}
}

// GenerateAccessToken takes a user as param and generates a signed access token
func GenerateAccessToken(user models.User) (string, error) {
	if string(accessKey) == "" {
		return "", fmt.Errorf("ACCESS_KEY is unset")
	}

	claims := createCustomClaims(user)
	claims.ExpiresAt = time.Now().Add(time.Minute * properties.AccessTokenExp).Unix()

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

	claims := createCustomClaims(user)
	claims.ExpiresAt = time.Now().Add(time.Minute * properties.RefreshTokenExp).Unix()

	refreshTokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(refreshKey)
	if err != nil {
		return "", err
	}

	_, err = repositories.CreateRefreshTokenAndSaveForUser(user, refreshTokenString, time.Unix(claims.ExpiresAt, 0))
	if err != nil {
		return "", err
	}

	return refreshTokenString, nil
}

// RefreshTokenAndSetTokensAsHeaders takes a refresh token and a writer,
// refreshes the user's token and sends back as headers
func RefreshTokenAndSetTokensAsHeaders(w http.ResponseWriter, refreshToken string) (string, error) {
	user, err := DecodeUserFromRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}

	var accessToken string
	if properties.RefreshTokenRotationOn {
		if err := repositories.DeleteRefreshTokenByTokenString(refreshToken); err != nil {
			LogErrorAndSendHTTPError(w, err, http.StatusInternalServerError)
			return "", err
		}
		accessToken, err = GenerateTokenPairsForUserAndSetThemAsHeaders(w, user)
	} else {
		accessToken, err = generateAccessTokenAndSetItAsHeader(w, user)
	}
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

	if err := checkIfRefreshTokenIsInDB(user, refreshTokenString); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func checkIfRefreshTokenIsInDB(user models.User, refreshTokenString string) error {
	tokensInDB, err := repositories.GetRefreshTokensByUserID(user.ID)
	if err != nil {
		return err
	}

	for _, t := range tokensInDB {
		if t.TokenString == refreshTokenString {
			return nil
		}
	}
	return fmt.Errorf("refresh token %v for user %v is not in db", refreshTokenString, user.Username)
}
