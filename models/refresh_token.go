/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package models

// RefreshToken stores the structure of a refresh token
type RefreshToken struct {
	Username    string `json:"username"`
	TokenString string `json:"tokenString"`
	ExpiresAt   int64  `json:"expiresAt"`
}
