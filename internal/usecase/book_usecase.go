package usecase

import (
	"context"

	"github.com/mrfachri19/digital-library-backend/internal/entity"
	"github.com/mrfachri19/digital-library-backend/internal/repository"
)

type BookUsecase interface {
	GetAll(ctx context.Context) ([]entity.Book, error)
	GetByID(ctx context.Context, id int) (*entity.Book, error)
	Create(ctx context.Context, book entity.Book) error
	Update(ctx context.Context, book entity.Book) error
	Delete(ctx context.Context, id int) error
}

type bookUsecase struct {
	repo repository.BookRepository
}

func NewBookUsecase(r repository.BookRepository) BookUsecase {
	return &bookUsecase{
		repo: r,
	}
}

func (u *bookUsecase) GetAll(ctx context.Context) ([]entity.Book, error) {
	return u.repo.GetAll(ctx)
}

func (u *bookUsecase) GetByID(ctx context.Context, id int) (*entity.Book, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *bookUsecase) Create(ctx context.Context, book entity.Book) error {
	// Validasi sederhana bisa ditambah di sini
	return u.repo.Create(ctx, book)
}

func (u *bookUsecase) Update(ctx context.Context, book entity.Book) error {
	return u.repo.Update(ctx, book)
}

func (u *bookUsecase) Delete(ctx context.Context, id int) error {
	return u.repo.Delete(ctx, id)
}
