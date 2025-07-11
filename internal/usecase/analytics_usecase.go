package usecase

import (
	"context"

	"github.com/mrfachri19/digital-library-backend/internal/repository"
)

type AnalyticsUsecase interface {
	GetTotalLendings(ctx context.Context) (int, error)
	GetTopBooks(ctx context.Context, limit int) ([]repository.AnalyticsResult, error)
	GetUniqueBorrowers(ctx context.Context) (int, error)

	// Tambahan untuk dashboard
	GetMostBorrowedBooks() ([]repository.MostBorrowedBook, error)
	GetMonthlyLendingTrends() ([]repository.MonthlyLending, error)
	GetBooksByCategory() ([]repository.CategoryCount, error)
}

type analyticsUsecase struct {
	repo repository.AnalyticsRepository
}

func NewAnalyticsUsecase(r repository.AnalyticsRepository) AnalyticsUsecase {
	return &analyticsUsecase{repo: r}
}

// ========== Analytics untuk dashboard ==========
func (u *analyticsUsecase) GetMostBorrowedBooks() ([]repository.MostBorrowedBook, error) {
	return u.repo.GetMostBorrowedBooks()
}

func (u *analyticsUsecase) GetMonthlyLendingTrends() ([]repository.MonthlyLending, error) {
	return u.repo.GetMonthlyLendingTrends()
}

func (u *analyticsUsecase) GetBooksByCategory() ([]repository.CategoryCount, error) {
	return u.repo.GetBooksByCategory()
}

// ========== Analytics dasar ==========
func (u *analyticsUsecase) GetTotalLendings(ctx context.Context) (int, error) {
	return u.repo.TotalLendings(ctx)
}

func (u *analyticsUsecase) GetTopBooks(ctx context.Context, limit int) ([]repository.AnalyticsResult, error) {
	return u.repo.TopBooks(ctx, limit)
}

func (u *analyticsUsecase) GetUniqueBorrowers(ctx context.Context) (int, error) {
	return u.repo.UniqueBorrowers(ctx)
}
