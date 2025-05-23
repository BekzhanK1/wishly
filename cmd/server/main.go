package main

import (
	"log"

	"github.com/BekzhanK1/wishly/config"
	"github.com/BekzhanK1/wishly/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	r := gin.Default()

	envDBConfig := config.NewEnvDBConfig()
	db, err := config.ConnectDB(*envDBConfig)

	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	api := r.Group("/api")
	user.RegisterRoutes(api, userHandler)

	if err := r.Run(":8000"); err != nil {
		panic(err)
	}
}
