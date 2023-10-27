package middleware

import (
	"fmt"

	"github.com/Ahmad940/assetly_server/app/model"
	"github.com/Ahmad940/assetly_server/app/service"
	"github.com/Ahmad940/assetly_server/pkg/util"
	"github.com/gofiber/fiber/v2"
)

type RoleConfig struct {
	// List Of Role To br allowed
	Roles []string
}

func RoleAuthorization(config RoleConfig) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// retrieving token meta data
		tokenData, err := util.ExtractTokenMetadata(ctx)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
				Message: err.Error(),
			})
		}

		// verify that the request user is admin
		user, err := service.GetAUser(tokenData.ID)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
				Message: err.Error(),
			})
		}
		// fmt.Printf("user %v\n", user)

		if err != nil {
			return err
		}

		for _, role := range config.Roles {
			fmt.Printf("Role %v\n", role)
			if user.Role == role {
				return ctx.Next()
			}
		}
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Message: "Unauthorized",
		})
	}
}
