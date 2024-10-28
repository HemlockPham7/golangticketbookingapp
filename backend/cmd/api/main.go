package main

import (
	"github.com/HemlockPham7/backend/handlers"
	"github.com/HemlockPham7/backend/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "TicketBooking",
		ServerHeader: "Fiber",
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(nil)

	// Routing
	server := app.Group("/api")

	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(":3000")
}
