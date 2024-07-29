package repositories

import (
	"github.com/AnkitNayan83/EngineerMandi-Backend/models"
	"gorm.io/gorm"
)

type SpecializationRepository interface {
	CreateSpecialization(specializationData *models.Specialization) (*models.Specialization, error)
	GetSpecializationByID(id string) (*models.Specialization, error)
}

type specializationRepository struct {
	DB *gorm.DB
}

func NewSpecializationRepository(db *gorm.DB) SpecializationRepository {
	return &specializationRepository{DB: db}
}

func (r *specializationRepository) CreateSpecialization(specializationData *models.Specialization) (*models.Specialization, error) {

	specialization := models.Specialization{
		Title: specializationData.Title,
	}

	resp := r.DB.Create(&specialization)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &specialization, nil
}

func (r *specializationRepository) GetSpecializationByID(id string) (*models.Specialization, error) {

	var specialization models.Specialization

	resp := r.DB.Where("id = ?", id).First(&specialization)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &specialization, nil
}
