package controller

import (
	"errors"
	"log"
	"net/http"

	"github.com/alfannurfaiz9/shortlink_backend/internal/dto"
	custom_error "github.com/alfannurfaiz9/shortlink_backend/internal/error"
	"github.com/alfannurfaiz9/shortlink_backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type LinkController struct {
	ls *service.LinkService
}

func NewLinkController(ls *service.LinkService) *LinkController {
	return &LinkController{
		ls: ls,
	}
}

// CreateLink
//
// @Summary			Create link
// @Description		Create link
// @Tags			links
// @Produce			json
// @Security		BearerToken
// @Router			/links			[post]
// @Param			data			body		dto.Link			true	"body"
// @Success			201				{object}	dto.Response
// @Failure			401				{object}	dto.ErrorResponse
// @Failure			500				{object}	dto.ErrorResponse
func (l *LinkController) CreateLink(ctx *gin.Context) {
	var body dto.Link

	if err := ctx.ShouldBindWith(&body, binding.JSON); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	id, exist := ctx.Get("id")
	if !exist {
		log.Println("Error get token")
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	if err := l.ls.CreateLink(ctx.Request.Context(), id.(int), body); err != nil {
		log.Println(err.Error())
		if errors.Is(err, custom_error.EmptyField) {
			ctx.JSON(http.StatusNotAcceptable, dto.ErrorResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		if errors.Is(err, custom_error.ReservedKeyword) {
			ctx.JSON(http.StatusNotAcceptable, dto.ErrorResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		if errors.Is(err, custom_error.MinimumSlug) {
			ctx.JSON(http.StatusNotAcceptable, dto.ErrorResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		if errors.Is(err, custom_error.MaximumSlug) {
			ctx.JSON(http.StatusNotAcceptable, dto.ErrorResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "link created successfully",
	})
}
