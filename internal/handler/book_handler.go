package handler

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mrfachri19/digital-library-backend/internal/entity"
	"github.com/mrfachri19/digital-library-backend/internal/usecase"
)

type BookHandler struct {
	Usecase usecase.BookUsecase
}

func NewBookHandler(u usecase.BookUsecase) *BookHandler {
	return &BookHandler{
		Usecase: u,
	}
}

// GET /books
func (h *BookHandler) GetAll(c *fiber.Ctx) error {
	books, err := h.Usecase.GetAll(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(books)
}

// GET /books/:id
func (h *BookHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	book, err := h.Usecase.GetByID(context.Background(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(book)
}

// POST /books
func (h *BookHandler) Create(c *fiber.Ctx) error {
	var book entity.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := h.Usecase.Create(context.Background(), book); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"message": "Book created"})
}

// PUT /books/:id
func (h *BookHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var book entity.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	book.ID = id

	if err := h.Usecase.Update(context.Background(), book); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Book updated"})
}

// DELETE /books/:id
func (h *BookHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.Usecase.Delete(context.Background(), id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Book deleted"})
}
