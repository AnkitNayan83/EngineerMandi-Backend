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

var userRepo repositories.UserRepository

func InitializeAuthRoutes(router *gin.RouterGroup) {

	services.InitializeOAuth()
	userRepo = repositories.NewUserRepository(DB)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	routes.AuthRoutes(router, authController)
}

func InitializeUserRoutes(router *gin.RouterGroup) {
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	routes.UserRoutes(router, userController)
}

func InitializeEngineerRoutes(router *gin.RouterGroup) {
	engineerRepo := repositories.NewEngineerRepository(DB)
	engineerService := services.NewEngineerService(engineerRepo)
	engineerController := controllers.NewEngineerController(engineerService)

	routes.EngineerRoutes(router, engineerController)
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

	v1PublicRouter := router.Group("/api/v1")

	v1ProtectedRouter := router.Group("/api/v1")
	v1ProtectedRouter.Use(middlewares.AuthMiddleware())

	// Public Routes
	InitializeAuthRoutes(v1PublicRouter)
	InitializeTestRoutes(v1PublicRouter)

	// Protected Routes
	InitializeUserRoutes(v1ProtectedRouter)
	InitializeEngineerRoutes(v1ProtectedRouter)

	router.Run(":" + port)
}
