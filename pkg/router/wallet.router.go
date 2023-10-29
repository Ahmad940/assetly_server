package router

import (
	"github.com/Ahmad940/assetly_server/app/handler"
	"github.com/gofiber/fiber/v2"
)

func Wallet(app fiber.Router) {
	auth := app.Group("/wallet")

	auth.Get("/detail/:wallet_no", handler.GetWalletByNo)
	auth.Post("/top-up", handler.TopUpWallet)
}
