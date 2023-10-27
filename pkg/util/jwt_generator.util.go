package util

import (
	"time"

	"github.com/Ahmad940/assetly_server/pkg/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id string) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"id":  id,
		"age": time.Now().Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	encodedToken, err := token.SignedString([]byte(config.GetEnv().JWT_SECRET))
	if err != nil {
		return "", err
	}
	return encodedToken, nil
}
