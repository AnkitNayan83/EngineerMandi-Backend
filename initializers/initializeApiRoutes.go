package initializers

import (
	"os"

	"github.com/AnkitNayan83/SMA-backend/controllers"
	"github.com/AnkitNayan83/SMA-backend/repositories"
	"github.com/AnkitNayan83/SMA-backend/routes"
	"github.com/AnkitNayan83/SMA-backend/services"
	"github.com/gin-gonic/gin"
)

func InitializeApiRoutes() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	services.InitializeOAuth()

	router := gin.Default()

	userRepo := repositories.NewUserRepository(DB)
	userService := services.NewUserService(userRepo)
	authController := controllers.NewAuthController(userService)

	routes.InitializeTestRoutes(router)
	routes.AuthRoutes(router, authController)
	router.Run(":" + port)
}
