/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package repositories

import (
	"time"

	"github.com/IstvanN/cashcalc-backend/database"
	"github.com/IstvanN/cashcalc-backend/models"
)

// CreateRefreshToken creates a refresh token struct from data
// and saves it in DB
func CreateRefreshToken(user models.User, tokenString string, expiresAt time.Time) (models.RefreshToken, error) {
	rt := models.RefreshToken{
		TokenString: tokenString,
		ExpiresAt:   expiresAt,
	}

	if err := SaveRefreshTokenForUser(user, rt); err != nil {
		return models.RefreshToken{}, nil
	}
	return rt, nil
}

// GetRefreshTokenByUserID retrieves the refresh token by given user ID
func GetRefreshTokenByUserID(userID uint) (models.RefreshToken, error) {
	var rt models.RefreshToken
	result := database.GetPostgresDB().Where("userID = ?", userID).First(&rt)
	if result.Error != nil {
		return models.RefreshToken{}, nil
	}
	return rt, nil
}
