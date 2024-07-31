package routes

import (
	"github.com/AnkitNayan83/EngineerMandi-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func EngineerRoutes(router *gin.RouterGroup, engineerController *controllers.EngineerController) {
	router.POST("/engineer", engineerController.CreateEngineer)
}
