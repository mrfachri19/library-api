package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/mrfachri19/digital-library-backend/internal/config"
	"github.com/mrfachri19/digital-library-backend/internal/entity"
)

type LendingRepository interface {
	Create(ctx context.Context, lending entity.Lending) error
	GetAll(ctx context.Context) ([]entity.Lending, error)
	MarkAsReturned(ctx context.Context, id int, returnDate time.Time) error
}

type lendingRepository struct {
	db *pgx.Conn
}

func NewLendingRepository() LendingRepository {
	conn, err := config.DB.Acquire(context.Background())
	if err != nil {
		panic("failed to acquire db connection: " + err.Error())
	}
	return &lendingRepository{db: conn.Conn()}
}

func (r *lendingRepository) Create(ctx context.Context, lending entity.Lending) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO lendings (book_id, borrower, borrow_date)
		VALUES ($1, $2, $3)
	`, lending.BookID, lending.Borrower, lending.BorrowDate)
	return err
}

func (r *lendingRepository) GetAll(ctx context.Context) ([]entity.Lending, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, book_id, borrower, borrow_date, return_date
		FROM lendings ORDER BY borrow_date DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []entity.Lending
	for rows.Next() {
		var lending entity.Lending
		err := rows.Scan(&lending.ID, &lending.BookID, &lending.Borrower, &lending.BorrowDate, &lending.ReturnDate)
		if err != nil {
			return nil, err
		}
		result = append(result, lending)
	}
	return result, nil
}

func (r *lendingRepository) MarkAsReturned(ctx context.Context, id int, returnDate time.Time) error {
	_, err := r.db.Exec(ctx, `
		UPDATE lendings SET return_date = $1 WHERE id = $2
	`, returnDate, id)
	return err
}
