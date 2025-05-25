package user

import (
	"log"
	"net/http"

	"github.com/BekzhanK1/wishly/config"
	"github.com/BekzhanK1/wishly/internal/auth"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) ProfileHandler(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	userResponse, err := h.service.Me(userID)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": userResponse})

}

func (h *Handler) Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Register(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (h *Handler) Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginOutput, err := h.service.ValidateCredentials(input)
	if err != nil {
		if err == ErrUserNotFound || err == ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to login"})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    loginOutput.AccessToken,
		MaxAge:   config.AccessTokenExpiryTime,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    loginOutput.RefreshToken,
		MaxAge:   config.RefreshTokenExpiryTime,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})

	c.JSON(http.StatusOK, gin.H{"user": loginOutput.UserResponse})
}

func (h *Handler) RefreshAccessTokenHandler(c *gin.Context) {
	refreshToken, err := auth.ExtractRefreshTokenFromRequest(c)
	if err != nil {
		log.Println("Error retrieving refresh token:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
		return
	}

	loginOutput, err := h.service.RefreshAccessToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    loginOutput.AccessToken,
		MaxAge:   config.AccessTokenExpiryTime,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})

	c.JSON(http.StatusOK, gin.H{"user": loginOutput.UserResponse})
}
