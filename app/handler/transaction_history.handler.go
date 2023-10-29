package handler

import (
	"github.com/Ahmad940/assetly_server/app/service"
	"github.com/gofiber/fiber/v2"
)

func GetWalletHistory(ctx *fiber.Ctx) error {
	wallet_no := ctx.Params("wallet_no")

	history, err := service.GetWalletHistory(wallet_no)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(history)
}
