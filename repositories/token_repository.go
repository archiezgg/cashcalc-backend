/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package repositories

import (
	"github.com/IstvanN/cashcalc-backend/database"
	"github.com/IstvanN/cashcalc-backend/models"
	"github.com/IstvanN/cashcalc-backend/properties"
)

// GetRoleFromRefreshToken retrieves the role to the corresponding refresh token
func GetRoleFromRefreshToken(rt string) (models.Role, error) {
	role, err := database.RedisClient().Get(rt).Result()
	if err != nil {
		return "", err
	}
	return models.Role(role), nil
}

// SaveRefreshToken saves the token with the role to the DB
func SaveRefreshToken(rt string, role models.Role) error {
	err := database.RedisClient().Set(rt, role, properties.RefreshTokenExp).Err()
	return err
}
