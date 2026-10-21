package service

import (
	"context"
	"log"

	"github.com/alfannurfaiz9/shortlink_backend/internal/dto"
	custom_error "github.com/alfannurfaiz9/shortlink_backend/internal/error"
	"github.com/alfannurfaiz9/shortlink_backend/internal/model"
	"github.com/alfannurfaiz9/shortlink_backend/internal/repo"
	"github.com/alfannurfaiz9/shortlink_backend/internal/utils"
)

type LinkService struct {
	lr *repo.LinkRepo
}

func NewLinkService(lr *repo.LinkRepo) *LinkService {
	return &LinkService{
		lr: lr,
	}
}

func (l *LinkService) CreateLink(ctx context.Context, id int, body dto.Link) error {
	if len(body.OriginalUrl) == 0 {
		return custom_error.EmptyField
	}

	if len(body.Slug) != 0 && len(body.Slug) < 3 {
		return custom_error.MinimumSlug
	}

	if len(body.Slug) >= 50 {
		return custom_error.MaximumSlug
	}

	if body.Slug == "api" || body.Slug == "login" || body.Slug == "register" || body.Slug == "dashboard" {
		return custom_error.ReservedKeyword
	}

	slug := body.Slug
	if len(body.Slug) == 0 {
		slug = utils.GenSlug()
	}

	for {
		_, err := l.lr.GetSlug(ctx, slug)
		if err != nil {
			break
		}

		if err == nil {
			log.Println("slug is exist, generate new slug")
			slug = utils.GenSlug()
		}
	}

	return l.lr.CreateLink(ctx, model.Link{
		UserId:      id,
		OriginalUrl: body.OriginalUrl,
		Slug:        slug,
	})
}

func (l *LinkService) GetUserLink(ctx context.Context, id int) ([]dto.UserLink, error) {
	links, err := l.lr.GetUserLink(ctx, id)
	if err != nil {
		return nil, err
	}

	result := make([]dto.UserLink, 0, len(links))
	for _, v := range links {
		result = append(result, dto.UserLink{
			Id:          v.Id,
			OriginalUrl: v.OriginalUrl,
			Slug:        v.Slug,
		})
	}

	return result, nil
}

func (l *LinkService) DeleteUserLink(ctx context.Context, id, userId int) error {
	return l.lr.DeleteUserLink(ctx, id, userId)
}
