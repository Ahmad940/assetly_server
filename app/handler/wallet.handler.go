package handler

import (
	"github.com/Ahmad940/assetly_server/app/model"
	"github.com/Ahmad940/assetly_server/app/service"
	"github.com/Ahmad940/assetly_server/pkg/util"
	"github.com/gofiber/fiber/v2"
)

func GetWalletByNo(ctx *fiber.Ctx) error {
	wallet_no := ctx.Params("wallet_no")
	fetchFullDetail := util.ParseBoolean(ctx.Query("full-detail"))

	wallet, err := service.GetWalletByNo(wallet_no)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	if fetchFullDetail {
		return ctx.JSON(wallet)
	} else {
		return ctx.JSON(fiber.Map{
			"first_name": wallet.User.FirstName,
			"last_name":  wallet.User.LastName,
			"wallet_no":  wallet_no,
		})
	}
}

func TopUpWallet(ctx *fiber.Ctx) error {
	var body model.TopUp
	// parsing response body
	err := ctx.BodyParser(&body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	// validating the body
	errors := util.ValidateStruct(body)
	if len(errors) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Admin update users admin
	user, err := service.TopUpWallet(body.WalletNo, body.Amount)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(user)
}
