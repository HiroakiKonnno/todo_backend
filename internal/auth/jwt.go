package auth

import (
	"errors"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
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

func ExtractAndVerifyToken(ctx *gin.Context) (map[string]interface{}, bool) {
	tokenString, err := ctx.Cookie("jwt")
	if err != nil {
		return nil, false
	}
	claims, err := parseAndValidateToken(tokenString)
	if err != nil {
		return nil, false
	}
	return claims, true
}

func parseAndValidateToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}