/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package repositories

import (
	"fmt"

	"github.com/IstvanN/cashcalc-backend/database"
	"github.com/IstvanN/cashcalc-backend/models"
	"github.com/IstvanN/cashcalc-backend/properties"
	"golang.org/x/crypto/bcrypt"
)

// GetUserByID retrieves the user by id
func GetUserByID(id uint) (models.User, error) {
	var u models.User
	result := database.GetPostgresDB().First(&u, id)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return u, nil
}

// GetUserByUsername retrieves the user by its username
func GetUserByUsername(username string) (models.User, error) {
	var u models.User
	result := database.GetPostgresDB().Where("username = ?", username).First(&u)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return u, nil
}

// GetAllUsers retrieves all users
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := database.GetPostgresDB().Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// GetAllUsernames retrieves all usernames
func GetAllUsernames() ([]string, error) {
	users, err := GetAllUsers()
	if err != nil {
		return nil, err
	}

	var usernames []string
	for _, u := range users {
		usernames = append(usernames, u.Username)
	}
	return usernames, nil
}

// GetUsersByRole retrieves all users by given role
func GetUsersByRole(role models.Role) ([]models.User, error) {
	var users []models.User
	result := database.GetPostgresDB().Where("role = ?", role).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// CreateUser creates user based on a username, password and a role
func CreateUser(username, password string, role models.Role) error {
	if err := checkIfUsernameIsTaken(username); err != nil {
		return err
	}

	if err := checkIfUsernameAndPasswordFulfillsRequirements(username, password); err != nil {
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

	result := database.GetPostgresDB().Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// SaveUser updates the user in DB with the given user object
func SaveUser(user models.User) error {
	result := database.GetPostgresDB().Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteUserByID checks if user with given ID is in DB and deletes it
// and all the refresh tokens of the user
func DeleteUserByID(id uint) error {
	_, err := GetUserByID(id)
	if err != nil {
		return err
	}

	if err := DeleteAllRefreshTokensForUser(id); err != nil {
		return err
	}

	result := database.GetPostgresDB().Unscoped().Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteUserByIDAndRole deletes the user by given ID and role
// returns error if the user by id is not matched for given role
func DeleteUserByIDAndRole(id uint, role models.Role) error {
	user, err := GetUserByID(id)
	if err != nil {
		return err
	}

	if user.Role != role {
		return fmt.Errorf("user with id: %v has no role: %v", id, role)
	}

	if err := DeleteUserByID(id); err != nil {
		return err
	}

	return nil
}

func checkIfUsernameIsTaken(username string) error {
	users, err := GetAllUsers()
	if err != nil {
		return err
	}

	for _, u := range users {
		if u.Username == username {
			return fmt.Errorf("username already taken: %v", username)
		}
	}
	return nil
}

// CreateUserDTOFromUser creates a user DTO that is sent back via endpoints
func CreateUserDTOFromUser(user models.User) models.UserDTO {
	return models.UserDTO{
		ID:        user.ID,
		Username:  user.Username,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt.Time,
	}
}

// CreateUserDTOsFromUsers takes a slice of users
// and returns with a slice with userDTOs
func CreateUserDTOsFromUsers(users []models.User) []models.UserDTO {
	var userDTOs []models.UserDTO

	for _, u := range users {
		userDTOs = append(userDTOs, CreateUserDTOFromUser(u))
	}

	return userDTOs
}

// SaveRefreshTokenForUser saves the refresh token for the user in the DB
func SaveRefreshTokenForUser(user models.User, rt models.RefreshToken) error {
	user.RefreshTokens = append(user.RefreshTokens, rt)
	if err := SaveUser(user); err != nil {
		return err
	}
	return nil
}

// GetAllLoggedInUsers retrieves all the users if they have refresh tokens stored in DB
func GetAllLoggedInUsers() ([]models.User, error) {
	refreshTokens, err := GetAllRefreshTokens()
	if err != nil {
		return nil, err
	}

	var loggedInUsersIDs []uint
	for _, rt := range refreshTokens {
		loggedInUsersIDs = append(loggedInUsersIDs, rt.UserID)
	}

	var loggedInUsers []models.User
	result := database.GetPostgresDB().Where("id IN ?", loggedInUsersIDs).Find(&loggedInUsers)
	if result.Error != nil {
		return nil, result.Error
	}
	return loggedInUsers, nil
}

func checkIfUsernameAndPasswordFulfillsRequirements(username, pw string) error {
	if len(username) < properties.UserUsernameMinLength || len(username) > properties.UserUsernameMaxLength {
		return fmt.Errorf("username must be between %v and %v characters", properties.UserUsernameMinLength, properties.UserUsernameMaxLength)
	}

	if len(pw) < properties.UserPasswordMinLength || len(pw) > properties.UserPasswordMaxLength {
		return fmt.Errorf("password must be between %v and %v characters", properties.UserPasswordMinLength, properties.UserPasswordMaxLength)
	}
	return nil
}
