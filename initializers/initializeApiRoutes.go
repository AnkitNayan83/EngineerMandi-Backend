package initializers

import (
	"os"

	"github.com/AnkitNayan83/SMA-backend/routes"
	"github.com/gin-gonic/gin"
)

func InitializeApiRoutes() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	routes.InitializeTestRoutes(router)
	router.Run(":8080")
}
