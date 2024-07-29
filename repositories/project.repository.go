package repositories

import (
	"github.com/AnkitNayan83/EngineerMandi-Backend/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	CreateProject(projectData *models.Project, id uuid.UUID) (*models.Project, error)
	GetProjectByID(id string) (*models.Project, error)
}

type projectRepository struct {
	DB *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{DB: db}
}

func (r *projectRepository) CreateProject(projectData *models.Project, engineerId uuid.UUID) (*models.Project, error) {

	project := models.Project{
		Name:        projectData.Name,
		Description: projectData.Description,
		ProjectUrls: projectData.ProjectUrls,
		EngineerID:  engineerId,
	}

	resp := r.DB.Create(&project)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &project, nil
}

func (r *projectRepository) GetProjectByID(id string) (*models.Project, error) {

	var project models.Project

	resp := r.DB.Where("id = ?", id).First(&project)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &project, nil
}
