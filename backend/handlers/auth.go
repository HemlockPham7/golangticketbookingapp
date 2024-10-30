package handlers

import (
	"github.com/HemlockPham7/backend/models"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service models.AuthService
}

func NewAuthHandler(route fiber.Router, service models.AuthService) {
	handler := &AuthHandler{
		service: service,
	}

	route.Post("/login", handler.Login)
	route.Post("/register", handler.Register)
}

func (h *AuthHandler) Login(ctx *fiber.Ctx) error {

	return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
		"status":  "fail",
		"message": "nil",
	})
}

func (h *AuthHandler) Register(ctx *fiber.Ctx) error {

	return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
		"status":  "fail",
		"message": "nil",
	})
}
