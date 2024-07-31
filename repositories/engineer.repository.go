package repositories

import (
	"log"

	"github.com/AnkitNayan83/EngineerMandi-Backend/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EngineerRepository interface {
	CreateEngineer(engineerData *models.EngineerModel, userId uuid.UUID) (*models.EngineerModel, error)
	UpdateEngineer(engineerData *models.EngineerModel) (*models.EngineerModel, error)
	GetEngineerByID(id uuid.UUID) (*models.EngineerModel, error)
	CreateEngineerSkill(engineerSkillData *models.EngineerSkills, userId uuid.UUID) (*models.EngineerSkills, error)
	GetEngineerSkillByID(id string) (*models.EngineerSkills, error)
	CreateProject(projectData *models.Project, engineerId uuid.UUID) (*models.Project, error)
	GetProjectByID(id string) (*models.Project, error)
	CreateSkill(skillData *models.Skill) (*models.Skill, error)
	GetSkillByID(id string) (*models.Skill, error)
	CreateSpecialization(specializationData *models.Specialization) (*models.Specialization, error)
	GetSpecializationByID(id string) (*models.Specialization, error)
	CreateEducation(educationData *models.Education, engineerId uuid.UUID) (*models.Education, error)
	GetEducationByID(id string) (*models.Education, error)
	CreateCertification(certificationData *models.Certification, engineerId uuid.UUID) (*models.Certification, error)
	GetCertificationByID(id string) (*models.Certification, error)
	CreateEngineerExperience(engineerExperienceData *models.EngineerExperience, engineerId uuid.UUID) (*models.EngineerExperience, error)
	GetEngineerExperienceByID(id string) (*models.EngineerExperience, error)
}

type engineerRepository struct {
	DB *gorm.DB
}

func NewEngineerRepository(db *gorm.DB) EngineerRepository {
	return &engineerRepository{DB: db}
}

func (r *engineerRepository) CreateEngineer(engineerData *models.EngineerModel, userId uuid.UUID) (*models.EngineerModel, error) {

	engineer := models.EngineerModel{
		UserId: userId,
	}

	resp := r.DB.Create(&engineer)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &engineer, nil
}

func (r *engineerRepository) UpdateEngineer(engineerData *models.EngineerModel) (*models.EngineerModel, error) {

	err := r.DB.Model(&models.EngineerModel{}).Where("user_id = ?", engineerData.UserId).Updates(&engineerData).Error

	if err != nil {
		return nil, err
	}

	return engineerData, nil
}

func (r *engineerRepository) GetEngineerByID(id uuid.UUID) (*models.EngineerModel, error) {

	var engineer models.EngineerModel

	resp := r.DB.Where("user_id = ?", id).First(&engineer)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &engineer, nil
}

func (r *engineerRepository) CreateEngineerSkill(engineerSkillData *models.EngineerSkills, userId uuid.UUID) (*models.EngineerSkills, error) {

	log.Println(engineerSkillData.SkillID)
	log.Println(userId)
	engineerSkill := models.EngineerSkills{
		EngineerID:        userId,
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

func (r *engineerRepository) GetEngineerSkillByID(id string) (*models.EngineerSkills, error) {

	var engineerSkill models.EngineerSkills

	resp := r.DB.Where("id = ?", id).First(&engineerSkill)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &engineerSkill, nil
}

func (r *engineerRepository) CreateProject(projectData *models.Project, engineerId uuid.UUID) (*models.Project, error) {
	// Start a transaction
	tx := r.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	project := models.Project{
		Name:        projectData.Name,
		Description: projectData.Description,
		EngineerID:  engineerId,
	}

	// Create the project
	if err := tx.Create(&project).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if len(projectData.ProjectUrls) > 0 {
		var projectUrls []models.ProjectUrl

		// Create associated project URLs
		for _, projectUrlData := range projectData.ProjectUrls {
			projectUrl := models.ProjectUrl{
				ProjectID: project.ID,
				Url:       projectUrlData.Url,
				Type:      projectUrlData.Type,
			}

			if err := tx.Create(&projectUrl).Error; err != nil {
				tx.Rollback()
				return nil, err
			}

			projectUrls = append(projectUrls, projectUrl)
		}

		project.ProjectUrls = projectUrls
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &project, nil
}

func (r *engineerRepository) GetProjectByID(id string) (*models.Project, error) {

	var project models.Project

	resp := r.DB.Where("id = ?", id).First(&project)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &project, nil
}

func (r *engineerRepository) CreateSkill(skillData *models.Skill) (*models.Skill, error) {

	skill := models.Skill{
		Name:      skillData.Name,
		ShortName: skillData.ShortName,
	}

	resp := r.DB.Create(&skill)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &skill, nil
}

func (r *engineerRepository) GetSkillByID(id string) (*models.Skill, error) {

	var skill models.Skill

	resp := r.DB.Where("id = ?", id).First(&skill)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &skill, nil
}

func (r *engineerRepository) CreateSpecialization(specializationData *models.Specialization) (*models.Specialization, error) {

	specialization := models.Specialization{
		Title: specializationData.Title,
	}

	resp := r.DB.Create(&specialization)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &specialization, nil
}

func (r *engineerRepository) GetSpecializationByID(id string) (*models.Specialization, error) {

	var specialization models.Specialization

	resp := r.DB.Where("id = ?", id).First(&specialization)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &specialization, nil
}

func (r *engineerRepository) CreateEducation(educationData *models.Education, engineerId uuid.UUID) (*models.Education, error) {

	education := models.Education{
		Degree:        educationData.Degree,
		Institute:     educationData.Institute,
		YearOfPassing: educationData.YearOfPassing,
		CGPA:          educationData.CGPA,
		EngineerID:    engineerId,
	}

	resp := r.DB.Create(&education)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &education, nil
}

func (r *engineerRepository) GetEducationByID(id string) (*models.Education, error) {

	var education models.Education

	resp := r.DB.Where("id = ?", id).First(&education)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &education, nil
}

func (r *engineerRepository) CreateCertification(certificationData *models.Certification, engineerId uuid.UUID) (*models.Certification, error) {

	certification := models.Certification{
		Name:           certificationData.Name,
		Authority:      certificationData.Authority,
		CertificateUrl: certificationData.CertificateUrl,
		Description:    certificationData.Description,
		IssuedDate:     certificationData.IssuedDate,
		EngineerID:     engineerId,
	}

	resp := r.DB.Create(&certification)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &certification, nil
}

func (r *engineerRepository) GetCertificationByID(id string) (*models.Certification, error) {

	var certification models.Certification

	resp := r.DB.Where("id = ?", id).First(&certification)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &certification, nil
}

func (r *engineerRepository) CreateEngineerExperience(engineerExperienceData *models.EngineerExperience, engineerId uuid.UUID) (*models.EngineerExperience, error) {

	engineerExperience := models.EngineerExperience{
		Company:     engineerExperienceData.Company,
		Location:    engineerExperienceData.Location,
		Role:        engineerExperienceData.Role,
		Description: engineerExperienceData.Description,
		StartDate:   engineerExperienceData.StartDate,
		IsCurrent:   engineerExperienceData.IsCurrent,
		EndDate:     engineerExperienceData.EndDate,
		EngineerID:  engineerId,
	}

	resp := r.DB.Create(&engineerExperience)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &engineerExperience, nil
}

func (r *engineerRepository) GetEngineerExperienceByID(id string) (*models.EngineerExperience, error) {

	var engineerExperience models.EngineerExperience

	resp := r.DB.Where("id = ?", id).First(&engineerExperience)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &engineerExperience, nil
}
