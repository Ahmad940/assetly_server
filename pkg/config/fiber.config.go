package config

import "github.com/gofiber/fiber/v2"

// "github.com/goccy/go-json"
// "github.com/gofiber/fiber/v2"

func FiberConfig() fiber.Config {
	return fiber.Config{
		// UnescapePath: true,
		// JSONEncoder: json.Marshal,
		// JSONDecoder: json.Unmarshal,
	}
}
