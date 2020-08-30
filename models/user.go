/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package models

import "gorm.io/gorm"

// User is the main struct for users such as carrier, admin and superuser
type User struct {
	gorm.Model
	Username string `gorm:"string;not null;unique"`
	Password string `gorm:"string;not null"`
	Role     Role   `gorm:"string;not null"`
}

// Role is an enum that can either be carrier, admin or superuser
type Role string

const (
	// RoleCarrier is the basic role, with minimum privileges
	RoleCarrier = "carrier"
	// RoleAdmin has privileges to set pricing variables
	RoleAdmin = "admin"
	// RoleSuperuser has privileges to modify database, revoke tokens and such
	RoleSuperuser = "superuser"
)
