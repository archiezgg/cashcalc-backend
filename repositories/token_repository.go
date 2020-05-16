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
	"github.com/IstvanN/cashcalc-backend/models"
	"github.com/IstvanN/cashcalc-backend/properties"
)

// GetRoleFromRefreshToken retrieves the role to the corresponding refresh token
func GetRoleFromRefreshToken(refreshToken string) (models.Role, error) {
	role, err := database.RedisClient().Get(refreshToken).Result()
	if err != nil {
		return "", fmt.Errorf("error retreiving refresh token %v: %v", refreshToken, err)
	}
	return models.Role(role), nil
}

// SaveRefreshToken saves the token with the role to the DB
func SaveRefreshToken(refreshToken string, role models.Role) error {
	err := database.RedisClient().Set(refreshToken, string(role), time.Minute*properties.RefreshTokenExp).Err()
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
