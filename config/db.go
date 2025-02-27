package config

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/roshanAnsy/go-project-new/models"
)

var DB *gorm.DB
func ConnectDB() {
	dns := "postgresql://db-zippyfin_owner:wG7Yo5CvRINq@ep-misty-hill-a5m81po4.us-east-2.aws.neon.tech/db-zippyfin?sslmode=require"

	var err error

	DB, err = gorm.Open(postgres.Open(dsn),  &gorm.Config{})

	if err !=nil {
		log.Fatal("failed to connect database: ", err)
	}
	// Auto-migrate User model
	DB.AutoMigrate(&models.User{})
	log.Println("Database connected successfully!")

	log.Println("Connected to the database!")

}