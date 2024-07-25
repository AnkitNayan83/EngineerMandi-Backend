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
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FirstName      string    `gorm:"not null"`
	LastName       string
	Email          string `gorm:"unique;not null"`
	Bio            string
	ProfilePicture string
	Address        string
	PinCode        string
	City           string
	State          string
	Country        string
	Role           Role `gorm:"type:VARCHAR(20);not null"`
	Ratings        []Rating
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Rating struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index"`
	Stars     int       `gorm:"not null"`
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
