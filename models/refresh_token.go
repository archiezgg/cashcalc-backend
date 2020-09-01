/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package models

import (
	"time"

	"gorm.io/gorm"
)

// RefreshToken stores the structure of a refresh token
type RefreshToken struct {
	gorm.Model
	UserID      uint
	TokenString string    `json:"tokenString"`
	ExpiresAt   time.Time `json:"expiresAt"`
}
