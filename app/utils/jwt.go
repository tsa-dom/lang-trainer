package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GetAuthToken(username, priviledges string) (string, error) {
	jwtKey := []byte(os.Getenv("SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":    username,
		"priviledges": priviledges,
		"time":        time.Now(),
	})

	tokenString, err := token.SignedString(jwtKey)
	fmt.Println(tokenString)
	if err != nil {
		fmt.Println()
		return "", err
	}

	return tokenString, nil
}
