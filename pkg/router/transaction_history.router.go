package router

import (
	"github.com/Ahmad940/assetly_server/app/handler"
	"github.com/gofiber/fiber/v2"
)

func TransactionHistory(app fiber.Router) {
	auth := app.Group("/wallet/history")

	auth.Get("/:wallet_no", handler.GetWalletHistory)
}
