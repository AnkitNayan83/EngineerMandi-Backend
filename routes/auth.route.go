package routes

import (
	"github.com/AnkitNayan83/SMA-backend/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.GET("/auth/google/login", controllers.GoogleLogin)
	router.GET("/api/v1/auth/google/callback", controllers.GoogleCallback)
}
