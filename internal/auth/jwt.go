package auth

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"

	"time"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func GenerateJWT(userId int, loginId string) (string, error) {
	claims := jwt.MapClaims{
		"userId":   userId,
		"loginId":  loginId,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey , nil
	})
}