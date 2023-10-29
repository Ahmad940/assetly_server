package handler

import (
	"github.com/Ahmad940/assetly_server/app/model"
	"github.com/Ahmad940/assetly_server/app/service"
	"github.com/Ahmad940/assetly_server/pkg/util"
	"github.com/gofiber/fiber/v2"
)

func GetAUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	users, err := service.GetAUser(id)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(users)
}

func GetAllUsers(ctx *fiber.Ctx) error {
	users, err := service.GetAllUsers()
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(users)
}

func UpdateUser(ctx *fiber.Ctx) error {
	var body model.UpdateUser
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

	// Update user
	user, err := service.UpdateUser(body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}
	return ctx.JSON(user)
}

func UpdateUserPassCode(ctx *fiber.Ctx) error {
	// retrieving token meta data
	tokenData, err := util.ExtractTokenMetadata(ctx)

	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	var body model.UpdateUserPassCode
	// parsing response body
	err = ctx.BodyParser(&body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	// validating the user
	errors := util.ValidateStruct(body)
	if len(errors) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Admin update users admin
	err = service.UpdateUserPassCode(tokenData.ID, body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(fiber.Map{
		"message": "pass code updated successfully",
	})
}

func UpdateUserAdmin(ctx *fiber.Ctx) error {
	var body model.UpdateUserAdmin
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

	// Admin update users admin
	user, err := service.UpdateUserAdmin(body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(user)
}
