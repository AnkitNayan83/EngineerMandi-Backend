package services

import (
	"github.com/AnkitNayan83/EngineerMandi-Backend/models"
	"github.com/AnkitNayan83/EngineerMandi-Backend/repositories"
)

type UserService interface {
	ProfileSetup(userId string, user *models.User) (*models.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) ProfileSetup(userId string, updatedUser *models.User) (*models.User, error) {

	currentUser, err := s.userRepository.FindUserById(userId)

	if err != nil {
		return nil, err
	}

	if updatedUser.FirstName != "" {
		currentUser.FirstName = updatedUser.FirstName
	}
	if updatedUser.LastName != "" {
		currentUser.LastName = updatedUser.LastName
	}
	if updatedUser.Bio != "" {
		currentUser.Bio = updatedUser.Bio
	}
	if updatedUser.ProfilePicture != "" {
		currentUser.ProfilePicture = updatedUser.ProfilePicture
	}
	if updatedUser.Address != "" {
		currentUser.Address = updatedUser.Address
	}
	if updatedUser.PinCode != "" {
		currentUser.PinCode = updatedUser.PinCode
	}
	if updatedUser.City != "" {
		currentUser.City = updatedUser.City
	}
	if updatedUser.State != "" {
		currentUser.State = updatedUser.State
	}
	if updatedUser.Country != "" {
		currentUser.Country = updatedUser.Country
	}

	if err := s.userRepository.UpdateUserById(userId, currentUser); err != nil {
		return nil, err
	}

	return currentUser, nil

}
