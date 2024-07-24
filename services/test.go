package services

import (
	"errors"

	"github.com/AnkitNayan83/SMA-backend/models"
)

var users = []models.UserTest{
	{ID: "1", Name: "John Doe"},
	{ID: "2", Name: "Michel Scholfield"},
}

func GetAllUsers() ([]models.UserTest, error) {
	if len(users) == 0 {
		return nil, errors.New("no users found")
	}

	return users, nil

}

func CreateUser(newUser models.UserTest) (models.UserTest, error) {

	if newUser.Name == "" {
		return models.UserTest{}, errors.New("name cannot be empty")
	}

	for _, user := range users {
		if user.ID == newUser.ID {
			return models.UserTest{}, errors.New("user already exists")
		}
	}

	users = append(users, newUser)

	return newUser, nil
}
