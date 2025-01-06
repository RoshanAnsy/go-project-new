package config

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectDB() {
	dns := "postgresql://db-zippyfin_owner:wG7Yo5CvRINq@ep-misty-hill-a5m81po4.us-east-2.aws.neon.tech/db-zippyfin?sslmode=require"

	var err error

	DB, err = gorm.Open(postgres.Open(dsn),  &gorm.Config{})

	if err !=nil {
		log.Fatal("failed to connect database: ", err)
	}
	log.Println("Connected to the database!")

}