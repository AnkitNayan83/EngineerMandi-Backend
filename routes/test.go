package routes

import (
	"github.com/AnkitNayan83/SMA-backend/controllers"
	"github.com/AnkitNayan83/SMA-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func InitializeTestRoutes(router *gin.RouterGroup) {

	router.GET("/users", middlewares.TestMiddleware(), controllers.GetUsers)
	router.POST("/users", controllers.CreateUser)
}
