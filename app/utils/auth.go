package utils

import (
	"errors"
	"strings"
)

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
