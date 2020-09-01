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

// GetUserDTOsByRole retrieves all user DTO by given role
func GetUserDTOsByRole(role models.Role) ([]models.UserDTO, error) {
	users, err := GetUsersByRole(role)
	if err != nil {
		return nil, err
	}

	var userDTOs []models.UserDTO
	for _, u := range users {
		userDTOs = append(userDTOs, CreateUserDTOFromUser(u))
	}
	return userDTOs, nil
}

// CreateUser creates user based on a username, password and a role
func CreateUser(username, password string, role models.Role) error {
	if err := checkIfUsernameIsTaken(username); err != nil {
		return err
	}

	if len(password) < properties.UserPasswordMinLength {
		return fmt.Errorf("password must be at least %v characters", properties.UserPasswordMinLength)
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

func SaveUser(user models.User) error {
	result := database.GetPostgresDB().Save(&user)
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

	result := database.GetPostgresDB().Where("role = ?", role).Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
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

// SaveRefreshTokenForUser saves the refresh token for the user in the DB
func SaveRefreshTokenForUser(user models.User, rt models.RefreshToken) error {
	user.RefreshTokens = append(user.RefreshTokens, rt)
	if err := SaveUser(user); err != nil {
		return err
	}
	return nil
}
