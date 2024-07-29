package repositories

import (
	"github.com/AnkitNayan83/EngineerMandi-Backend/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EngineerRepository interface {
}

type engineerRepository struct {
	DB *gorm.DB
}

func NewEngineerRepository(db *gorm.DB) EngineerRepository {
	return &engineerRepository{DB: db}
}

func (r *engineerRepository) CreateEngineer(engineerData *models.EngineerModel, userId uuid.UUID) (*models.EngineerModel, error) {

	engineer := models.EngineerModel{
		UserId:          userId,
		Specializations: engineerData.Specializations,
		Experience:      engineerData.Experience,
		Skills:          engineerData.Skills,
		Education:       engineerData.Education,
		Certifications:  engineerData.Certifications,
		Projects:        engineerData.Projects,
	}

	resp := r.DB.Create(&engineer)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &engineer, nil
}
