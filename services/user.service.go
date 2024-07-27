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

// func (s *userService) ProfileSetup(userId string, user *models.User) (*models.User, error) {

// 	currentUser, err := s.userRepository.FindUserById(userId)

// 	if err != nil {
// 		return nil, err
// 	}

// 	currentUser = s.userRepository.UpdateUserById()

// }
