package app

import (
	"github.com/Ahmad940/assetly_server/app/handler/ws"
	"github.com/Ahmad940/assetly_server/pkg/middleware"
	"github.com/Ahmad940/assetly_server/pkg/router"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

var cfg = websocket.Config{
	RecoverHandler: func(conn *websocket.Conn) {
		if err := recover(); err != nil {
			conn.WriteJSON(fiber.Map{"customError": "error occurred"})
		}
	},
}

func AttachRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error { return c.Status(200).JSON(fiber.Map{"message": "Hello, World!"}) })
	app.Get("/ws/:id", websocket.New(ws.Stream, cfg))

	// routes
	router.Authentication(api)
	router.User(api)
	router.OTP(api)
	router.Wallet(api)
	router.TransactionHistory(api)

	// not found
	middleware.NotFoundMiddleware(app)
}
