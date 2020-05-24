/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package repositories

import (
	"fmt"
	"time"

	"github.com/IstvanN/cashcalc-backend/database"
	"github.com/IstvanN/cashcalc-backend/properties"
)

// GetUsernameFromRefreshToken retrieves the username to the corresponding refresh token
func GetUsernameFromRefreshToken(refreshToken string) (string, error) {
	username, err := database.RedisClient().Get(refreshToken).Result()
	if err != nil {
		return "", fmt.Errorf("error retreiving refresh token %v: %v", refreshToken, err)
	}
	return username, nil
}

// SaveRefreshToken saves the token with the role to the DB
func SaveRefreshToken(username string, refreshTokenString string) error {
	err := database.RedisClient().Set(username, refreshTokenString, time.Minute*properties.RefreshTokenExp).Err()
	if err != nil {
		return fmt.Errorf("error saving refresh token: %v", err)
	}
	return nil
}

// DeleteRefreshToken deletes the given refresh token from DB
func DeleteRefreshToken(refreshToken string) error {
	err := database.RedisClient().Del(refreshToken).Err()
	if err != nil {
		return fmt.Errorf("error deleting refresh token: %v", err)
	}
	return nil
}

// DeleteBulkRefreshToken deletes multiple refresh tokens from DB
func DeleteBulkRefreshToken(refreshTokens []string) error {
	for _, rt := range refreshTokens {
		err := DeleteRefreshToken(rt)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetAllTokens returns with a map of string:string containing all token data
func GetAllTokens() (map[string]string, error) {
	tokens, err := database.RedisClient().Keys("*").Result()
	if err != nil {
		return nil, err
	}

	tokensMap := make(map[string]string)
	for _, t := range tokens {
		username, err := GetUsernameFromRefreshToken(t)
		if err != nil {
			return nil, err
		}
		tokensMap[t] = username
	}
	return tokensMap, nil
}

// DeleteAllTokens removes all tokens from DB
func DeleteAllTokens() error {
	if err := database.RedisClient().FlushDB().Err(); err != nil {
		return fmt.Errorf("error flushing all tokens from DB")
	}
	return nil
}
