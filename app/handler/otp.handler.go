package handler

import (
	"github.com/Ahmad940/assetly_server/app/service"
	"github.com/gofiber/fiber/v2"
)

func RequestOTP(ctx *fiber.Ctx) error {
	phone := ctx.Query("phone")
	// fetching the current logged user base on the mid retrieved from meta data
	err := service.RequestOtp(phone)

	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(fiber.Map{
		"message": "OTP sent successfully",
	})
}

func VerifyOTP(ctx *fiber.Ctx) error {
	phone := ctx.Query("phone")
	code := ctx.Query("code")

	// fetching the current logged user base on the mid retrieved from meta data
	err := service.VerifyOtp(phone, code)

	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(fiber.Map{
		"message": "OTP matched",
	})
}
