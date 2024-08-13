package services

import (
	"errors"
	"log"

	"github.com/AnkitNayan83/EngineerMandi-Backend/models"
	"github.com/AnkitNayan83/EngineerMandi-Backend/repositories"
	"github.com/google/uuid"
)

type EngineerService interface {
	CreateEngineer(engineerData models.EngineerModel, userId uuid.UUID) (*models.EngineerModel, error)
	GetEngineerByID(userId uuid.UUID) (*models.EngineerModel, error)
	UpdateEngineerResume(resumeUrl string, userId uuid.UUID) error

	CreateEngineerSkill(engineerSkillData models.EngineerSkills, userId uuid.UUID) (*models.EngineerSkills, error)
	GetEngineerSkills(engineerId uuid.UUID) ([]models.EngineerSkills, error)
	UpdateEngineerSkill(engineerSkillData models.EngineerSkills, userId uuid.UUID) (*models.EngineerSkills, error)
	RemoveEngineerSkill(id uuid.UUID, userId uuid.UUID) error

	CreateEducation(educationData models.Education, engineerId uuid.UUID) (*models.Education, error)
	GetEducations(engineerId uuid.UUID) ([]models.Education, error)
	UpdateEducation(educationData models.Education, engineerId uuid.UUID) (*models.Education, error)
	RemoveEducation(id uuid.UUID, userId uuid.UUID) error

	CreateCertification(certificationData models.Certification, engineerId uuid.UUID) (*models.Certification, error)
	GetCertifications(engineerId uuid.UUID) ([]models.Certification, error)
	UpdateCertification(certificationData models.Certification, engineerId uuid.UUID) (*models.Certification, error)
	RemoveCertification(id uuid.UUID, userId uuid.UUID) error

	CreateProject(projectData models.Project, engineerId uuid.UUID) (*models.Project, error)
	GetProjects(engineerId uuid.UUID) ([]models.Project, error)
	UpdateProject(projectData models.Project, engineerId uuid.UUID) (*models.Project, error)
	RemoveProject(id uuid.UUID, userId uuid.UUID) error

	CreateSpecialization(specializationData models.Specialization) (*models.Specialization, error)
	GetSpecializations(engineerId uuid.UUID) ([]models.Specialization, error)
	AddEngineerSpecailization(specializationId uuid.UUID, userId uuid.UUID) error
	RemoveSpecialization(id uuid.UUID, userId uuid.UUID) error

	CreateEngineerExperience(engineerExperienceData models.EngineerExperience, engineerId uuid.UUID) (*models.EngineerExperience, error)
	GetEngineerExperiences(engineerId uuid.UUID) ([]models.EngineerExperience, error)
	UpdateEngineerExperience(engineerExperienceData models.EngineerExperience, engineerId uuid.UUID) (*models.EngineerExperience, error)
	RemoveEngineerExperience(id uuid.UUID, userId uuid.UUID) error
}

type engineerService struct {
	repo repositories.EngineerRepository
}

func NewEngineerService(repo repositories.EngineerRepository) EngineerService {
	return &engineerService{repo: repo}
}

func (s *engineerService) GetEngineerByID(userId uuid.UUID) (*models.EngineerModel, error) {
	engineer, err := s.repo.GetEngineerByID(userId)

	if err != nil {
		return nil, err
	}

	return engineer, nil
}

func (s *engineerService) UpdateEngineerResume(resumeUrl string, userId uuid.UUID) error {

	err := s.repo.UpdateEngineerResume(resumeUrl, userId)

	if err != nil {
		return err
	}

	return nil
}

func (s *engineerService) CreateEngineer(engineerData models.EngineerModel, userId uuid.UUID) (*models.EngineerModel, error) {

	if engineerData.Resume == "" {
		return nil, errors.New("engineer resume is required")
	}

	newEngineer, err := s.repo.CreateEngineer(&engineerData, userId)

	if err != nil {
		return nil, err
	}

	return newEngineer, nil

}

func (s *engineerService) GetEngineerSkills(engineerId uuid.UUID) ([]models.EngineerSkills, error) {
	engineerSkill, err := s.repo.GetEngineerSkills(engineerId)

	if err != nil {
		return nil, err
	}

	return engineerSkill, nil
}

func (s *engineerService) CreateEngineerSkill(engineerSkillData models.EngineerSkills, userId uuid.UUID) (*models.EngineerSkills, error) {
	if engineerSkillData.SkillID == uuid.Nil {
		return nil, errors.New("skill id is required")
	}

	engineerSkill, err := s.repo.CreateEngineerSkill(&engineerSkillData, userId)

	if err != nil {
		return nil, err
	}

	return engineerSkill, nil
}

