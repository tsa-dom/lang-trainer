package utils

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/tsa-dom/lang-trainer/app/models/users"
)

type Claims struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	Priviledges string `json:"priviledges"`
	jwt.StandardClaims
}

func CreateAuthToken(username string) (string, error) {
	user, err := users.GetUserByUsername(username)
	if err != nil {
		return "", err
	}
	jwtKey := []byte(os.Getenv("SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":          user.Id,
		"username":    user.Username,
		"priviledges": user.Priviledges,
		"time":        time.Now(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func VerifyUser(authorization string) (*Claims, error) {
	if authorization == "" {
		return nil, errors.New("no Authorization header provided")
	}
	token := strings.TrimPrefix(authorization, "Bearer ")
	if authorization == token {
		return nil, errors.New("token should have Bearer prefix")
	}

	verification := verifyAuthToken(token)

	if verification == nil {
		return nil, errors.New("invalid Authorization token")
	}

	return verification, nil
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
