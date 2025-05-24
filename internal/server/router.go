package server

import (
	"github.com/BekzhanK1/wishly/internal/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	api := r.Group("/api")
	user.RegisterRoutes(api, userHandler)

	return r
}
