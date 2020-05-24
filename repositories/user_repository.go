/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

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
		errMsg := fmt.Errorf("error while retrieving collection %v from database: %v",
			properties.UsersCollection, err)
		return models.Users{}, errMsg
	}

	return u, nil
}

// GetUserByUsername retrieves the user by its username
func GetUserByUsername(username string) (models.User, error) {
	users, err := GetUsers()
	if err != nil {
		return models.User{}, err
	}

	for _, u := range users.Users {
		if u.Username == username {
			return u, nil
		}
	}
	err = fmt.Errorf("user cannot be found in db by username: %v", username)
	return models.User{}, err
}
