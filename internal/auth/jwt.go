package auth

import (
	"os"
	"time"

	"errors"

	"github.com/BekzhanK1/wishly/config"
	"github.com/golang-jwt/jwt/v5"
)

var accessSecret = []byte(os.Getenv("JWT_ACCESS_SECRET"))
var refreshSecret = []byte(os.Getenv("JWT_REFRESH_SECRET"))

func GenerateAccessToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(config.AccessTokenExpiryTime * time.Second).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(accessSecret)
}

func GenerateRefreshToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(config.RefreshTokenExpiryTime * time.Second).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshSecret)
}

func GenerateTokenPair(userID uint) (string, string, error) {
	accessToken, err := GenerateAccessToken(userID)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := GenerateRefreshToken(userID)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func ParseAccessToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return accessSecret, nil
	})
}

func ParseRefreshToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpeced signing method")
		}
		return refreshSecret, nil
	})
}

func ExtractUserID(token *jwt.Token) (uint, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		idFloat, ok := claims["user_id"].(float64)
		if !ok {
			return 0, errors.New("user_id claim missing or invalid")
		}
		return uint(idFloat), nil
	}
	return 0, errors.New("invalid token")
}
