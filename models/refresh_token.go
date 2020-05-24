/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package models

// Token represents a refresh token structure
type Token struct {
	Username    string `json:"username"`
	Role        Role   `json:"role"`
	TokenString string `json:"tokenString"`
}
