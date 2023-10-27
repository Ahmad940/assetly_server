package service

import (
	"github.com/Ahmad940/assetly_server/app/model"
	"github.com/Ahmad940/assetly_server/pkg/constant"
	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(err error, ctx *fiber.Ctx) error {
	if err.Error() == constant.SqlNotFoundText {
		return ctx.Status(fiber.StatusNotFound).JSON(model.ErrorResponse{
			Message: "no record found",
		})
	}
	if err.Error() == "redis: nil" {
		return ctx.Status(fiber.StatusNotFound).JSON(model.ErrorResponse{
			Message: "no key found or session expired",
		})
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
		Message: err.Error(),
	})
}

func SqlErrorNotFound(err error) bool {
	return err.Error() == constant.SqlNotFoundText
}

func SqlErrorIgnoreNotFound(err error) error {
	if err == nil {
		return nil
	}
	if err.Error() == constant.SqlNotFoundText {
		return nil
	}
	return err
}
