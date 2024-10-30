package main

import (
	"fmt"

	"github.com/HemlockPham7/backend/config"
	"github.com/HemlockPham7/backend/db"
	"github.com/HemlockPham7/backend/handlers"
	"github.com/HemlockPham7/backend/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      envConfig.APPName,
		ServerHeader: envConfig.SVHeader,
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(db)

	// Routing
	server := app.Group("/api")

	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
