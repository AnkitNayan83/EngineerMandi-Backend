package initializers

import (
	"os"

	"github.com/AnkitNayan83/EngineerMandi-Backend/controllers"
	"github.com/AnkitNayan83/EngineerMandi-Backend/middlewares"
	"github.com/AnkitNayan83/EngineerMandi-Backend/repositories"
	"github.com/AnkitNayan83/EngineerMandi-Backend/routes"
	"github.com/AnkitNayan83/EngineerMandi-Backend/services"
	"github.com/gin-gonic/gin"
)

func InitializeAuthRoutes(router *gin.RouterGroup) {

	services.InitializeOAuth()
	userRepo := repositories.NewUserRepository(DB)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	routes.AuthRoutes(router, authController)
}

func InitializeTestRoutes(router *gin.RouterGroup) {
	routes.InitializeTestRoutes(router)
}

func InitializeApiRoutes() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	v1AuthRouter := router.Group("/api/v1")

	v1ProtectedRouter := router.Group("/api/v1")
	v1ProtectedRouter.Use(middlewares.AuthMiddleware())

	InitializeTestRoutes(v1ProtectedRouter)
	InitializeAuthRoutes(v1AuthRouter)

	router.Run(":" + port)
}
