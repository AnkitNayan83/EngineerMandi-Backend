package initializers

import (
	"os"

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
	routes.InitializeTestRoutes(router)
	routes.AuthRoutes(router)
	router.Run(":" + port)
}
