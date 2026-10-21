package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/alfannurfaiz9/shortlink_backend/internal/dto"
	custom_error "github.com/alfannurfaiz9/shortlink_backend/internal/error"
	"github.com/alfannurfaiz9/shortlink_backend/internal/model"
	"github.com/alfannurfaiz9/shortlink_backend/internal/repo"
	"github.com/alfannurfaiz9/shortlink_backend/pkg"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type AuthService struct {
	ar  *repo.AuthRepo
	rdb *redis.Client
}

func NewAuthService(ar *repo.AuthRepo, rdb *redis.Client) *AuthService {
	return &AuthService{
		ar:  ar,
		rdb: rdb,
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

func (a *AuthService) Login(ctx context.Context, body dto.Login) (dto.LoginResponse, error) {
	if len(body.Email) == 0 || len(body.Password) == 0 {
		return dto.LoginResponse{}, custom_error.EmptyLoginField
	}

	user, err := a.ar.FindUser(ctx, body.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return dto.LoginResponse{}, custom_error.InvalidEmailOrPassword
		}
		return dto.LoginResponse{}, err
	}

	if err := pkg.CompareHash(body.Password, user.Password_hash); err != nil {
		return dto.LoginResponse{}, custom_error.InvalidEmailOrPassword
	}

	claims := pkg.NewJWTClaims(user.Id)
	token, err := claims.GenToken()

	return dto.LoginResponse{Token: token}, err
}

func (a *AuthService) Logout(ctx context.Context, token string) error {
	key := fmt.Sprintf("shortlink:token:%s", token)

	err := a.rdb.Set(ctx, key, token, 5*time.Minute).Err()
	if err != nil {
		log.Printf("Failed to set key %s", err)
		return err
	}

	return nil
}
