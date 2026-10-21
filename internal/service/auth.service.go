package service

import (
	"context"
	"errors"

	"github.com/alfannurfaiz9/shortlink_backend/internal/dto"
	custom_error "github.com/alfannurfaiz9/shortlink_backend/internal/error"
	"github.com/alfannurfaiz9/shortlink_backend/internal/model"
	"github.com/alfannurfaiz9/shortlink_backend/internal/repo"
	"github.com/alfannurfaiz9/shortlink_backend/pkg"
	"github.com/jackc/pgx/v5"
)

type AuthService struct {
	ar *repo.AuthRepo
}

func NewAuthService(ar *repo.AuthRepo) *AuthService {
	return &AuthService{
		ar: ar,
	}
}

func (a *AuthService) Register(ctx context.Context, body dto.Register) error {
	_, err := a.ar.FindUserByEmail(ctx, body.Email)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}
	if err == nil {
		return custom_error.UserAlreadyExist
	}

	hash := pkg.NewHashConfig()
	hashedPass := hash.GenerateHash(body.Password)

	if err := a.ar.Register(ctx, model.User{
		Email:         body.Email,
		Fullname:      body.Fullname,
		Password_hash: hashedPass,
	}); err != nil {
		return err
	}

	return nil
}
