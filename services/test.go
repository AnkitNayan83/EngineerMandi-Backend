package services

import (
	"errors"

	"github.com/AnkitNayan83/SMA-backend/models"
)

var users = []models.User{
	{ID: "1", Name: "John Doe"},
	{ID: "2", Name: "Michel Scholfield"},
}

func GetAllUsers() ([]models.User,error) {
	if(len(users) == 0) {
		return nil, errors.New("no users found");
	}

	return users, nil

}


func CreateUser(newUser models.User) (models.User, error) {

	if(newUser.Name == "") {
		return models.User{}, errors.New("name cannot be empty")
	}

	for _, user := range users {
		if user.ID == newUser.ID {
			return models.User{}, errors.New("user already exists")
		}
	}

	users = append(users, newUser)

	return newUser, nil
}