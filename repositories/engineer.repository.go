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
	GetEngineerSkills(engineerId uuid.UUID) ([]models.EngineerSkills, error)
	UpdateEngineerSkill(engineerSkillData *models.EngineerSkills, userId uuid.UUID) (*models.EngineerSkills, error)
	RemoveEngineerSkill(id uuid.UUID, userId uuid.UUID) error

	CreateProject(projectData *models.Project, engineerId uuid.UUID) (*models.Project, error)
	GetProjects(engineerId uuid.UUID) ([]models.Project, error)
	UpdateProject(projectData *models.Project, engineerId uuid.UUID) (*models.Project, error)
	RemoveProject(id uuid.UUID, userId uuid.UUID) error

	CreateSkill(skillData *models.Skill) (*models.Skill, error)
	GetSkillByID(id uuid.UUID) (*models.Skill, error)

	CreateSpecialization(specializationData *models.Specialization) (*models.Specialization, error)
	GetSpecializations(engineerId uuid.UUID) ([]models.Specialization, error)
	RemoveSpecializationFromEngineer(id uuid.UUID, engineerId uuid.UUID) error

	CreateEducation(educationData *models.Education, engineerId uuid.UUID) (*models.Education, error)
	GetEducations(engineerId uuid.UUID) ([]models.Education, error)
	UpdateEducation(educationData *models.Education, engineerId uuid.UUID) (*models.Education, error)
	RemoveEducation(id uuid.UUID, userId uuid.UUID) error

	CreateCertification(certificationData *models.Certification, engineerId uuid.UUID) (*models.Certification, error)
	GetCertifications(engineerId uuid.UUID) ([]models.Certification, error)
	UpdateCertification(certificationData *models.Certification, engineerId uuid.UUID) (*models.Certification, error)
	RemoveCertification(id uuid.UUID, userId uuid.UUID) error

	CreateEngineerExperience(engineerExperienceData *models.EngineerExperience, engineerId uuid.UUID) (*models.EngineerExperience, error)
	GetEngineerExperiences(engineerId uuid.UUID) ([]models.EngineerExperience, error)
	UpdateEngineerExperience(engineerExperienceData *models.EngineerExperience, engineerId uuid.UUID) (*models.EngineerExperience, error)
	RemoveEngineerExperience(id uuid.UUID, userId uuid.UUID) error
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
		Resume:          engineerData.Resume,
		Specializations: engineerData.Specializations,
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

func (r *engineerRepository) GetEngineerSkills(engineerId uuid.UUID) ([]models.EngineerSkills, error) {

	var engineerSkill []models.EngineerSkills

	resp := r.DB.Where("engineer_id = ?", engineerId).Preload("Skill").Find(&engineerSkill)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return engineerSkill, nil
}

func (r *engineerRepository) UpdateEngineerSkill(engineerSkillData *models.EngineerSkills, userId uuid.UUID) (*models.EngineerSkills, error) {
	err := r.DB.Model(&models.EngineerSkills{}).Where("id = ? AND engineer_id = ?", engineerSkillData.SkillID, userId).Updates(&engineerSkillData).Error

	if err != nil {
		return nil, err
	}

	return engineerSkillData, nil
}

