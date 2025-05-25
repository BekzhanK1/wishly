package user

import (
	"github.com/BekzhanK1/wishly/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
	users := rg.Group("/users")

	users.POST("/register", h.Register)
	users.POST("/login", h.Login)
	users.POST("/refresh", h.RefreshAccessTokenHandler)

	usersAuth := users.Group("/")
	usersAuth.Use(auth.JWTMiddleware())
	usersAuth.GET("/me", h.ProfileHandler)
}
