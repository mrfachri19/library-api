package usecase

import (
	"context"
	"time"

	"github.com/mrfachri19/digital-library-backend/internal/entity"
	"github.com/mrfachri19/digital-library-backend/internal/repository"
)

type LendingUsecase interface {
	Create(ctx context.Context, lending entity.Lending) error
	GetAll(ctx context.Context) ([]entity.Lending, error)
	MarkAsReturned(ctx context.Context, id int) error
}

type lendingUsecase struct {
	repo repository.LendingRepository
}

func NewLendingUsecase(r repository.LendingRepository) LendingUsecase {
	return &lendingUsecase{
		repo: r,
	}
}

func (u *lendingUsecase) Create(ctx context.Context, lending entity.Lending) error {
	lending.BorrowDate = time.Now() // Auto set waktu pinjam
	return u.repo.Create(ctx, lending)
}

func (u *lendingUsecase) GetAll(ctx context.Context) ([]entity.Lending, error) {
	return u.repo.GetAll(ctx)
}

func (u *lendingUsecase) MarkAsReturned(ctx context.Context, id int) error {
	return u.repo.MarkAsReturned(ctx, id, time.Now())
}
