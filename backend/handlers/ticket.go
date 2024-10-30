package handlers

import (
	"context"
	"time"

	"github.com/HemlockPham7/backend/models"
	"github.com/gofiber/fiber/v2"
)

type TicketHandler struct {
	repository models.TicketRepository
}

func NewTicketHandler(router fiber.Router, repository models.TicketRepository) {
	handler := &TicketHandler{
		repository: repository,
	}

	router.Get("/", handler.GetMany)
	router.Post("/", handler.CreateOne)
	router.Get("/:ticketId", handler.GetOne)
	router.Put("/validate/:ticketId", handler.ValidateOne)
}

func (h *TicketHandler) GetMany(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	tickets, err := h.repository.GetMany(context)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "",
		"data":    tickets,
	})
}

func (h *TicketHandler) GetOne(ctx *fiber.Ctx) error {
	ticketId := ctx.Params("ticketId")

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	ticket, err := h.repository.GetOne(context, ticketId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "",
		"data": &fiber.Map{
			"ticket": ticket,
		},
	})
}

func (h *TicketHandler) CreateOne(ctx *fiber.Ctx) error {
	ticket := &models.Ticket{}

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(ticket); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	ticket, err := h.repository.CreateOne(context, ticket)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "Ticket created",
		"data":    ticket,
	})
}

func (h *TicketHandler) ValidateOne(ctx *fiber.Ctx) error {
	ticketId := ctx.Params("ticketId")
	validateData := make(map[string]interface{})
	validateData["entered"] = true

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	ticket, err := h.repository.UpdateOne(context, ticketId, validateData)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "Ticket valid, welcome to the show!",
		"data":    ticket,
	})
}
