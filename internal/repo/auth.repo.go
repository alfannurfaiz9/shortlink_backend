package repo

import (
	"context"
	"errors"

	"github.com/alfannurfaiz9/shortlink_backend/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepo struct {
	db *pgxpool.Pool
}

func NewAuthRepo(db *pgxpool.Pool) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (a *AuthRepo) FindUserByEmail(ctx context.Context, email string) (model.User, error) {
	sql := "SELECT email FROM users WHERE email = $1"
	args := []any{email}

	var user model.User
	if err := a.db.QueryRow(ctx, sql, args...).Scan(
		&user.Email); err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (a *AuthRepo) Register(ctx context.Context, body model.User) error {
	sql := `
	INSERT INTO users(email, full_name, password_hash)
	VALUES($1, $2, $3)`
	args := []any{body.Email, body.Fullname, body.Password_hash}

	cmd, err := a.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return errors.New("No rows affected")
	}

	return nil
}
