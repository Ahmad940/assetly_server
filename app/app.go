package app

import (
	"github.com/Ahmad940/assetly_server/pkg/config"
	"github.com/Ahmad940/assetly_server/platform/hub"

	"github.com/gofiber/fiber/v2"
)

func StartApp() {
	app := fiber.New(config.FiberConfig())

	// enable middleware
	EnableMiddlewares(app)

	// start hub
	go hub.RunHub()

	// attach routes
	AttachRoutes(app)

	StartServerWithGracefulShutdown(app)
}
