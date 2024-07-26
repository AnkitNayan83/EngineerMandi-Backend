package routes

import (
	"github.com/AnkitNayan83/SMA-backend/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup, authController *controllers.AuthController) {
	router.GET("/auth/google/login", authController.GoogleLogin)
	router.GET("/auth/google/callback", authController.GoogleCallback)
}
