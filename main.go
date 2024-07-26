package main

import (
	"github.com/AnkitNayan83/EngineerMandi-Backend/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.InitializeApiRoutes()
}

func main() {
}
