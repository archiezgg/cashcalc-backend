package models

// User is the main struct for users such as carrier, admin and superuser
type User struct {
	ID       int    `bson:"id" json:"id"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	Role     Role   `bson:"role" json:"role"`
}

// Role is an enum that can either be carrier, admin or superuser
type Role string

const (
	// Carrier is the basic role, with minimum privileges
	Carrier = "carrier"
	// Admin has privileges to set pricing variables
	Admin = "admin"
	// Superuser has privileges to modify database, revoke tokens and such
	Superuser = "superuser"
)
