package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Username    string `json:"username"`
	Priviledges string `json:"priviledges"`
	jwt.StandardClaims
}

func CreateAuthToken(username, priviledges string) (string, error) {
	jwtKey := []byte(os.Getenv("SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":    username,
		"priviledges": priviledges,
		"time":        time.Now(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyAuthToken(token string) *Claims {
	jwtKey := []byte(os.Getenv("SECRET"))
	claims := &Claims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil
	}

	return claims
}
