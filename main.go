package main

import (
	"github.com/Ahmad940/assetly_server/app"
	"github.com/Ahmad940/assetly_server/platform/db"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// connecting to database and initialize migrations
	db.InitializeMigration()

	// start server
	app.StartApp()
}
