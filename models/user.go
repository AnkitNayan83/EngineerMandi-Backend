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
