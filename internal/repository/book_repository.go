package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/mrfachri19/digital-library-backend/internal/config"
	"github.com/mrfachri19/digital-library-backend/internal/entity"
)

type BookRepository interface {
	GetAll(ctx context.Context) ([]entity.Book, error)
	GetByID(ctx context.Context, id int) (*entity.Book, error)
	Create(ctx context.Context, book entity.Book) error
	Update(ctx context.Context, book entity.Book) error
	Delete(ctx context.Context, id int) error
}

type bookRepository struct {
	db *pgx.Conn
}

func NewBookRepository() BookRepository {
	conn, err := config.DB.Acquire(context.Background())
	if err != nil {
		panic("failed to acquire db connection: " + err.Error())
	}
	return &bookRepository{db: conn.Conn()}
}

// GetAll books
func (r *bookRepository) GetAll(ctx context.Context) ([]entity.Book, error) {
	rows, err := r.db.Query(ctx, "SELECT id, title, author, isbn, quantity, category FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []entity.Book
	for rows.Next() {
		var book entity.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Quantity, &book.Category); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (r *bookRepository) GetByID(ctx context.Context, id int) (*entity.Book, error) {
	row := r.db.QueryRow(ctx, "SELECT id, title, author, isbn, quantity, category FROM books WHERE id=$1", id)

	var book entity.Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Quantity, &book.Category)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *bookRepository) Create(ctx context.Context, book entity.Book) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO books (title, author, isbn, quantity, category)
		VALUES ($1, $2, $3, $4, $5)
	`, book.Title, book.Author, book.ISBN, book.Quantity, book.Category)

	return err
}

func (r *bookRepository) Update(ctx context.Context, book entity.Book) error {
	_, err := r.db.Exec(ctx, `
		UPDATE books SET title=$1, author=$2, isbn=$3, quantity=$4, category=$5
		WHERE id=$6
	`, book.Title, book.Author, book.ISBN, book.Quantity, book.Category, book.ID)

	return err
}

func (r *bookRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, "DELETE FROM books WHERE id=$1", id)
	return err
}
