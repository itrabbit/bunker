package main

import (
	"os"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/helmet"
	"github.com/itrabbit/bunker/config"
	"github.com/itrabbit/bunker/controllers"
	"github.com/itrabbit/bunker/db"
)

func main() {
	if err := db.Open(
		db.Dialect(config.GetDbDialect()),
		config.GetDbArgs(),
	); err != nil {
		panic(err)
	}
	// Init storage catalog
	if _, err := os.Stat(config.GetStoragePath()); os.IsNotExist(err) {
		err = os.MkdirAll(config.GetStoragePath(), 0777)
		if err != nil {
			panic(err)
		}
	}
	// Create app
	app := fiber.New()
	app.Use(
		middleware.Recover(),
		middleware.RequestID(),
		helmet.New(),
		cors.New(),
	)
	defer func() {
		if err := db.Close(); err != nil {
			println(err)
		}
	}()
	db.AutoMigrate()
	api := app.Group("/api", cors.New())
	api.Get("/versions", func(c *fiber.Ctx) {
		if err := c.Status(200).JSON(controllers.GetVersions()); err != nil {
			c.Next(err)
		}
	})
	controllers.Init(api, "v1")
	if err := app.Listen(config.GetBindPort()); err != nil {
		panic(err)
	}
}
