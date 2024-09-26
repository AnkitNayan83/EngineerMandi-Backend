package repositories

import (
	"github.com/AnkitNayan83/EngineerMandi-Backend/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.OAuthUser) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	FindUserById(id string) (*models.User, error)
	UpdateUserById(id string, user *models.User) error
	CreateEngineer(engineer models.EngineerModel) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) CreateUser(user *models.OAuthUser) (*models.User, error) {

	userData := models.User{
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		ProfilePicture: user.Picture,
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

func (r *userRepository) CreateEngineer(engineer models.EngineerModel) error {
	resp := r.DB.Create(&engineer)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

func (r *userRepository) CreateConversation(conversation *models.Conversation) error {
	resp := r.DB.Create(&conversation)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

func (r *userRepository) FindConversationById(id uuid.UUID) (*models.Conversation, error) {
	var conversation models.Conversation
	err := r.DB.Where("id = ?", id).First(&conversation).Error
	if err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (r *userRepository) EndConversation(id uuid.UUID) error {
	err := r.DB.Model(&models.Conversation{}).Where("id = ?", id).Update("is_over", true).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetConversationMessages(conversationId uuid.UUID) ([]models.Message, error) {
	var messages []models.Message
	err := r.DB.Where("conversation_id = ?", conversationId).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *userRepository) EditMessage(messageId uuid.UUID, senderId uuid.UUID, content string) error {
	err := r.DB.Model(&models.Message{}).Where("id = ? AND sender_id = ?", messageId, senderId).Update("content", content).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) DeleteMessage(messageId uuid.UUID, senderId uuid.UUID) error {
	err := r.DB.Where("id = ? AND sender_id = ?", messageId, senderId).Update(
		"is_deleted = ? content = This message has been deleted", true).Error

	if err != nil {
		return err
	}
	return nil
}
