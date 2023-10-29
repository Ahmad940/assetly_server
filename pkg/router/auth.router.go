package router

import (
	"github.com/Ahmad940/assetly_server/app/handler"
	"github.com/gofiber/fiber/v2"
)

func Authentication(app fiber.Router) {
	auth := app.Group("/auth")

	auth.Get("/profile", handler.Profile)
	auth.Post("/login", handler.Login)
	auth.Post("/register", handler.Register)
}
