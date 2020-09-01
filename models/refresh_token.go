/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package models

import (
	"time"
)

// RefreshToken stores the structure of a refresh token
type RefreshToken struct {
	TokenString string `gorm:"primaryKey;not null" json:"tokenString"`
	UserID      uint
	CreatedAt   time.Time
	ExpiresAt   time.Time `json:"expiresAt"`
}
