package repositories

import (
	"github.com/AnkitNayan83/EngineerMandi-Backend/models"
	"gorm.io/gorm"
)

type EngineerSkillRepository interface {
	CreateEngineerSkill(engineerSkillData *models.EngineerSkills) (*models.EngineerSkills, error)
	GetEngineerSkillByID(id string) (*models.EngineerSkills, error)
}

type engineerSkillRepository struct {
	DB *gorm.DB
}

func NewEngineerSkillRepository(db *gorm.DB) EngineerSkillRepository {
	return &engineerSkillRepository{DB: db}
}

func (r *engineerSkillRepository) CreateEngineerSkill(engineerSkillData *models.EngineerSkills) (*models.EngineerSkills, error) {

	engineerSkill := models.EngineerSkills{
		EngineerID:        engineerSkillData.EngineerID,
		SkillID:           engineerSkillData.SkillID,
		ProficiencyLevel:  engineerSkillData.ProficiencyLevel,
		YearsOfExperience: engineerSkillData.YearsOfExperience,
	}

	resp := r.DB.Create(&engineerSkill)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &engineerSkill, nil
}

func (r *engineerSkillRepository) GetEngineerSkillByID(id string) (*models.EngineerSkills, error) {

	var engineerSkill models.EngineerSkills

	resp := r.DB.Where("id = ?", id).First(&engineerSkill)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &engineerSkill, nil
}
