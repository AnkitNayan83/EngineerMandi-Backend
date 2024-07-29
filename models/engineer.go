package models

import (
	"time"

	"github.com/google/uuid"
)

type EngineerModel struct {
	UserId          uuid.UUID        `gorm:"type:uuid;primaryKey" json:"userId"`
	User            User             `gorm:"foreignKey:UserId" json:"user"`
	Resume          string           `gorm:"not null" json:"resume"`
	Specializations []Specialization `gorm:"many2many:engineer_specializations" json:"specializations"`
	Experience      float64          `json:"experience"`
	Skills          []Skill          `gorm:"many2many:engineer_skills" json:"skills"`
	Education       []Education      `gorm:"foreignKey:EngineerID;references:UserId" json:"education"`
	Certifications  []Certification  `gorm:"foreignKey:EngineerID;references:UserId" json:"certifications"`
	Projects        []Project        `gorm:"foreignKey:EngineerID;references:UserId" json:"projects"`
}

type ProficiencyLevelEnum string

const (
	Beginner     ProficiencyLevelEnum = "beginner"
	Intermediate ProficiencyLevelEnum = "intermediate"
	Advanced     ProficiencyLevelEnum = "advanced"
)

type EngineerSkills struct {
	EngineerID        uuid.UUID            `gorm:"type:uuid;primaryKey" json:"engineerId"`
	SkillID           uuid.UUID            `gorm:"type:uuid;primaryKey" json:"skillId"`
	ProficiencyLevel  ProficiencyLevelEnum `gorm:"type:VARCHAR(20)" json:"proficiencyLevel"`
	YearsOfExperience float64              `gorm:"default:0.0" json:"yearsOfExperience"`
}

type Skill struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	ShortName string    `gorm:"not null" json:"shortName"`
}

type Specialization struct {
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title string    `gorm:"not null" json:"title"`
}

type Education struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	EngineerID uuid.UUID `gorm:"type:uuid;not null" json:"engineerId"`
	Degree     string    `gorm:"not null" json:"degree"`
	Institute  string    `gorm:"not null" json:"institute"`
	CGPA       float64   `gorm:"not null" json:"cgpa"`
}

type Certification struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Authority   string    `gorm:"not null" json:"authority"`
	Description string    `json:"description,omitempty"`
	IssuedDate  time.Time `gorm:"not null" json:"issuedDate"`
	EngineerID  uuid.UUID `gorm:"type:uuid;not null" json:"engineerId"`
}

type Project struct {
	ID          uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string       `gorm:"not null" json:"name"`
	Description string       `json:"description,omitempty"`
	ProjectUrls []ProjectUrl `gorm:"foreignKey:ProjectID;references:ID" json:"projectUrls"`
	EngineerID  uuid.UUID    `gorm:"type:uuid;not null" json:"engineerId"`
}

type ProjectUrl struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	ProjectID uuid.UUID `gorm:"type:uuid;not null;index" json:"projectId"`
	Url       string    `gorm:"not null;uniqueIndex:idx_projectid_url" json:"url"`
	Type      string    `gorm:"not null;uniqueIndex:idx_projectid_type" json:"type"`
}
