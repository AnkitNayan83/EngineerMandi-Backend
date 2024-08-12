package controllers

import (
	"log"
	"net/http"

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

func (ctrl *EngineerController) CreateEngineer(c *gin.Context) {

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

	engineer, err := ctrl.engineerService.CreateEngineer(engineerData, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	engineer.User = models.User{}

	c.JSON(http.StatusOK, gin.H{"data": engineer})
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

	engineerExperienceData := models.EngineerExperience{}
	err = c.ShouldBindJSON(&engineerExperienceData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty data in the request"})
		return
	}

	err = ctrl.engineerService.RemoveEngineerExperience(engineerExperienceData.ID, userID)

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

	educationData := models.Education{}
	err = c.ShouldBindJSON(&educationData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty data in the request"})
	}

	err = ctrl.engineerService.RemoveEducation(educationData.ID, userID)

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

	isNew := c.Params.ByName("isNew")

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

	skillData := models.EngineerSkills{}
	err = c.ShouldBindJSON(&skillData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty data in the request"})
	}

	err = ctrl.engineerService.RemoveEngineerSkill(skillData.SkillID, userID)

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