func (r *engineerRepository) RemoveEngineerSkill(id uuid.UUID, userId uuid.UUID) error {
	err := r.DB.Where("id = ? AND engineer_id = ?", id, userId).Delete(&models.EngineerSkills{}).Error

	if err != nil {
		return err
	}

	return nil
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

func (r *engineerRepository) GetProjects(engineerId uuid.UUID) ([]models.Project, error) {

	var project []models.Project

	resp := r.DB.Where("engineer_id = ?", engineerId).Find(&project)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return project, nil
}

func (r *engineerRepository) UpdateProject(projectData *models.Project, engineerId uuid.UUID) (*models.Project, error) {
	err := r.DB.Model(&models.Project{}).Where("id = ? AND engineer_id = ?", projectData.ID, engineerId).Updates(&projectData).Error

	if err != nil {
		return nil, err
	}

	return projectData, nil
}

func (r *engineerRepository) RemoveProject(id uuid.UUID, userId uuid.UUID) error {
	err := r.DB.Where("id = ? AND engineer_id = ?", id, userId).Delete(&models.Project{}).Error

	if err != nil {
		return err
	}

	return nil
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

func (r *engineerRepository) GetSkillByID(id uuid.UUID) (*models.Skill, error) {

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

func (r *engineerRepository) GetSpecializations(engineerId uuid.UUID) ([]models.Specialization, error) {

	var specialization []models.Specialization

	resp := r.DB.Where("engineer_id = ?", engineerId).Find(&specialization)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return specialization, nil
}

func (r *engineerRepository) RemoveSpecializationFromEngineer(id uuid.UUID, engineerId uuid.UUID) error {
	err := r.DB.Model(&models.EngineerModel{}).Where("user_id = ?", engineerId).Association("Specializations").Delete(&models.Specialization{ID: id})

	if err != nil {
		return err
	}

	return nil
}

func (r *engineerRepository) CreateEducation(educationData *models.Education, engineerId uuid.UUID) (*models.Education, error) {

	education := models.Education{
		Degree:        educationData.Degree,
		Institute:     educationData.Institute,
		YearOfPassing: educationData.YearOfPassing,
		CGPA:          educationData.CGPA,
		Branch:        educationData.Branch,
		EngineerID:    engineerId,
	}

	resp := r.DB.Create(&education)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &education, nil
}

func (r *engineerRepository) GetEducations(engineerId uuid.UUID) ([]models.Education, error) {

	var educations []models.Education

	resp := r.DB.Where("engineer_id = ?", engineerId).Find(&educations)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return educations, nil
}

func (r *engineerRepository) UpdateEducation(educationData *models.Education, engineerId uuid.UUID) (*models.Education, error) {
	err := r.DB.Model(&models.Education{}).Where("id = ? AND engineer_id = ?", educationData.ID, engineerId).Updates(&educationData).Error

	if err != nil {
		return nil, err
	}

	return educationData, nil
}

func (r *engineerRepository) RemoveEducation(id uuid.UUID, userId uuid.UUID) error {
	err := r.DB.Where("id = ? AND engineer_id = ?", id, userId).Delete(&models.Education{}).Error

	if err != nil {
		return err
	}

	return nil
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

func (r *engineerRepository) GetCertifications(engineerId uuid.UUID) ([]models.Certification, error) {
	var certification []models.Certification

	resp := r.DB.Where("engineer_id = ?", engineerId).Find(&certification)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return certification, nil
}

func (r *engineerRepository) UpdateCertification(certificationData *models.Certification, engineerId uuid.UUID) (*models.Certification, error) {
	err := r.DB.Model(&models.Certification{}).Where("id = ? AND engineer_id = ?", certificationData.ID, engineerId).Updates(&certificationData).Error

	if err != nil {
		return nil, err
	}

	return certificationData, nil
}

func (r *engineerRepository) RemoveCertification(id uuid.UUID, userId uuid.UUID) error {
	err := r.DB.Where("id = ? AND engineer_id = ?", id, userId).Delete(&models.Certification{}).Error

	if err != nil {
		return err
	}

	return nil
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

func (r *engineerRepository) GetEngineerExperiences(engineerId uuid.UUID) ([]models.EngineerExperience, error) {
	// fetch all experiences of the engineer
	var engineerExperience []models.EngineerExperience

	resp := r.DB.Where("engineer_id = ?", engineerId).Find(&engineerExperience)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return engineerExperience, nil
}

func (r *engineerRepository) UpdateEngineerExperience(engineerExperienceData *models.EngineerExperience, engineerId uuid.UUID) (*models.EngineerExperience, error) {

	err := r.DB.Model(&models.EngineerExperience{}).Where("id = ? AND engineer_id = ?", engineerExperienceData.ID, engineerId).Updates(&engineerExperienceData).Error

	if err != nil {
		return nil, err
	}

	return engineerExperienceData, nil
}

func (r *engineerRepository) RemoveEngineerExperience(id uuid.UUID, userId uuid.UUID) error {
	err := r.DB.Where("id = ? AND engineer_id = ?", id, userId).Delete(&models.EngineerExperience{}).Error

	if err != nil {
		return err
	}

	return nil
}
