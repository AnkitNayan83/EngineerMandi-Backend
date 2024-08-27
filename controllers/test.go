package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckup(c *gin.Context) {

	msg := fmt.Sprintf("Server is up and running on port %s", c.Request.Host)

	c.JSON(http.StatusOK, gin.H{"message": msg})
}
