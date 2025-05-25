package wishlist

import (
	"github.com/BekzhanK1/wishly/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
	wishlists := rg.Group("/wishlists")
	wishlists.GET("user/:username/:slug", h.GetWishlistByUsernameAndSlugHandler)
	wishlists.Use(auth.JWTMiddleware())

	wishlists.POST("/", h.CreateWishlistHandler)
	wishlists.GET("/slug/:slug", h.GetWishlistBySlugHandler)
	wishlists.GET("/:id", h.GetWishlistByIDHandler)
	wishlists.GET("/", h.GetAllWishlistsByUserHandler)
	wishlists.PUT("/:id", h.UpdateWishlistHandler)
	wishlists.DELETE("/:id", h.DeleteWishlistHandler)
}
