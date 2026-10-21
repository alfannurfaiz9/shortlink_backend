package repo

import (
	"context"

	custom_error "github.com/alfannurfaiz9/shortlink_backend/internal/error"
	"github.com/alfannurfaiz9/shortlink_backend/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LinkRepo struct {
	db *pgxpool.Pool
}

func NewLinkRepo(db *pgxpool.Pool) *LinkRepo {
	return &LinkRepo{
		db: db,
	}
}

func (l *LinkRepo) GetSlug(ctx context.Context, slug string) (string, error) {
	sql := `
	SELECT slug FROM links
	WHERE slug = $1`
	args := []any{slug}

	var exist string
	if err := l.db.QueryRow(ctx, sql, args...).Scan(&exist); err != nil {
		return "", custom_error.ErrNoRows
	}

	return exist, nil
}

func (l *LinkRepo) CreateLink(ctx context.Context, body model.Link) error {
	sql := `
	INSERT INTO links(user_id, original_url, slug)
	VALUES($1, $2, $3)`
	args := []any{body.UserId, body.OriginalUrl, body.Slug}

	cmd, err := l.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return custom_error.ErrNoRows
	}

	return nil
}
