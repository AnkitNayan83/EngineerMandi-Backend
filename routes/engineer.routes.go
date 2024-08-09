package routes

import (
	"github.com/AnkitNayan83/EngineerMandi-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func EngineerRoutes(router *gin.RouterGroup, engineerController *controllers.EngineerController) {
	// Engineer Routes
	router.POST("/engineer", engineerController.CreateEngineer)

	// Experience Routes
	router.PATCH("/engineer/update-experience", engineerController.UpdateOrAddEngineerExperience)
	router.DELETE("/engineer/delete-experience", engineerController.RemoveExperience)
	router.GET("/engineer/experience", engineerController.GetEngineerExperiences)
}
