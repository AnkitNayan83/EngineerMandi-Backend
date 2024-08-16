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
	router.GET(fmt.Sprint(pathPrefix), engineerController.GetEngineer)
	router.PATCH(fmt.Sprint(pathPrefix+"/resume"), engineerController.UpdateEngineerResume)

	// Experience Routes
	router.PATCH(fmt.Sprint(pathPrefix+"/experience/append-experience"), engineerController.UpdateOrAddEngineerExperience)
	router.DELETE(fmt.Sprint(pathPrefix+"/experience/:id"), engineerController.RemoveExperience)
	router.GET(fmt.Sprint(pathPrefix+"/experience"), engineerController.GetEngineerExperiences)

	//Education Routes
	router.PATCH(fmt.Sprint(pathPrefix+"/education/append-education"), engineerController.UpdateOrAddEducation)
	router.DELETE(fmt.Sprint(pathPrefix+"/education/:id"), engineerController.RemoveEducation)
	router.GET(fmt.Sprint(pathPrefix+"/education"), engineerController.GetEngineerEducations)

	//Skill Routes
	router.PATCH(fmt.Sprint(pathPrefix+"/skill/append-skill"), engineerController.UpdateOrAddEngineerSkill)
	router.DELETE(fmt.Sprint(pathPrefix+"/skill/:id"), engineerController.RemoveEngineerSkill)
	router.GET(fmt.Sprint(pathPrefix+"/skill"), engineerController.GetEngineerSkills)

	//Certification Routes
	router.PATCH(fmt.Sprint(pathPrefix+"/certification/append-certification"), engineerController.UpdateOrAddEngineerCertification)
	router.DELETE(fmt.Sprint(pathPrefix+"/certification/:id"), engineerController.RemoveEngineerCertification)
	router.GET(fmt.Sprint(pathPrefix+"/certification"), engineerController.GetEngineerCertifications)

	//Project Routes
	router.PATCH(fmt.Sprint(pathPrefix+"/project/append-project"), engineerController.UpdateOrAddEngineerProject)
	router.DELETE(fmt.Sprint(pathPrefix+"/project/:id"), engineerController.RemoveEngineerProject)
	router.GET(fmt.Sprint(pathPrefix+"/project"), engineerController.GetEngineerProjects)

	//Specialization Routes
	router.POST(fmt.Sprint(pathPrefix+"/specialization/append-specialization"), engineerController.AddEngineerSpecailization)
	router.DELETE(fmt.Sprint(pathPrefix+"/specialization/:id"), engineerController.RemoveEngineerSpecailization)
	router.GET(fmt.Sprint(pathPrefix+"/specialization"), engineerController.GetEngineerSpecialization)

	//Rating Routes
	router.POST(fmt.Sprint(pathPrefix+"/rating"), engineerController.AddRating)
	router.GET(fmt.Sprint(pathPrefix+"/rating"), engineerController.GetEngineerRating)
	router.GET(fmt.Sprint(pathPrefix+"/rating/average"), engineerController.GetRatingsAverage)
	router.PATCH(fmt.Sprint(pathPrefix+"/rating"), engineerController.UpdateRating)
	router.DELETE(fmt.Sprint(pathPrefix+"/rating/:id"), engineerController.RemoveRating)

}
