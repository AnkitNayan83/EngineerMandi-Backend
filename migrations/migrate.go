package main

import (
	"log"

	"github.com/AnkitNayan83/EngineerMandi-Backend/initializers"
	"github.com/AnkitNayan83/EngineerMandi-Backend/models"
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
		&models.EngineerExperience{},
	)

	if err != nil {
		log.Fatalf("Migration failed🚫: %v", err)
	}

	log.Println("Migration successful🗃️")
}
