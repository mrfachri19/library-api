package usecase

import (
	"context"

	"github.com/mrfachri19/digital-library-backend/internal/repository"
)

type AnalyticsUsecase interface {
	GetTotalLendings(ctx context.Context) (int, error)
	GetTopBooks(ctx context.Context, limit int) ([]repository.AnalyticsResult, error)
	GetUniqueBorrowers(ctx context.Context) (int, error)
}

type analyticsUsecase struct {
	repo repository.AnalyticsRepository
}

func NewAnalyticsUsecase(r repository.AnalyticsRepository) AnalyticsUsecase {
	return &analyticsUsecase{repo: r}
}

func (u *analyticsUsecase) GetTotalLendings(ctx context.Context) (int, error) {
	return u.repo.TotalLendings(ctx)
}

func (u *analyticsUsecase) GetTopBooks(ctx context.Context, limit int) ([]repository.AnalyticsResult, error) {
	return u.repo.TopBooks(ctx, limit)
}

func (u *analyticsUsecase) GetUniqueBorrowers(ctx context.Context) (int, error) {
	return u.repo.UniqueBorrowers(ctx)
}
