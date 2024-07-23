package main

import (
	"github.com/AnkitNayan83/SMA-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.InitializeTestRoutes(router)

	router.Run(":8080")
}