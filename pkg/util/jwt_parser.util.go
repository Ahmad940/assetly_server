package util

import (
	"errors"
	"os"
	"strings"

	"github.com/Ahmad940/assetly_server/app/model"
	"github.com/Ahmad940/assetly_server/pkg/constant"
	"github.com/Ahmad940/assetly_server/platform/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	ID  string  `json:"id"`
	Age float64 `json:"age"`
}

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// Expires time.
		id := claims["id"].(string)
		age := claims["age"].(float64)

		// validating user
		var user model.User

		err := db.DB.Where("id = ?", id).First(&user).Error
		if err != nil {
			if err.Error() == constant.SqlNotFoundText {
				return &TokenMetadata{}, errors.New("invalid token")
			} else {
				return &TokenMetadata{}, err
			}
		}

		return &TokenMetadata{
			ID:  id,
			Age: age,
		}, nil
	}

	return nil, err
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET")), nil
}
