package routes

import (
	"github.com/AnkitNayan83/EngineerMandi-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, userController *controllers.UserController) {
	router.PATCH("/user/profile-setup", userController.ProfileSetup)
}
