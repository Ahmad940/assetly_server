package util

import (
	"time"

	"github.com/Ahmad940/assetly_server/app/model"
	"github.com/Ahmad940/assetly_server/pkg/config"
	"github.com/Ahmad940/assetly_server/platform/db"
	"github.com/golang-jwt/jwt/v5"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gopkg.in/guregu/null.v4"
)

func GenerateToken(user model.User) (string, error) {
	if user.Session.String == "" {
		session_id := gonanoid.Must()
		err := db.DB.Model(&model.User{}).Where("id = ?", user.ID).Update("session", session_id).Error
		if err != nil {
			return "", err
		}
		user.Session = null.StringFrom(session_id)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"id":      user.ID,
		"session": user.Session.String,
		"age":     time.Now().Unix(),
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
