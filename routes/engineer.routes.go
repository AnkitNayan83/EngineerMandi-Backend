package routes

import (
	"fmt"

	"github.com/AnkitNayan83/EngineerMandi-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func EngineerRoutes(router *gin.RouterGroup, engineerController *controllers.EngineerController) {

	pathPrefix := "/engineer"

	// Engineer Routes
	router.POST(fmt.Sprint(pathPrefix), engineerController.CreateEngineer)

	// Experience Routes
	router.PATCH(fmt.Sprint(pathPrefix+"/experience/append-experience"), engineerController.UpdateOrAddEngineerExperience)
	router.DELETE(fmt.Sprint(pathPrefix+"/experience/remove-experience"), engineerController.RemoveExperience)
	router.GET(fmt.Sprint(pathPrefix+"/experience"), engineerController.GetEngineerExperiences)

	//Education Routes
	router.PATCH(fmt.Sprint(pathPrefix+"/education/append-education"), engineerController.UpdateOrAddEducation)
	router.DELETE(fmt.Sprint(pathPrefix+"/education/remove-education"), engineerController.RemoveEducation)
	router.GET(fmt.Sprint(pathPrefix+"/education"), engineerController.GetEngineerEducations)

	//Skill Routes
	router.PATCH(fmt.Sprint(pathPrefix+"/skill/append-skill"), engineerController.UpdateOrAddEngineerSkill)
	router.DELETE(fmt.Sprint(pathPrefix+"/skill/remove-skill"), engineerController.RemoveEngineerSkill)
	router.GET(fmt.Sprint(pathPrefix+"/skill"), engineerController.GetEngineerSkills)
}
