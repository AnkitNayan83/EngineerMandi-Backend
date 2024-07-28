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
