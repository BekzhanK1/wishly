package server

import (
	"github.com/BekzhanK1/wishly/internal/user"
	"github.com/BekzhanK1/wishly/internal/wishlist"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	wishlistRepo := wishlist.NewRepository(db)
	wishlistService := wishlist.NewService(wishlistRepo)
	wishlistHandler := wishlist.NewHandler(wishlistService)

	api := r.Group("/api")
	user.RegisterRoutes(api, userHandler)
	wishlist.RegisterRoutes(api, wishlistHandler)

	return r
}
