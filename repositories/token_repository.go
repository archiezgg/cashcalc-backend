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

// SaveRefreshToken saves the token with the role to the DB
func SaveRefreshToken(rt string, role models.Role) {
	database.RedisClient().Set(rt, role, properties.RefreshTokenExp)
}
