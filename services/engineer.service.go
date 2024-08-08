package services

import (
	"errors"

	"github.com/AnkitNayan83/EngineerMandi-Backend/models"
	"github.com/AnkitNayan83/EngineerMandi-Backend/repositories"
	"github.com/google/uuid"
)

type EngineerService interface {
	CreateEngineer(engineerData models.EngineerModel, userId uuid.UUID) (*models.EngineerModel, error)
	CreateEngineerSkill(engineerSkillData models.EngineerSkills, userId uuid.UUID) (*models.EngineerSkills, error)
	CreateEducation(educationData models.Education, userId uuid.UUID) (*models.Education, error)
	CreateCertification(certificationData models.Certification, userId uuid.UUID) (*models.Certification, error)
	CreateProject(projectData models.Project, userId uuid.UUID) (*models.Project, error)
	CreateSpecialization(specializationData models.Specialization) (*models.Specialization, error)
	CreateEngineerExperience(experienceData models.EngineerExperience, userId uuid.UUID) (*models.EngineerExperience, error)
	UpdateEngineer(engineerData models.EngineerModel, userId uuid.UUID) (*models.EngineerModel, error)
}

type engineerService struct {
	repo repositories.EngineerRepository
}

func NewEngineerService(repo repositories.EngineerRepository) EngineerService {
	return &engineerService{repo: repo}
}

func (s *engineerService) CreateEngineer(engineerData models.EngineerModel, userId uuid.UUID) (*models.EngineerModel, error) {
	var specializations []models.Specialization
	var skills []models.EngineerSkills
	var educations []models.Education
	var certifications []models.Certification
	var projects []models.Project
	var engineerExperiences []models.EngineerExperience

	newEngineer, err := s.repo.CreateEngineer(&engineerData, userId)

	if err != nil {
		return nil, err
	}

	if engineerData.Resume == "" {
		return nil, errors.New("engineer resume is required")
	}

	if len(engineerData.Specializations) > 0 {
		for _, specialization := range engineerData.Specializations {
			if specialization.ID != uuid.Nil {
				specializations = append(specializations, specialization)
			} else {
				return nil, errors.New("specialization id not found")
			}
		}

		newEngineer.Specializations = specializations
	}

	if len(engineerData.Skills) > 0 {
		for _, skill := range engineerData.Skills {
			if skill.SkillID != uuid.Nil {
				engineerSkill, err := s.CreateEngineerSkill(skill, userId)

				if err != nil {
					return nil, err
				}

				skills = append(skills, *engineerSkill)
			} else {
				return nil, errors.New("skill id not found")
			}
		}
		newEngineer.Skills = skills
	}

	if len(engineerData.Education) > 0 {
		for _, education := range engineerData.Education {
			newEducation, err := s.CreateEducation(education, userId)

			if err != nil {
				return nil, err
			}

			educations = append(educations, *newEducation)

		}
		newEngineer.Education = educations
	}

	if len(engineerData.Certifications) > 0 {
		for _, certification := range engineerData.Certifications {
			newCertification, err := s.CreateCertification(certification, userId)

			if err != nil {
				return nil, err
			}

			certifications = append(certifications, *newCertification)
		}
		newEngineer.Certifications = certifications
	}

	if len(engineerData.Projects) > 0 {
		for _, project := range engineerData.Projects {
			newProject, err := s.CreateProject(project, userId)

			if err != nil {
				return nil, err
			}

			projects = append(projects, *newProject)
		}
		newEngineer.Projects = projects
	}

	if len(engineerData.Experiences) > 0 {
		for _, experience := range engineerData.Experiences {
			newExperience, err := s.CreateEngineerExperience(experience, userId)

			if err != nil {
				return nil, err
			}

			engineerExperiences = append(engineerExperiences, *newExperience)
		}
		newEngineer.Experiences = engineerExperiences
	}

	updatedEngineer, err := s.UpdateEngineer(*newEngineer, userId)

	if err != nil {
		return nil, err
	}

	return updatedEngineer, nil

}

func (s *engineerService) UpdateEngineer(engineerData models.EngineerModel, userId uuid.UUID) (*models.EngineerModel, error) {
	currentEngineer, err := s.repo.GetEngineerByID(userId)

	if err != nil {
		return nil, err
	}

	if engineerData.Resume != "" {
		currentEngineer.Resume = engineerData.Resume
	}
	if len(engineerData.Specializations) > 0 {
		currentEngineer.Specializations = engineerData.Specializations
	}

	if len(engineerData.Skills) > 0 {
		currentEngineer.Skills = engineerData.Skills
	}

	if len(engineerData.Education) > 0 {
		currentEngineer.Education = engineerData.Education
	}

	if len(engineerData.Certifications) > 0 {
		currentEngineer.Certifications = engineerData.Certifications
	}

	if len(engineerData.Projects) > 0 {
		currentEngineer.Projects = engineerData.Projects
	}

	if len(engineerData.Experiences) > 0 {
		currentEngineer.Experiences = engineerData.Experiences
	}

	engineer, err := s.repo.UpdateEngineer(currentEngineer)

	if err != nil {
		return nil, err
	}

	return engineer, nil

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

	education, err := s.repo.UpdateEducation(&educationData, userId)

	if err != nil {
		return nil, err
	}

	return education, nil
}
