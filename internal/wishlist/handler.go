package wishlist

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) CreateWishlistHandler(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var input CreateWishlistInput
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.CreateWishlist(&input, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Wishlist created successfully"})
}

func (h *Handler) GetWishlistByIDHandler(c *gin.Context) {
	idParam := c.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wishlist ID"})
		return
	}

	id := uint(id64)
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wishlist ID cannot be zero"})
		return
	}

	wishlist, err := h.service.GetWishlistByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, wishlist)
}

func (h *Handler) GetWishlistByUsernameAndSlugHandler(c *gin.Context) {
	username := c.Param("username")
	slug := c.Param("slug")
	if username == "" || slug == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and wishlist slug cannot be empty"})
		return
	}

	wishlist, err := h.service.GetWishlistByUsernameAndSlug(username, slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, wishlist)
}

func (h *Handler) GetWishlistBySlugHandler(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wishlist slug cannot be empty"})
		return
	}

	wishlist, err := h.service.GetWishlistBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, wishlist)
}

func (h *Handler) GetAllWishlistsByUserHandler(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	wishlists, err := h.service.GetAllWishlistsByUser(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, wishlists)
}

func (h *Handler) UpdateWishlistHandler(c *gin.Context) {
	idParam := c.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wishlist ID"})
		return
	}

	id := uint(id64)
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wishlist ID cannot be zero"})
		return
	}

	var input Wishlist
	if err = c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.ID = id

	err = h.service.UpdateWishlist(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Wishlist updated successfully"})
}

func (h *Handler) DeleteWishlistHandler(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	idParam := c.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wishlist ID"})
		return
	}

	id := uint(id64)
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wishlist ID cannot be zero"})
		return
	}

	err = h.service.DeleteWishlist(id, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Wishlist deleted successfully"})
}
