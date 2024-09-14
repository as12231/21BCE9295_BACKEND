package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var keyy = []byte("L8fAApjINlIJdmP+qLLnaBOtL0tEsZZhAi5oB8oM4l4=") // Replace with a secure key

// GenerateJWT generates a new JWT token for the user
func GenerateJWT(userID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token valid for 24 hours
	claims := &jwt.StandardClaims{
		Subject:   userID,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(keyy)
	return tokenString, err
}
