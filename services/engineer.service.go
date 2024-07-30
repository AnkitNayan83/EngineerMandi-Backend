package services

import (
	"errors"

	"github.com/AnkitNayan83/EngineerMandi-Backend/models"
	"github.com/AnkitNayan83/EngineerMandi-Backend/repositories"
	"github.com/google/uuid"
)

type EngineerService interface {
}

type engineerService struct {
	repo repositories.EngineerRepository
}

func NewEngineerService(repo repositories.EngineerRepository) EngineerService {
	return &engineerService{repo: repo}
}

func (s *engineerService) CreateEngineer(engineerData models.EngineerModel, userId uuid.UUID) (*models.EngineerModel, error) {
	var engineer models.EngineerModel
	var specializations []models.Specialization
	var skills []models.EngineerSkills
	var educations []models.Education
	var certifications []models.Certification
	var projects []models.Project
	var engineerExperiences []models.EngineerExperience

	if engineerData.Resume == "" {
		return nil, errors.New("Engineer resume is required")
	}

	if len(engineerData.Specializations) > 0 {
		for _, specialization := range engineerData.Specializations {
			if specialization.ID != uuid.Nil {
				specializations = append(specializations, specialization)
			} else {
				newSpecialization, err := s.CreateSpecialization(specialization)

				if err != nil {
					return nil, err
				}

				specializations = append(specializations, *newSpecialization)
			}
		}
	}

	if len(engineerData.Skills) > 0 {
		for _, skill := range engineerData.Skills {
			if skill.SkillID != uuid.Nil {
				engineerSkill, err := s.CreateEngineerSkill(skill, userId)

				if err != nil {
					return nil, err
				}

				skills = append(skills, *engineerSkill)
			}
		}
	}

	if len(engineerData.Education) > 0 {
		for _, education := range engineerData.Education {
			newEducation, err := s.CreateEducation(education, userId)

			if err != nil {
				return nil, err
			}

			educations = append(educations, *newEducation)

		}
	}

	if len(engineerData.Certifications) > 0 {
		for _, certification := range engineerData.Certifications {
			newCertification, err := s.CreateCertification(certification, userId)

			if err != nil {
				return nil, err
			}

			certifications = append(certifications, *newCertification)
		}
	}

	if len(engineerData.Projects) > 0 {
		for _, project := range engineerData.Projects {
			newProject, err := s.CreateProject(project, userId)

			if err != nil {
				return nil, err
			}

			projects = append(projects, *newProject)
		}
	}

	if len(engineerData.Experiences) > 0 {
		for _, experience := range engineerData.Experiences {
			newExperience, err := s.CreateEngineerExperience(experience, userId)

			if err != nil {
				return nil, err
			}

			engineerExperiences = append(engineerExperiences, *newExperience)
		}
	}

	newEngineerData := models.EngineerModel{
		UserId:          userId,
		Resume:          engineerData.Resume,
		Specializations: specializations,
		Skills:          skills,
		Education:       educations,
		Certifications:  certifications,
		Projects:        projects,
		Experiences:     engineerExperiences,
	}

	engineer, err := s.repo.CreateEngineer(&newEngineerData, userId)

}

func (s *engineerService) CreateEngineerSkill(engineerSkillData models.EngineerSkills, userId uuid.UUID) (*models.EngineerSkills, error) {
	if engineerSkillData.SkillID == uuid.Nil {
		return nil, errors.New("Skill id is required")
	}

	engineerSkill, err := s.repo.CreateEngineerSkill(&engineerSkillData, userId)

	if err != nil {
		return nil, err
	}

	return engineerSkill, nil
}

func (s *engineerService) CreateEducation(educationData models.Education, userId uuid.UUID) (*models.Education, error) {
	if educationData.Degree == "" {
		return nil, errors.New("Education degree is required")
	}

	if educationData.Institute == "" {
		return nil, errors.New("Education institute is required")
	}

	if educationData.Branch == "" {
		return nil, errors.New("Education branch is required")
	}

	if educationData.YearOfPassing == 0 {
		return nil, errors.New("Education year of passing is required")
	}

	if educationData.CGPA == 0 {
		return nil, errors.New("Education cgpa is required")
	}

	education, err := s.repo.CreateEducation(&educationData, userId)

	if err != nil {
		return nil, err
	}

	return education, nil
}

func (s *engineerService) CreateCertification(certificationData models.Certification, userId uuid.UUID) (*models.Certification, error) {
	if certificationData.Name == "" {
		return nil, errors.New("Certification name is required")
	}

	if certificationData.CertificateUrl == "" {
		return nil, errors.New("Certification certificate url is required")
	}

	if certificationData.IssuedDate.IsZero() {
		return nil, errors.New("Certification issued date is required")
	}

	certificate, err := s.repo.CreateCertification(&certificationData, userId)

	if err != nil {
		return nil, err
	}

	return certificate, nil
}

func (s *engineerService) CreateProject(projectData models.Project, userId uuid.UUID) (*models.Project, error) {
	if projectData.Name == "" {
		return nil, errors.New("Project name is required")
	}

	project, err := s.repo.CreateProject(&projectData, userId)

	if err != nil {
		return nil, err
	}

	return project, nil
}

func (s *engineerService) CreateSpecialization(specializationData models.Specialization) (*models.Specialization, error) {

	if specializationData.Title == "" {
		return nil, errors.New("Specialization title is required")
	}

	specialization, err := s.repo.CreateSpecialization(&specializationData)

	if err != nil {
		return nil, err
	}

	return specialization, nil
}

func (s *engineerService) CreateEngineerExperience(experienceData models.EngineerExperience, userId uuid.UUID) (*models.EngineerExperience, error) {
	if experienceData.Company == "" {
		return nil, errors.New("Experience company is required")
	}

	if experienceData.Location == "" {
		return nil, errors.New("Experience location is required")
	}

	if experienceData.Role == "" {
		return nil, errors.New("Experience role is required")
	}

	if experienceData.StartDate.IsZero() {
		return nil, errors.New("Experience start date is required")
	}

	if experienceData.IsCurrent == false && experienceData.EndDate.IsZero() {
		return nil, errors.New("Experience end date is required for a past experience")
	}

	experience, err := s.repo.CreateEngineerExperience(&experienceData, userId)

	if err != nil {
		return nil, err
	}

	return experience, nil
}
