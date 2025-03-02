package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"auth-service/pkg/apperrors"
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
		return "", jwt.ErrTokenMalformed // или просто return "", errors.New("JWT_SECRET not set")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return "", err // Возвращаем ошибку как есть
	}

	if !token.Valid {
		return "", jwt.ErrTokenMalformed
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", jwt.ErrTokenMalformed
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", jwt.ErrTokenMalformed
	}

	return userID, nil
}
