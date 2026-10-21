package repo

import (
	"context"
	"errors"

	custom_error "github.com/alfannurfaiz9/shortlink_backend/internal/error"
	"github.com/alfannurfaiz9/shortlink_backend/internal/model"
	"github.com/jackc/pgx/v5"
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

func (l *LinkRepo) GetUserLink(ctx context.Context, id int) ([]model.Link, error) {
	sql := `
	SELECT id, slug, original_url FROM links
	WHERE user_id = $1 AND deleted_at IS NULL`
	args := []any{id}

	rows, err := l.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	var links []model.Link
	for rows.Next() {
		var link model.Link
		if err := rows.Scan(
			&link.Id,
			&link.Slug,
			&link.OriginalUrl,
		); err != nil {
			return nil, err
		}
		links = append(links, link)
	}

	return links, nil
}

func (l *LinkRepo) DeleteUserLink(ctx context.Context, id, userId int) error {
	sql := `
	UPDATE links 
	SET deleted_at = NOW()
	WHERE id = $1 AND user_id = $2 AND deleted_at IS NULL`
	args := []any{id, userId}

	cmd, err := l.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return custom_error.ErrNoRows
	}

	return nil
}

func (l *LinkRepo) GetOriginalLink(ctx context.Context, slug string) (string, error) {
	sql := `
	SELECT original_url FROM links
	WHERE slug = $1 AND deleted_at IS NULL`
	args := []any{slug}

	var originalLink string
	if err := l.db.QueryRow(ctx, sql, args...).Scan(&originalLink); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", custom_error.ErrNoRows
		}
		return "", err
	}

	return originalLink, nil
}
