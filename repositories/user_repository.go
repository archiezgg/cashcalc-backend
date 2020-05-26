/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package repositories

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/IstvanN/cashcalc-backend/properties"

	"github.com/IstvanN/cashcalc-backend/database"
	"github.com/IstvanN/cashcalc-backend/models"
)

// GetUsers retrieves all users from the database
func GetUsers() ([]models.User, error) {
	coll := database.GetCollectionByName(properties.UsersCollection)

	var users []models.User
	err := coll.Find(nil).All(&users)
	if err != nil {
		errMsg := fmt.Errorf("error while retrieving collection %v from database: %v",
			properties.UsersCollection, err)
		return nil, errMsg
	}

	return users, nil
}

// GetUsernames returns all usernames from DB
func GetUsernames() ([]string, error) {
	users, err := GetUsers()
	if err != nil {
		return nil, err
	}

	var usernames []string
	for _, u := range users {
		usernames = append(usernames, u.Username)
	}
	return usernames, nil
}

// GetUserByUsername retrieves the user by its username
func GetUserByUsername(username string) (models.User, error) {
	users, err := GetUsers()
	if err != nil {
		return models.User{}, err
	}

	for _, u := range users {
		if u.Username == username {
			return u, nil
		}
	}
	err = fmt.Errorf("user cannot be found in db by username: %v", username)
	return models.User{}, err
}

// CreateUser creates a user by given username, password and role and saves it to DB
func CreateUser(username, password string, role models.Role) error {
	if err := checkIfUserNameIsTaken(username); err != nil {
		return err
	}

	hashedPw, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	user := models.User{
		Username: username,
		Password: string(hashedPw),
		Role:     role,
	}

	coll := database.GetCollectionByName(properties.UsersCollection)
	if err := coll.Insert(user); err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes given user from the DB
func DeleteUser(user models.User) error {
	coll := database.GetCollectionByName(properties.UsersCollection)
	if err := coll.Remove(user); err != nil {
		return err
	}
	return nil
}

func checkIfUserNameIsTaken(username string) error {
	usernames, err := GetUsernames()
	if err != nil {
		return err
	}

	for _, u := range usernames {
		if username == u {
			return fmt.Errorf("username already taken: %v", username)
		}
	}
	return nil
}
