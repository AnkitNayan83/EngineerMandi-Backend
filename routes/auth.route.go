package routes

import (
	"github.com/AnkitNayan83/SMA-backend/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, authController *controllers.AuthController) {
	router.GET("/auth/google/login", authController.GoogleLogin)
	router.GET("/api/v1/auth/google/callback", authController.GoogleCallback)
}
