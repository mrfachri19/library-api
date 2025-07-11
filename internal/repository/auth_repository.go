package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/mrfachri19/digital-library-backend/internal/config"
	"github.com/mrfachri19/digital-library-backend/internal/entity"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, user entity.User) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

type authRepository struct {
	db *pgx.Conn
}

func NewAuthRepository() AuthRepository {
	conn, _ := config.DB.Acquire(context.Background())
	return &authRepository{db: conn.Conn()}
}

func (r *authRepository) CreateUser(ctx context.Context, user entity.User) error {
	_, err := r.db.Exec(ctx, "INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password)
	return err
}

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	row := r.db.QueryRow(ctx, "SELECT id, email, password FROM users WHERE email = $1", email)
	var user entity.User
	if err := row.Scan(&user.ID, &user.Email, &user.Password); err != nil {
		return nil, err
	}
	return &user, nil
}
