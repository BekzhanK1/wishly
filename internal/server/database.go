package server

import (
	"log"

	"github.com/BekzhanK1/wishly/config"
	"gorm.io/gorm"
)

func ConnectToDatabase() *gorm.DB {
	envDBConfig := config.NewEnvDBConfig()
	db, err := config.ConnectDB(*envDBConfig)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}
	log.Print("Successfully connected to database")
	return db
}
