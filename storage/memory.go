package storage

import (
	"errors"
	"openim/models"
)

var Users = map[string]models.User{}

func SaveUser(user models.User) {
	Users[user.Email] = user
}

func GetUser(email string) (models.User, error) {
	user, exists := Users[email]
	if !exists {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func UserExists(email string) bool {
	_, exists := Users[email]
	return exists
}
