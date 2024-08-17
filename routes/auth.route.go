package routes

import (
	"fmt"

	"github.com/AnkitNayan83/EngineerMandi-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup, authController *controllers.AuthController) {

	const pathPrefix = "/auth"

	router.GET(fmt.Sprint(pathPrefix+"/google/login"), authController.GoogleLogin)
	router.GET(fmt.Sprint(pathPrefix+"/google/callback"), authController.GoogleCallback)
}
