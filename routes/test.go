package routes

import (
	"github.com/AnkitNayan83/EngineerMandi-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeTestRoutes(router *gin.RouterGroup) {
	router.GET("/ping", controllers.HealthCheckup)
}
