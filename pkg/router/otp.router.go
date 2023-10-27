package router

import (
	"github.com/Ahmad940/assetly_server/app/handler"
	"github.com/gofiber/fiber/v2"
)

func OTP(app fiber.Router) {
	auth := app.Group("/auth/otp")

	auth.Get("/request-otp", handler.RequestOTP)
	auth.Get("/verify-otp", handler.VerifyOTP)
}
