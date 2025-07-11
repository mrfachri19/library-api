package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrfachri19/digital-library-backend/internal/usecase"
)

type AnalyticsHandler struct {
	Usecase usecase.AnalyticsUsecase
}

func NewAnalyticsHandler(u usecase.AnalyticsUsecase) *AnalyticsHandler {
	return &AnalyticsHandler{Usecase: u}
}

func (h *AnalyticsHandler) GetSummary(c *fiber.Ctx) error {
	mostBorrowed, err := h.Usecase.GetMostBorrowedBooks()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	monthly, err := h.Usecase.GetMonthlyLendingTrends()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	category, err := h.Usecase.GetBooksByCategory()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"most_borrowed":         mostBorrowed,
		"monthly_lending":       monthly,
		"category_distribution": category,
	})
}
