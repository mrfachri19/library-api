package handler

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mrfachri19/digital-library-backend/internal/entity"
	"github.com/mrfachri19/digital-library-backend/internal/usecase"
)

type LendingHandler struct {
	Usecase usecase.LendingUsecase
}

func NewLendingHandler(u usecase.LendingUsecase) *LendingHandler {
	return &LendingHandler{Usecase: u}
}

// GET /lendings
func (h *LendingHandler) GetAll(c *fiber.Ctx) error {
	data, err := h.Usecase.GetAll(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(data)
}

// POST /lendings
func (h *LendingHandler) Create(c *fiber.Ctx) error {
	var lending entity.Lending
	if err := c.BodyParser(&lending); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if lending.BookID == 0 || lending.Borrower == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Book ID and borrower required"})
	}

	if err := h.Usecase.Create(context.Background(), lending); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"message": "Lending created"})
}

// PUT /lendings/:id/return
func (h *LendingHandler) MarkAsReturned(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.Usecase.MarkAsReturned(context.Background(), id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Book returned"})
}
