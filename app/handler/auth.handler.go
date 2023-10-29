package handler

import (
	"github.com/Ahmad940/assetly_server/app/model"
	"github.com/Ahmad940/assetly_server/app/service"
	"github.com/Ahmad940/assetly_server/pkg/util"
	"github.com/gofiber/fiber/v2"
)

func Profile(ctx *fiber.Ctx) error {
	// retrieving token meta data
	tokenData, err := util.ExtractTokenMetadata(ctx)

	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	// fetching the current logged user base on the mid retrieved from meta data
	user, err := service.GetAUser(tokenData.ID)

	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(user)
}

func Login(ctx *fiber.Ctx) error {
	var body model.Login
	// parsing response body
	err := ctx.BodyParser(&body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	// validating the user
	errors := util.ValidateStruct(body)
	if len(errors) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Requesting OTP
	response, err := service.Login(body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(response)
}

func Register(ctx *fiber.Ctx) error {
	var body model.CreateUser

	// parsing response body
	err := ctx.BodyParser(&body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	// validating the user
	errors := util.ValidateStruct(body)
	if len(errors) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// retrieving the token by passing request body
	response, err := service.CreateAccount(body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(response)
}
