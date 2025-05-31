package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateJWT(jwtSecret []byte, userID string, username string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 144).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
