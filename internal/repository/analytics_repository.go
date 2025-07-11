package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/mrfachri19/digital-library-backend/internal/config"
)

type AnalyticsResult struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

type AnalyticsRepository interface {
	TotalLendings(ctx context.Context) (int, error)
	TopBooks(ctx context.Context, limit int) ([]AnalyticsResult, error)
	UniqueBorrowers(ctx context.Context) (int, error)
}

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
