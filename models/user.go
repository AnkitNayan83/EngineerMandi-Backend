package models

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	Admin    Role = "admin"
	Client   Role = "client"
	Engineer Role = "engineer"
)

type User struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	FirstName      string    `gorm:"not null" json:"firstName"`
	LastName       string    `json:"lastName"`
	Email          string    `gorm:"unique;not null" json:"email"`
	Bio            string    `json:"bio"`
	ProfilePicture string    `json:"profilePicture"`
	Address        string    `json:"address"`
	PinCode        string    `json:"pinCode"`
	City           string    `json:"city"`
	State          string    `json:"state"`
	Country        string    `json:"country"`
	Role           Role      `gorm:"type:VARCHAR(20);default:client" json:"role"`
	Conversations  []*Conversation
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type OAuthUser struct {
	ID        string `json:"id"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	Email     string `json:"email"`
	Picture   string `json:"picture"`
}

type Conversation struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	IsOver       bool       `gorm:"default:false" json:"isOver"`
	Participants []*User    `gorm:"many2many:conversation_participants;" json:"participants"`
	Messages     []*Message `gorm:"foreignKey:ConversationID" json:"messages"`
}

type Message struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	ConversationID  uuid.UUID `json:"conversationID"`
	SenderID        uuid.UUID `json:"senderID"`
	Content         string    `json:"content"`
	SentAt          time.Time `json:"sentAt"`
	IsEdited        bool      `gorm:"default:true" json:"isEdited"`
	IsDeleted       bool      `gorm:"default:true" json:"isDeleted"`
	Attachments     []string  `json:"attachments,omitempty"`
	AttachmentsType string    `json:"attachmentsType,omitempty"`
}
