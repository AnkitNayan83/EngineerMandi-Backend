package routes

import (
	"fmt"

	"github.com/AnkitNayan83/EngineerMandi-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, userController *controllers.UserController) {

	const pathPrefix = "/user"

	router.GET(fmt.Sprint(pathPrefix+"/info"), userController.GetUserInfo)
	router.PATCH(fmt.Sprint(pathPrefix+"/profile-setup"), userController.ProfileSetup)
	router.POST(fmt.Sprint(pathPrefix+"/engineer"), userController.CreateEngineer)
}
