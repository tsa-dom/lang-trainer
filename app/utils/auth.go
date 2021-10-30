package utils

import (
	"errors"
	"strings"
)

func VerifyUser(auth string) (*Claims, error) {
	if auth == "" {
		return nil, errors.New("no Authorization header provided")
	}
	token := strings.TrimPrefix(auth, "Bearer ")
	if auth == token {
		return nil, errors.New("no token provided")
	}

	verification := verifyAuthToken(token)

	if verification == nil {
		return nil, errors.New("invalid Authorization token")
	}

	return verification, nil
}
