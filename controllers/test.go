package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckup(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Server is up and running"})
}
