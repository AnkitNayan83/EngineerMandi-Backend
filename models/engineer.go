package models

import (
	"time"

	"github.com/google/uuid"
)

type EngineerModel struct {
	UserId          uuid.UUID            `gorm:"type:uuid;primaryKey" json:"userId"`
	User            User                 `gorm:"foreignKey:UserId" json:"user"`
	Resume          string               `json:"resume"`
	Specializations []Specialization     `gorm:"many2many:engineer_specializations" json:"specializations"`
	Experiences     []EngineerExperience `gorm:"foreignKey:EngineerID;references:UserId;constraint:OnDelete:CASCADE;" json:"experiences"`
	Skills          []EngineerSkills     `gorm:"foreignKey:EngineerID;references:UserId;constraint:OnDelete:CASCADE;" json:"skills"`
	Education       []Education          `gorm:"foreignKey:EngineerID;references:UserId;constraint:OnDelete:CASCADE;" json:"education"`
	Certifications  []Certification      `gorm:"foreignKey:EngineerID;references:UserId;constraint:OnDelete:CASCADE;" json:"certifications"`
	Projects        []Project            `gorm:"foreignKey:EngineerID;references:UserId;constraint:OnDelete:CASCADE;" json:"projects"`
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
	Skill             Skill                `gorm:"foreignKey:SkillID;references:ID" json:"skill"`
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
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	EngineerID    uuid.UUID `gorm:"type:uuid;not null" json:"engineerId"`
	Degree        string    `gorm:"not null" json:"degree"`
	Branch        string    `gorm:"not null" json:"branch"`
	Institute     string    `gorm:"not null" json:"institute"`
	YearOfPassing int       `gorm:"not null" json:"yearOfPassing"`
	CGPA          float64   `gorm:"not null" json:"cgpa"`
}

type Certification struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name           string    `gorm:"not null" json:"name"`
	Authority      string    `gorm:"not null" json:"authority"`
	CertificateUrl string    `json:"certificateUrl"`
	Description    string    `json:"description,omitempty"`
	IssuedDate     time.Time `gorm:"not null" json:"issuedDate"`
	EngineerID     uuid.UUID `gorm:"type:uuid;not null" json:"engineerId"`
}

type Project struct {
	ID          uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string       `gorm:"not null" json:"name"`
	Description string       `json:"description,omitempty"`
	ProjectUrls []ProjectUrl `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE;references:ID" json:"projectUrls"`
	EngineerID  uuid.UUID    `gorm:"type:uuid;not null" json:"engineerId"`
}

type ProjectUrl struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	ProjectID uuid.UUID `gorm:"type:uuid;not null;index;constraint:OnDelete:CASCADE;" json:"projectId"`
	Url       string    `gorm:"not null" json:"url"`
	Type      string    `gorm:"not null" json:"type"`
}

type EngineerExperience struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	EngineerID  uuid.UUID `gorm:"type:uuid;not null;index" json:"engineerId"`
	Company     string    `gorm:"not null" json:"company"`
	Location    string    `gorm:"not null" json:"location"`
	Role        string    `gorm:"not null" json:"role"`
	Description string    `json:"description,omitempty"`
	StartDate   time.Time `gorm:"not null" json:"startDate"`
	IsCurrent   bool      `gorm:"not null" json:"isCurrent"`
	EndDate     time.Time `json:"endDate,omitempty"`
}