func (s *engineerService) UpdateEngineerSkill(engineerSkillData models.EngineerSkills, userId uuid.UUID) (*models.EngineerSkills, error) {
	if engineerSkillData.SkillID == uuid.Nil {
		return nil, errors.New("skill id is required")
	}

	engineerSkill, err := s.repo.UpdateEngineerSkill(&engineerSkillData, userId)

	if err != nil {
		return nil, err
	}

	return engineerSkill, nil
}

func (s *engineerService) RemoveEngineerSkill(skillId uuid.UUID, userId uuid.UUID) error {
	if skillId == uuid.Nil {
		return errors.New("skill id is required")
	}

	err := s.repo.RemoveEngineerSkill(skillId, userId)

	if err != nil {
		return err
	}

	return nil
}

func (s *engineerService) GetEducations(engineerId uuid.UUID) ([]models.Education, error) {
	education, err := s.repo.GetEducations(engineerId)

	if err != nil {
		return nil, err
	}

	return education, nil
}

func (s *engineerService) CreateEducation(educationData models.Education, userId uuid.UUID) (*models.Education, error) {
	if educationData.Degree == "" {
		return nil, errors.New("education degree is required")
	}

	if educationData.Institute == "" {
		return nil, errors.New("education institute is required")
	}

	if educationData.Branch == "" {
		return nil, errors.New("education branch is required")
	}

	if educationData.YearOfPassing == 0 {
		return nil, errors.New("education year of passing is required")
	}

	if educationData.CGPA == 0 {
		return nil, errors.New("education cgpa is required")
	}

	education, err := s.repo.CreateEducation(&educationData, userId)

	if err != nil {
		return nil, err
	}

	return education, nil
}

func (s *engineerService) UpdateEducation(educationData models.Education, userId uuid.UUID) (*models.Education, error) {
	if educationData.Degree == "" {
		return nil, errors.New("education degree is required")
	}

	if educationData.Institute == "" {
		return nil, errors.New("education institute is required")
	}

	if educationData.Branch == "" {
		return nil, errors.New("education branch is required")
	}

	if educationData.YearOfPassing == 0 {
		return nil, errors.New("education year of passing is required")
	}

	if educationData.CGPA == 0 {
		return nil, errors.New("education cgpa is required")
	}

	log.Print(educationData)

	education, err := s.repo.UpdateEducation(&educationData, userId)

	if err != nil {
		return nil, err
	}

	return education, nil
}

func (s *engineerService) RemoveEducation(educationId uuid.UUID, userId uuid.UUID) error {
	if educationId == uuid.Nil {
		return errors.New("education id is required")
	}

	err := s.repo.RemoveEducation(educationId, userId)

	if err != nil {
		return err
	}

	return nil
}

func (s *engineerService) GetCertifications(engineerId uuid.UUID) ([]models.Certification, error) {
	certifications, err := s.repo.GetCertifications(engineerId)

	if err != nil {
		return nil, err
	}

	return certifications, nil
}

func (s *engineerService) CreateCertification(certificationData models.Certification, userId uuid.UUID) (*models.Certification, error) {
	if certificationData.Name == "" {
		return nil, errors.New("certification name is required")
	}

	if certificationData.CertificateUrl == "" {
		return nil, errors.New("certification certificate url is required")
	}

	if certificationData.IssuedDate.IsZero() {
		return nil, errors.New("certification issued date is required")
	}

	certificate, err := s.repo.CreateCertification(&certificationData, userId)

	if err != nil {
		return nil, err
	}

	return certificate, nil
}

func (s *engineerService) UpdateCertification(certificationData models.Certification, userId uuid.UUID) (*models.Certification, error) {
	if certificationData.Name == "" {
		return nil, errors.New("certification name is required")
	}

	if certificationData.CertificateUrl == "" {
		return nil, errors.New("certification certificate url is required")
	}

	if certificationData.IssuedDate.IsZero() {
		return nil, errors.New("certification issued date is required")
	}

	certificate, err := s.repo.UpdateCertification(&certificationData, userId)

	if err != nil {
		return nil, err
	}

	return certificate, nil
}

func (s *engineerService) RemoveCertification(certificationId uuid.UUID, userId uuid.UUID) error {
	if certificationId == uuid.Nil {
		return errors.New("certification id is required")
	}

	err := s.repo.RemoveCertification(certificationId, userId)

	if err != nil {
		return err
	}

	return nil
}

