package main

import (
	"log"

	"github.com/AnkitNayan83/SMA-backend/initializers"
	"github.com/AnkitNayan83/SMA-backend/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	Migrate()
}

func Migrate() {
	err := initializers.DB.AutoMigrate(
		&models.User{},
		&models.Rating{},
		&models.EngineerModel{},
		&models.EngineerSkills{},
		&models.Skill{},
		&models.Specialization{},
		&models.Education{},
		&models.Certification{},
		&models.Project{},
		&models.ProjectUrl{},
	)

	if err != nil {
		log.Fatalf("Migration failed🚫: %v", err)
	}

	log.Println("Migration successful🗃️")
}
