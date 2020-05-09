package repositories

import (
	"fmt"

	"github.com/IstvanN/cashcalc-backend/properties"

	"github.com/IstvanN/cashcalc-backend/database"
	"github.com/IstvanN/cashcalc-backend/models"
)

//GetUsers retrieves all users from the database
func GetUsers() (models.Users, error) {
	coll := database.GetCollectionByName(properties.UsersCollection)

	var u models.Users
	err := coll.Find(nil).One(&u)
	if err != nil {
		return models.Users{}, fmt.Errorf("error while retrieving collection %v from database: %v",
			"users", err)
	}

	return u, nil
}

// GetUserByRole retrieves the user by its role
func GetUserByRole(role models.Role) (models.User, error) {
	users, err := GetUsers()
	if err != nil {
		return models.User{}, err
	}

	for _, u := range users.Users {
		if u.Role == role {
			return u, nil
		}
	}
	return models.User{}, nil
}
