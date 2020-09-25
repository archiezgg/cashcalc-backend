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

// CreateRefreshTokenAndSaveForUser creates a refresh token struct from data
// and saves it in DB
func CreateRefreshTokenAndSaveForUser(user models.User, tokenString string, expiresAt time.Time) (models.RefreshToken, error) {
	rt := models.RefreshToken{
		TokenString: tokenString,
		CreatedAt:   time.Now(),
		ExpiresAt:   expiresAt,
	}

	if err := SaveRefreshTokenForUser(user, rt); err != nil {
		return models.RefreshToken{}, nil
	}
	return rt, nil
}

// GetRefreshTokenByTokenString retrieves the rt by ID (token string)
func GetRefreshTokenByTokenString(tokenString string) (models.RefreshToken, error) {
	var rt models.RefreshToken
	result := database.GetPostgresDB().First(&rt)
	if result.Error != nil {
		return models.RefreshToken{}, result.Error
	}
	return rt, nil
}

// GetRefreshTokensByUserID retrieves the refresh token by given user ID
func GetRefreshTokensByUserID(userID uint) ([]models.RefreshToken, error) {
	var refreshTokens []models.RefreshToken
	result := database.GetPostgresDB().Where("user_id = ?", userID).Find(&refreshTokens)
	if result.Error != nil {
		return nil, result.Error
	}
	return refreshTokens, nil
}

// GetAllRefreshTokens retrieves all refresh tokens stored in DB
func GetAllRefreshTokens() ([]models.RefreshToken, error) {
	var refreshTokens []models.RefreshToken
	result := database.GetPostgresDB().Find(&refreshTokens)
	if result.Error != nil {
		return nil, result.Error
	}
	return refreshTokens, nil
}

// DeleteRefreshTokenByTokenString deletes the refresh token by ID (token string)
func DeleteRefreshTokenByTokenString(tokenString string) error {
	rt, err := GetRefreshTokenByTokenString(tokenString)
	if err != nil {
		return err
	}

	result := database.GetPostgresDB().Delete(rt)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteAllRefreshTokensForUser deletes all refresh tokens for given user
func DeleteAllRefreshTokensForUser(userID uint) error {
	result := database.GetPostgresDB().Where("user_id = ?", userID).Delete(&models.RefreshToken{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteExpiredRefreshTokens delete all expired refresh tokens
func DeleteExpiredRefreshTokens() error {
	result := database.GetPostgresDB().Where("expires_at > ?", time.Now()).Delete(&models.RefreshToken{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
