package controllers

import (
	"net/http"

	"github.com/AnkitNayan83/SMA-backend/models"
	"github.com/AnkitNayan83/SMA-backend/services"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users,err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	createdUser, err := services.CreateUser(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdUser)
}