/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package repositories

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/IstvanN/cashcalc-backend/models"

	"github.com/IstvanN/cashcalc-backend/database"
	"github.com/IstvanN/cashcalc-backend/properties"
)

// GetRefreshTokenByUsername retrieves the token by the user's name
func GetRefreshTokenByUsername(username string) (models.RefreshToken, error) {
	tokenJSON, err := database.RedisClient().Get(username).Result()
	if err != nil {
		return models.RefreshToken{}, fmt.Errorf("error retreiving refresh token for %v: %v", username, err)
	}

	var token models.RefreshToken
	if err := json.Unmarshal([]byte(tokenJSON), &token); err != nil {
		return models.RefreshToken{}, fmt.Errorf("error unmarshaling token: %v", tokenJSON)
	}
	return token, nil
}

// SaveRefreshToken saves the username as key and the token data as JSON as value
func SaveRefreshToken(user models.User, refreshTokenString string) error {
	exp := time.Minute * properties.RefreshTokenExp
	token := models.RefreshToken{
		Username:    user.Username,
		Role:        user.Role,
		TokenString: refreshTokenString,
		ExpiresAt:   time.Now().Add(exp).Unix(),
	}

	tokenJSON, err := json.Marshal(token)
	if err != nil {
		return err
	}

	err = database.RedisClient().Set(user.Username, tokenJSON, exp).Err()
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
	usernames, err := database.RedisClient().Keys("*").Result()
	if err != nil {
		return nil, err
	}

	var tokens []models.RefreshToken
	for _, username := range usernames {
		token, err := GetRefreshTokenByUsername(username)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
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
