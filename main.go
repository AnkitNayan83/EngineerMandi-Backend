package main

import (
	"github.com/AnkitNayan83/SMA-backend/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.InitializeApiRoutes()
}

func main() {
}
