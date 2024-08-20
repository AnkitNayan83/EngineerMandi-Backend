package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AnkitNayan83/EngineerMandi-Backend/models"
	"github.com/AnkitNayan83/EngineerMandi-Backend/services"
	"github.com/AnkitNayan83/EngineerMandi-Backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EngineerController struct {
	engineerService services.EngineerService
}

func NewEngineerController(engineerService services.EngineerService) *EngineerController {
	return &EngineerController{engineerService: engineerService}
}

func (ctrl *EngineerController) GetEngineers(c *gin.Context) {
	var filterRequest struct {
		SpecializationIds []uuid.UUID `json:"specializationIds"`
		SkillIds          []uuid.UUID `json:"skillIds"`
	}

	_ = c.ShouldBindJSON(&filterRequest)

	engineers, err := ctrl.engineerService.GetEngineers(filterRequest.SpecializationIds, filterRequest.SkillIds)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": engineers})
}

func (ctrl *EngineerController) GetEngineer(c *gin.Context) {
	userID, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	engineer, err := ctrl.engineerService.GetEngineerByID(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": engineer})
}

func (ctrl *EngineerController) UpdateEngineerResume(c *gin.Context) {
	userID, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	engineerData := models.EngineerModel{}
	err = c.ShouldBindJSON(&engineerData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.engineerService.UpdateEngineerResume(engineerData.Resume, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "resume updated successfully"})
}

func (ctrl *EngineerController) UpdateOrAddEngineerExperience(c *gin.Context) {

	userID, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	engineerExperienceData := []models.EngineerExperience{}
	err = c.ShouldBindJSON(&engineerExperienceData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty data in the request"})
		return
	}

	for _, exp := range engineerExperienceData {
		if exp.ID == uuid.Nil {
			_, err := ctrl.engineerService.CreateEngineerExperience(exp, userID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		} else {
			log.Println(exp)
			_, err := ctrl.engineerService.UpdateEngineerExperience(exp, userID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "experiences updated successfully"})
}

func (ctrl *EngineerController) RemoveExperience(c *gin.Context) {

	userID, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	engineerExperienceIdStr := c.Params.ByName("id")
	engineerExperienceId, err := uuid.Parse(engineerExperienceIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid experience id"})
		return
	}

	err = ctrl.engineerService.RemoveEngineerExperience(engineerExperienceId, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "experience removed successfully"})
}

func (ctrl *EngineerController) GetEngineerExperiences(c *gin.Context) {

	userID, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	experiences, err := ctrl.engineerService.GetEngineerExperiences(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": experiences})
}

func (ctrl *EngineerController) UpdateOrAddEducation(c *gin.Context) {
	userID, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	educationData := models.Education{}
	err = c.ShouldBindJSON(&educationData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if educationData.ID == uuid.Nil {
		_, err := ctrl.engineerService.CreateEducation(educationData, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		x, err := ctrl.engineerService.UpdateEducation(educationData, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		log.Print(x)
	}

	c.JSON(http.StatusOK, gin.H{"message": "education updated successfully"})

}

func (ctrl *EngineerController) RemoveEducation(c *gin.Context) {
	userID, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	educationIdStr := c.Params.ByName("id")
	educationId, err := uuid.Parse(educationIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid education id"})
		return
	}

	err = ctrl.engineerService.RemoveEducation(educationId, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "education removed successfully"})
}

func (ctrl *EngineerController) GetEngineerEducations(c *gin.Context) {
	userId, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	educations, err := ctrl.engineerService.GetEducations(userId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "education not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": educations})

}

func (ctrl *EngineerController) UpdateOrAddEngineerSkill(c *gin.Context) {
	userID, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	isNew := c.Query("new")

	skillData := models.EngineerSkills{}
	err = c.ShouldBindJSON(&skillData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if isNew == "true" {
		_, err := ctrl.engineerService.CreateEngineerSkill(skillData, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		log.Print("gg")
		_, err := ctrl.engineerService.UpdateEngineerSkill(skillData, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "skill updated successfully"})
}

func (ctrl *EngineerController) RemoveEngineerSkill(c *gin.Context) {
	userID, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	skillIdStr := c.Params.ByName("id")
	skillId, err := uuid.Parse(skillIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid skill id"})
		return
	}

	err = ctrl.engineerService.RemoveEngineerSkill(skillId, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "skill removed successfully"})
}

func (ctrl *EngineerController) GetEngineerSkills(c *gin.Context) {
	userId, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	skills, err := ctrl.engineerService.GetEngineerSkills(userId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "skills not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": skills})
}

func (ctrl *EngineerController) UpdateOrAddEngineerCertification(c *gin.Context) {
	userID, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	certificationData := models.Certification{}
	err = c.ShouldBindJSON(&certificationData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if certificationData.ID == uuid.Nil {
		_, err := ctrl.engineerService.CreateCertification(certificationData, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		_, err := ctrl.engineerService.UpdateCertification(certificationData, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "certification updated successfully"})
}

func (ctrl *EngineerController) RemoveEngineerCertification(c *gin.Context) {
	userID, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	certificationIdStr := c.Params.ByName("id")
	certificationId, err := uuid.Parse(certificationIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid certification id"})
		return
	}

	err = ctrl.engineerService.RemoveCertification(certificationId, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "certification removed successfully"})
}

func (ctrl *EngineerController) GetEngineerCertifications(c *gin.Context) {
	userId, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	certifications, err := ctrl.engineerService.GetCertifications(userId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "certifications not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": certifications})
}

func (ctrl *EngineerController) UpdateOrAddEngineerProject(c *gin.Context) {
	userID, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	projectData := models.Project{}
	err = c.ShouldBindJSON(&projectData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if projectData.ID == uuid.Nil {
		_, err := ctrl.engineerService.CreateProject(projectData, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		_, err := ctrl.engineerService.UpdateProject(projectData, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "project updated successfully"})
}

func (ctrl *EngineerController) RemoveEngineerProject(c *gin.Context) {
	userID, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	projectIdStr := c.Params.ByName("id")
	projectId, err := uuid.Parse(projectIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project id"})
		return
	}

	err = ctrl.engineerService.RemoveProject(projectId, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "project removed successfully"})
}

func (ctrl *EngineerController) GetEngineerProjects(c *gin.Context) {
	userId, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	projects, err := ctrl.engineerService.GetProjects(userId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "projects not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": projects})
}

func (ctrl *EngineerController) GetEngineerSpecialization(c *gin.Context) {
	userId, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	specializations, err := ctrl.engineerService.GetSpecializations(userId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "specializations not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": specializations})
}

func (ctrl *EngineerController) AddEngineerSpecailization(c *gin.Context) {
	userID, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	specializationData := models.Specialization{}
	err = c.ShouldBindJSON(&specializationData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = ctrl.engineerService.AddEngineerSpecailization(specializationData.ID, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "specialization added successfully"})
}

func (ctrl *EngineerController) RemoveEngineerSpecailization(c *gin.Context) {
	userID, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	specializationIdStr := c.Params.ByName("id")
	specializationId, err := uuid.Parse(specializationIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid specialization id"})
		return
	}

	err = ctrl.engineerService.RemoveSpecialization(specializationId, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "specialization removed successfully"})
}

func (ctrl *EngineerController) AddRating(c *gin.Context) {
	userID, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ratingData := models.Rating{}
	err = c.ShouldBindJSON(&ratingData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = ctrl.engineerService.AddRating(userID, ratingData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "rating added successfully"})
}

func (ctrl *EngineerController) GetEngineerRating(c *gin.Context) {

	userID, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var page int
	pageStr := c.Query("page")

	if pageStr == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}
	}

	rating, err := ctrl.engineerService.GetRatings(userID, page)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

func (ctrl *EngineerController) UpdateRating(c *gin.Context) {

	userID, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ratingData := models.Rating{}
	err = c.ShouldBindJSON(&ratingData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = ctrl.engineerService.UpdateRating(ratingData, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "rating updated successfully"})
}

func (ctrl *EngineerController) RemoveRating(c *gin.Context) {

	userID, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ratingIdStr := c.Params.ByName("id")
	ratingId, err := uuid.Parse(ratingIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid rating id"})
		return
	}

	err = ctrl.engineerService.RemoveRating(ratingId, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "rating removed successfully"})
}

func (ctrl *EngineerController) GetRatingsAverage(c *gin.Context) {
	userID, err := utils.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	average, err := ctrl.engineerService.GetRatingsAverage(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": average})
}
