package repositories

import (
	"fmt"
	"time"

	"github.com/AnkitNayan83/EngineerMandi-Backend/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EngineerRepository interface {
	GetEngineers(specializationIds []uuid.UUID, skillIds []uuid.UUID) ([]models.EngineerModel, error)
	UpdateEngineer(engineerData *models.EngineerModel) (*models.EngineerModel, error)
	GetEngineerByID(id uuid.UUID) (*models.EngineerModel, error)
	UpdateEngineerResume(resumeUrl string, userId uuid.UUID) error

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
	AddEngineerSpecailization(specializationId uuid.UUID, engineerId uuid.UUID) error

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

	AddRating(rating *models.Rating, userId uuid.UUID) error
	GetRatingsByEngineerID(engineerId uuid.UUID, page int) ([]models.Rating, error)
	UpdateRating(rating *models.Rating, userId uuid.UUID) (*models.Rating, error)
	RemoveRating(id uuid.UUID, userId uuid.UUID) error
	GetRatingsAverage(engineerId uuid.UUID) (float64, error)
}

type engineerRepository struct {
	DB *gorm.DB
}

func NewEngineerRepository(db *gorm.DB) EngineerRepository {
	return &engineerRepository{DB: db}
}

func (r *engineerRepository) GetEngineers(specializationIds []uuid.UUID, skillIds []uuid.UUID) ([]models.EngineerModel, error) {
	var engineers []models.EngineerModel

	query := r.DB.Model(&models.EngineerModel{})

	if len(specializationIds) > 0 {
		query = query.Joins("JOIN engineer_specializations ON engineer_specializations.engineer_model_user_id = engineer_models.user_id").
			Where("engineer_specializations.specialization_id IN ?", specializationIds)
	}

	if len(skillIds) > 0 {
		query = query.Joins("JOIN engineer_skills ON engineer_skills.engineer_id = engineer_models.user_id").
			Where("engineer_skills.skill_id IN ?", skillIds)
	}

	resp := query.Preload("User").Find(&engineers)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return engineers, nil
}

func (r *engineerRepository) UpdateEngineer(engineerData *models.EngineerModel) (*models.EngineerModel, error) {

	err := r.DB.Model(&models.EngineerModel{}).Where("user_id = ?", engineerData.UserId).Updates(&engineerData).Error

	if err != nil {
		return nil, err
	}

	return engineerData, nil
}

func (r *engineerRepository) UpdateEngineerResume(resumeUrl string, userId uuid.UUID) error {

	err := r.DB.Model(&models.EngineerModel{}).Where("user_id = ?", userId).Update("resume", resumeUrl).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *engineerRepository) GetEngineerByID(id uuid.UUID) (*models.EngineerModel, error) {

	var engineer models.EngineerModel

	resp := r.DB.Preload("User").Where("user_id = ?", id).First(&engineer)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return &engineer, nil
}

func (r *engineerRepository) CreateEngineerSkill(engineerSkillData *models.EngineerSkills, userId uuid.UUID) (*models.EngineerSkills, error) {
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
	err := r.DB.Model(&models.EngineerSkills{}).Where("skill_id = ? AND engineer_id = ?", engineerSkillData.SkillID, userId).Updates(&engineerSkillData).Error

	if err != nil {
		return nil, err
	}

	return engineerSkillData, nil
}

func (r *engineerRepository) RemoveEngineerSkill(id uuid.UUID, userId uuid.UUID) error {
	err := r.DB.Where("skill_id = ? AND engineer_id = ?", id, userId).Delete(&models.EngineerSkills{}).Error

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

	resp := r.DB.Where("engineer_id = ?", engineerId).Preload("ProjectUrls").Find(&project)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return project, nil
}

