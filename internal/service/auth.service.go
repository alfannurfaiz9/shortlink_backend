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

func (a *AuthService) Login(ctx context.Context, body dto.Login) (dto.UserInfo, error) {
	if len(body.Email) == 0 || len(body.Password) == 0 {
		return dto.UserInfo{}, custom_error.EmptyLoginField
	}

	user, err := a.ar.FindUser(ctx, body.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return dto.UserInfo{}, custom_error.InvalidEmailOrPassword
		}
		return dto.UserInfo{}, err
	}

	if err := pkg.CompareHash(body.Password, user.Password_hash); err != nil {
		return dto.UserInfo{}, custom_error.InvalidEmailOrPassword
	}

	claims := pkg.NewJWTClaims(user.Id)
	token, err := claims.GenToken()

	result := dto.UserInfo{
		Email:    user.Email,
		FullName: user.Fullname,
		Token:    token,
	}

	return result, err
}
