package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := extractTokenFromRequest(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		token, err := ParseAccessToken(tokenStr)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}
		userID, err := ExtractUserID(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}

func extractTokenFromRequest(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer "), nil
	}

	cookie, err := c.Cookie("access_token")
	if err == nil && cookie != "" {
		return cookie, nil
	}

	return "", errors.New("missing access token")
}

func ExtractRefreshTokenFromRequest(c *gin.Context) (string, error) {
	cookie, err := c.Cookie("refresh_token")
	if err == nil && cookie != "" {
		return cookie, nil
	}

	return "", errors.New("missing refresh token")
}
