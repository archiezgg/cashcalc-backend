/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package models

// User is the main struct for users such as carrier, admin and superuser
type User struct {
	Role     Role   `bson:"role" json:"role"`
	Password string `bson:"password" json:"password"`
}

// Users holds all users in Users field
type Users struct {
	Users []User `bson:"users" json:"users"`
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
