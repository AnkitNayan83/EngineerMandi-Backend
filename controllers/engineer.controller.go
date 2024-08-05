package controllers

import (
	"net/http"

	"github.com/AnkitNayan83/EngineerMandi-Backend/models"
	"github.com/AnkitNayan83/EngineerMandi-Backend/services"
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

	user, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user id not found"})
		return
	}

	//convert user to uuid
	userIDStr, ok := user.(string)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	userID, err := uuid.Parse(userIDStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
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
