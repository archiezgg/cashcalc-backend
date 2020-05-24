/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package repositories

import (
	"fmt"
	"time"

	"github.com/IstvanN/cashcalc-backend/models"

	"github.com/IstvanN/cashcalc-backend/database"
	"github.com/IstvanN/cashcalc-backend/properties"
)

// GetRefreshTokenByUsername retrieves the token by the user's name
func GetRefreshTokenByUsername(username string) (string, error) {
	tokenString, err := database.RedisClient().Get(username).Result()
	if err != nil {
		return "", fmt.Errorf("error retreiving refresh token for %v: %v", username, err)
	}
	return tokenString, nil
}

// SaveRefreshToken saves the token with the role to the DB
func SaveRefreshToken(username string, refreshTokenString string) error {
	err := database.RedisClient().Set(username, refreshTokenString, time.Minute*properties.RefreshTokenExp).Err()
	if err != nil {
		return fmt.Errorf("error saving refresh token: %v", err)
	}
	return nil
}

// DeleteRefreshToken deletes the given user's refresh token from DB
func DeleteRefreshToken(username string) error {
	err := database.RedisClient().Del(username).Err()
	if err != nil {
		return fmt.Errorf("error deleting refresh token for user %v: %v", username, err)
	}
	return nil
}

// DeleteBulkRefreshToken deletes multiple refresh tokens from DB
func DeleteBulkRefreshToken(usernames []string) error {
	for _, user := range usernames {
		err := DeleteRefreshToken(user)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetAllTokens returns with an array containing all refresh tokens
func GetAllTokens() ([]models.RefreshToken, error) {
	var tokens []models.RefreshToken
	usernames, err := database.RedisClient().Keys("*").Result()
	if err != nil {
		return nil, err
	}

	for _, username := range usernames {
		token, err := createRefreshTokenDataFromUsername(username)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

func createRefreshTokenDataFromUsername(username string) (models.RefreshToken, error) {
	tokenString, err := GetRefreshTokenByUsername(username)
	if err != nil {
		return models.RefreshToken{}, err
	}
	expDate, err := getExpirationDateForToken(username)
	if err != nil {
		return models.RefreshToken{}, err
	}
	user, err := GetUserByUsername(username)
	if err != nil {
		return models.RefreshToken{}, err
	}

	token := models.RefreshToken{
		Username:    username,
		Role:        user.Role,
		TokenString: tokenString,
		ExpiresAt:   expDate,
	}
	return token, nil
}

// DeleteAllTokens removes all tokens from DB
func DeleteAllTokens() error {
	if err := database.RedisClient().FlushDB().Err(); err != nil {
		return fmt.Errorf("error flushing all tokens from DB")
	}
	return nil
}

func getExpirationDateForToken(username string) (int64, error) {
	expTime, err := database.RedisClient().PTTL(username).Result()
	if err != nil {
		return 0, err
	}
	expDate := time.Now().Add(expTime).Unix()
	return expDate, nil
}
