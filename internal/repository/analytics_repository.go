package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/mrfachri19/digital-library-backend/internal/config"
)

// Reusable struct for charts
type AnalyticsResult struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

// Dashboard-specific structs
type MostBorrowedBook struct {
	Title string `json:"title"`
	Count int    `json:"count"`
}

type MonthlyLending struct {
	Month string `json:"month"`
	Count int    `json:"count"`
}

type CategoryCount struct {
	Category string `json:"category"`
	Count    int    `json:"count"`
}

// Interface
type AnalyticsRepository interface {
	TotalLendings(ctx context.Context) (int, error)
	TopBooks(ctx context.Context, limit int) ([]AnalyticsResult, error)
	UniqueBorrowers(ctx context.Context) (int, error)

	GetMostBorrowedBooks() ([]MostBorrowedBook, error)
	GetMonthlyLendingTrends() ([]MonthlyLending, error)
	GetBooksByCategory() ([]CategoryCount, error)
}

// Implementation
type analyticsRepository struct {
	db *pgx.Conn
}

func NewAnalyticsRepository() AnalyticsRepository {
	conn, err := config.DB.Acquire(context.Background())
	if err != nil {
		panic("failed to acquire db connection: " + err.Error())
	}
	return &analyticsRepository{db: conn.Conn()}
}

// ========== BASIC ANALYTICS ==========

func (r *analyticsRepository) TotalLendings(ctx context.Context) (int, error) {
	var count int
	err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM lendings`).Scan(&count)
	return count, err
}

func (r *analyticsRepository) TopBooks(ctx context.Context, limit int) ([]AnalyticsResult, error) {
	rows, err := r.db.Query(ctx, `
		SELECT b.title, COUNT(*) as count
		FROM lendings l
		JOIN books b ON b.id = l.book_id
		GROUP BY b.title
		ORDER BY count DESC
		LIMIT $1
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []AnalyticsResult
	for rows.Next() {
		var r AnalyticsResult
		if err := rows.Scan(&r.Label, &r.Value); err != nil {
			return nil, err
		}
		result = append(result, r)
	}
	return result, nil
}

func (r *analyticsRepository) UniqueBorrowers(ctx context.Context) (int, error) {
	var count int
	err := r.db.QueryRow(ctx, `SELECT COUNT(DISTINCT borrower) FROM lendings`).Scan(&count)
	return count, err
}

// ========== DASHBOARD ANALYTICS ==========

func (r *analyticsRepository) GetMostBorrowedBooks() ([]MostBorrowedBook, error) {
	rows, err := r.db.Query(context.Background(), `
		SELECT b.title, COUNT(*) as count
		FROM lendings l
		JOIN books b ON l.book_id = b.id
		GROUP BY b.title
		ORDER BY count DESC
		LIMIT 5
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []MostBorrowedBook
	for rows.Next() {
		var b MostBorrowedBook
		if err := rows.Scan(&b.Title, &b.Count); err != nil {
			return nil, err
		}
		result = append(result, b)
	}
	return result, nil
}

func (r *analyticsRepository) GetMonthlyLendingTrends() ([]MonthlyLending, error) {
	rows, err := r.db.Query(context.Background(), `
		SELECT TO_CHAR(borrow_date, 'Mon') AS month, COUNT(*) 
		FROM lendings
		GROUP BY month
		ORDER BY MIN(borrow_date)
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []MonthlyLending
	for rows.Next() {
		var m MonthlyLending
		if err := rows.Scan(&m.Month, &m.Count); err != nil {
			return nil, err
		}
		result = append(result, m)
	}
	return result, nil
}

func (r *analyticsRepository) GetBooksByCategory() ([]CategoryCount, error) {
	rows, err := r.db.Query(context.Background(), `
		SELECT category, COUNT(*) 
		FROM books
		GROUP BY category
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []CategoryCount
	for rows.Next() {
		var c CategoryCount
		if err := rows.Scan(&c.Category, &c.Count); err != nil {
			return nil, err
		}
		result = append(result, c)
	}
	return result, nil
}
