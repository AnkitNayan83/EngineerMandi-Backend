package controllers

import (
	"net/http"

	"github.com/AnkitNayan83/EngineerMandi-Backend/models"
	"github.com/AnkitNayan83/EngineerMandi-Backend/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (ctrl *UserController) ProfileSetup(c *gin.Context) {

	user, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id not found"})
		return
	}

	userID, ok := user.(string)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id not found"})
		return
	}

	var updateUserData models.User

	if err := c.ShouldBindJSON(&updateUserData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userData, err := ctrl.userService.ProfileSetup(userID, &updateUserData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userData)

}

func (ctrl *UserController) GetUserInfo(c *gin.Context) {

	user, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user id not found in the token"})
		return
	}

	userId := user.(string)

	userInfo, err := ctrl.userService.GetUserInfo(userId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
	}

	c.JSON(http.StatusOK, userInfo)

}

func (ctrl *UserController) CreateEngineer(c *gin.Context) {

	var engineer models.EngineerModel

	if err := c.ShouldBindJSON(&engineer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.userService.CreateEngineer(engineer)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Engineer created successfully"})
}