func (r *engineerRepository) UpdateProject(projectData *models.Project, engineerId uuid.UUID) (*models.Project, error) {
	tx := r.DB.Begin()

	err := tx.Model(&models.Project{}).Where("id = ? AND engineer_id = ?", projectData.ID, engineerId).Updates(map[string]interface{}{
		"name":        projectData.Name,
		"description": projectData.Description,
	}).Error

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var existingUrls []models.ProjectUrl
	err = tx.Where("project_id = ?", projectData.ID).Find(&existingUrls).Error

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	existingUrlsMap := make(map[uuid.UUID]models.ProjectUrl)
	for _, url := range existingUrls {
		existingUrlsMap[url.ID] = url
	}

	for _, projectUrl := range projectData.ProjectUrls {
		if projectUrl.ID == uuid.Nil {
			projectUrl.ProjectID = projectData.ID
			err = tx.Create(&projectUrl).Error
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		} else {
			err = tx.Model(&models.ProjectUrl{}).Where("id = ? AND project_id = ?", projectUrl.ID, projectData.ID).Updates(map[string]interface{}{
				"url":  projectUrl.Url,
				"type": projectUrl.Type,
			}).Error

			if err != nil {
				tx.Rollback()
				return nil, err
			}

			delete(existingUrlsMap, projectUrl.ID)
		}
	}

	for _, url := range existingUrlsMap {
		err = tx.Delete(&url).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	err = tx.Commit().Error
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
	var specializations []models.Specialization

	resp := r.DB.Joins("JOIN engineer_specializations ON engineer_specializations.specialization_id = specializations.id").
		Where("engineer_specializations.engineer_model_user_id = ?", engineerId).
		Find(&specializations)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return specializations, nil
}

func (r *engineerRepository) AddEngineerSpecailization(specializationId uuid.UUID, engineerId uuid.UUID) error {
	var engineer models.EngineerModel
	if err := r.DB.First(&engineer, "user_id = ?", engineerId).Error; err != nil {
		return fmt.Errorf("engineer not found: %w", err)
	}

	// Ensure the specialization exists
	var specialization models.Specialization
	if err := r.DB.First(&specialization, "id = ?", specializationId).Error; err != nil {
		return fmt.Errorf("specialization not found: %w", err)
	}

	// Append the specialization to the engineer's Specializations association
	if err := r.DB.Model(&engineer).Association("Specializations").Append(&specialization); err != nil {
		return fmt.Errorf("failed to add specialization: %w", err)
	}

	return nil
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

func (r *engineerRepository) AddRating(rating *models.Rating, engineerId uuid.UUID) error {
	rating.EngineerID = engineerId
	rating.CreatedAt = time.Now()

	resp := r.DB.Create(&rating)

	if resp.Error != nil {
		return resp.Error
	}

	return nil
}

func (r *engineerRepository) GetRatingsByEngineerID(engineerId uuid.UUID, page int) ([]models.Rating, error) {

	var ratings []models.Rating

	resp := r.DB.Where("engineer_id = ?", engineerId).Find(&ratings).Limit(10).Offset((page - 1) * 10)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return ratings, nil
}

func (r *engineerRepository) UpdateRating(rating *models.Rating, engineerId uuid.UUID) (*models.Rating, error) {

	err := r.DB.Model(&models.Rating{}).Where("id = ? AND engineer_id = ?", rating.ID, engineerId).Updates(&rating).Error

	if err != nil {
		return nil, err
	}

	return rating, nil
}

func (r *engineerRepository) RemoveRating(id uuid.UUID, userId uuid.UUID) error {
	err := r.DB.Where("id = ? AND engineer_id = ?", id, userId).Delete(&models.Rating{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *engineerRepository) GetRatingsAverage(engineerId uuid.UUID) (float64, error) {
	var result struct {
		totalRating float64
		count       int
	}

	resp := r.DB.Model(&models.Rating{}).
		Where("engineer_id = ?", engineerId).
		Select("SUM(stars) as totalRating, COUNT(*) as count").
		Scan(&result)

	if resp.Error != nil {
		return 0, resp.Error
	}

	if result.count == 0 {
		return 0, nil
	}

	return result.totalRating / float64(result.count), nil
}
