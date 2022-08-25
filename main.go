package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiber-gorm-basic/database"
	"github.com/zakariawahyu/go-fiber-gorm-basic/database/migrations"
	"github.com/zakariawahyu/go-fiber-gorm-basic/routes"
)

func main() {
	// Connection database
	database.InitDB()

	// Migrate database
	migrations.MigrationDatabase()

	// Fiber init
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Hello World",
		})
	})

	routes.RouteInit(app)
	app.Listen(":8081")
}
