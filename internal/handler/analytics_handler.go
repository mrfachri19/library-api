package handler

import (
	"context"

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
	total, _ := h.Usecase.GetTotalLendings(context.Background())
	top, _ := h.Usecase.GetTopBooks(context.Background(), 5)
	users, _ := h.Usecase.GetUniqueBorrowers(context.Background())

	return c.JSON(fiber.Map{
		"total_lendings":   total,
		"top_books":        top,
		"unique_borrowers": users,
	})
}
