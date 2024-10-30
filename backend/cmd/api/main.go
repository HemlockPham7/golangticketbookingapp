package main

import (
	"fmt"

	"github.com/HemlockPham7/backend/config"
	"github.com/HemlockPham7/backend/db"
	"github.com/HemlockPham7/backend/handlers"
	"github.com/HemlockPham7/backend/middlewares"
	"github.com/HemlockPham7/backend/repositories"
	"github.com/HemlockPham7/backend/services"
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
	ticketRepository := repositories.NewTicketRepository(db)
	authRepository := repositories.NewAuthRepository(db)

	// Service
	authService := services.NewAuthService(authRepository)

	// Routing
	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	privateRoutes := server.Use(middlewares.AuthProtected(db))

	handlers.NewEventHandler(privateRoutes.Group("/event"), eventRepository)
	handlers.NewTicketHandler(privateRoutes.Group("/ticket"), ticketRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
