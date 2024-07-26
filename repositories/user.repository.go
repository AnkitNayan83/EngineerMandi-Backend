package repositories

import (
	"errors"
	"strings"

	"github.com/AnkitNayan83/SMA-backend/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.OAuthUser) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	FindUserById(id string) (*models.User, error)
	UpdateUserById(id string, user *models.User) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) CreateUser(user *models.OAuthUser) (*models.User, error) {
	names := strings.Split(user.Name, " ")
	if len(names) == 0 {
		return nil, errors.New("invalid user name")
	}

	var firstName string
	var lastName string

	firstName = names[0]
	if len(names) > 1 {
		lastName = names[1]
	}

	userData := models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     user.Email,
	}
	resp := r.DB.Create(&userData)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &userData, nil
}

func (r *userRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindUserById(id string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUserById(id string, user *models.User) error {
	err := r.DB.Model(&models.User{}).Where("id = ?", id).Updates(&user).Error
	if err != nil {
		return err
	}
	return nil
}
