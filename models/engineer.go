package models

import (
	"time"

	"github.com/google/uuid"
)

type EngineerModel struct {
	UserId          uuid.UUID        `gorm:"type:uuid;primaryKey"`
	User            User             `gorm:"foreignKey:UserId"`
	Specializations []Specialization `gorm:"many2many:engineer_specializations"`
	Experience      float64
	Skills          []Skill `gorm:"many2many:engineer_skills"`
	Education       []Education
	Certifications  []Certification
	Projects        []Project
}

type ProficiencyLevelEnum string

const (
	Beginner     ProficiencyLevelEnum = "beginner"
	Intermediate ProficiencyLevelEnum = "intermediate"
	Advanced     ProficiencyLevelEnum = "advanced"
)

type EngineerSkills struct {
	EngineerID        uuid.UUID            `gorm:"type:uuid;primaryKey"`
	SkillID           uuid.UUID            `gorm:"type:uuid;primaryKey"`
	ProficiencyLevel  ProficiencyLevelEnum `gorm:"type:VARCHAR(20)"`
	YearsOfExperience float64              `gorm:"default:0.0"`
}

type Skill struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string    `gorm:"not null"`
	ShortName string    `gorm:"not null"`
}

type Specialization struct {
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title string    `gorm:"not null"`
}

type Education struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	EngineerID uuid.UUID `gorm:"type:uuid;not null"`
	Degree     string    `gorm:"not null"`
	Institute  string    `gorm:"not null"`
	CGPA       float64   `gorm:"not null"`
}

type Certification struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string    `gorm:"not null"`
	Authority   string    `gorm:"not null"`
	Description string
	IssuedDate  time.Time `gorm:"not null"`
	EngineerID  uuid.UUID `gorm:"type:uuid;not null"`
}

type Project struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string    `gorm:"not null"`
	Description string
	Urls        []string
	EngineerID  uuid.UUID `gorm:"type:uuid;not null"`
}
