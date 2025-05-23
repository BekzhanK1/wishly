package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
	users := rg.Group("/users")
	{
		users.POST("/register", h.Register)
		users.POST("/login", h.Login)
	}

}
