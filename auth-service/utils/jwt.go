package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"auth-service/internal/apperrors"
)

func GenerateJWT(userID int) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")

	if jwtSecret == "" {
		return jwtSecret, apperrors.ErrNoJwtSecret
	}

	claims := jwt.MapClaims{
		"user_id": strconv.Itoa(userID),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecret))
}

func ValidateJWT(tokenString string) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")

	if jwtSecret == "" {
		return jwtSecret, apperrors.ErrNoJwtSecret
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, apperrors.ErrInvalidToken
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return "", apperrors.ErrInvalidToken
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			fmt.Println("Token exp:", exp)
			fmt.Println("Current time:", time.Now().Unix())
			if time.Now().Unix() > int64(exp) {
				return "", apperrors.ErrExpiredToken
			}
		}

		if userID, ok := claims["user_id"].(string); ok {
			return userID, nil
		}
	}

	return "", apperrors.ErrInvalidToken
}