func (s *engineerService) GetProjects(engineerId uuid.UUID) ([]models.Project, error) {
	projects, err := s.repo.GetProjects(engineerId)

	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (s *engineerService) CreateProject(projectData models.Project, userId uuid.UUID) (*models.Project, error) {
	if projectData.Name == "" {
		return nil, errors.New("project name is required")
	}

	project, err := s.repo.CreateProject(&projectData, userId)

	if err != nil {
		return nil, err
	}

	return project, nil
}

func (s *engineerService) UpdateProject(projectData models.Project, userId uuid.UUID) (*models.Project, error) {
	if projectData.Name == "" {
		return nil, errors.New("project name is required")
	}

	project, err := s.repo.UpdateProject(&projectData, userId)

	if err != nil {
		return nil, err
	}

	return project, nil
}

func (s *engineerService) RemoveProject(projectId uuid.UUID, userId uuid.UUID) error {
	if projectId == uuid.Nil {
		return errors.New("project id is required")
	}

	err := s.repo.RemoveProject(projectId, userId)

	if err != nil {
		return err
	}

	return nil
}

func (s *engineerService) CreateSpecialization(specializationData models.Specialization) (*models.Specialization, error) {

	if specializationData.Title == "" {
		return nil, errors.New("specialization title is required")
	}

	specialization, err := s.repo.CreateSpecialization(&specializationData)

	if err != nil {
		return nil, err
	}

	return specialization, nil
}

func (s *engineerService) AddEngineerSpecailization(specializationId uuid.UUID, userId uuid.UUID) error {
	if specializationId == uuid.Nil {
		return errors.New("specialization id is required")
	}

	err := s.repo.AddEngineerSpecailization(specializationId, userId)

	if err != nil {
		return err
	}

	return nil
}

func (s *engineerService) RemoveSpecialization(specializationId uuid.UUID, userId uuid.UUID) error {
	if specializationId == uuid.Nil {
		return errors.New("specialization id is required")
	}

	err := s.repo.RemoveSpecializationFromEngineer(specializationId, userId)

	if err != nil {
		return err
	}

	return nil
}

func (s *engineerService) GetSpecializations(engineerId uuid.UUID) ([]models.Specialization, error) {
	specializations, err := s.repo.GetSpecializations(engineerId)

	if err != nil {
		return nil, err
	}

	return specializations, nil
}

func (s *engineerService) CreateEngineerExperience(experienceData models.EngineerExperience, userId uuid.UUID) (*models.EngineerExperience, error) {
	if experienceData.Company == "" {
		return nil, errors.New("experience company is required")
	}

	if experienceData.Location == "" {
		return nil, errors.New("experience location is required")
	}

	if experienceData.Role == "" {
		return nil, errors.New("experience role is required")
	}

	if experienceData.StartDate.IsZero() {
		return nil, errors.New("experience start date is required")
	}

	if !experienceData.IsCurrent && experienceData.EndDate.IsZero() {
		return nil, errors.New("experience end date is required for a past experience")
	}

	experience, err := s.repo.CreateEngineerExperience(&experienceData, userId)

	if err != nil {
		return nil, err
	}

	return experience, nil
}

func (s *engineerService) UpdateEngineerExperience(experienceData models.EngineerExperience, userId uuid.UUID) (*models.EngineerExperience, error) {
	if experienceData.Company == "" {
		return nil, errors.New("experience company is required")
	}

	if experienceData.Location == "" {
		return nil, errors.New("experience location is required")
	}

	if experienceData.Role == "" {
		return nil, errors.New("experience role is required")
	}

	if experienceData.StartDate.IsZero() {
		return nil, errors.New("experience start date is required")
	}

	if !experienceData.IsCurrent {
		if experienceData.EndDate.IsZero() {

			return nil, errors.New("experience end date is required for a past experience")
		}
		if experienceData.EndDate.Before(experienceData.StartDate) {
			return nil, errors.New("experience end date should be after start date")
		}

	}

	log.Println(experienceData.Description)

	experience, err := s.repo.UpdateEngineerExperience(&experienceData, userId)

	if err != nil {
		return nil, err
	}

	return experience, nil
}

func (s *engineerService) RemoveEngineerExperience(experienceId uuid.UUID, userId uuid.UUID) error {
	if experienceId == uuid.Nil {
		return errors.New("experience id is required")
	}

	err := s.repo.RemoveEngineerExperience(experienceId, userId)

	if err != nil {
		return err
	}

	return nil
}

func (s *engineerService) GetEngineerExperiences(engineerId uuid.UUID) ([]models.EngineerExperience, error) {
	experience, err := s.repo.GetEngineerExperiences(engineerId)

	if err != nil {
		return nil, err
	}

	return experience, nil
}
