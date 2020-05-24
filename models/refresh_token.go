/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package models

// RefreshToken represents a refresh token structure
type RefreshToken struct {
	Username    string `json:"username"`
	Role        Role   `json:"role"`
	TokenString string `json:"tokenString"`
	ExpiresAt   int64  `json:"expiresAt"`
}
