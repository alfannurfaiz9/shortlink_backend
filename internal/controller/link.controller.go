package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"

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

// GetUserLink
//
// @Summary			Get user link
// @Description		Get user link
// @Tags			links
// @Produce			json
// @Security		BearerToken
// @Router			/links			[get]
// @Success			200				{object}	dto.Response
// @Failure			401				{object}	dto.ErrorResponse
// @Failure			500				{object}	dto.ErrorResponse
func (l *LinkController) GetUserLink(ctx *gin.Context) {
	id, exist := ctx.Get("id")
	if !exist {
		log.Println("Error get token")
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	result, err := l.ls.GetUserLink(ctx.Request.Context(), id.(int))
	if err != nil {
		log.Println("Error get token")
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Result:  result,
		Message: "successfully get links",
	})
}

// DeleteUserLink
//
// @Summary			Delete user link
// @Description		Delete user link
// @Tags			links
// @Produce			json
// @Security		BearerToken
// @Router			/links/{id}		[delete]
// @Param        	id   			path     	int  				true  "Link ID"
// @Success			200				{object}	dto.Response
// @Failure			401				{object}	dto.ErrorResponse
// @Failure			404				{object}	dto.ErrorResponse
// @Failure			500				{object}	dto.ErrorResponse
func (l *LinkController) DeleteUserLink(ctx *gin.Context) {
	paramId := ctx.Param("id")
	intParamId, err := strconv.Atoi(paramId)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	userId, exist := ctx.Get("id")
	if !exist {
		log.Println("Error get token")
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	if err := l.ls.DeleteUserLink(ctx.Request.Context(), intParamId, userId.(int)); err != nil {
		log.Println(err.Error())
		if errors.Is(err, custom_error.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, dto.ErrorResponse{
				Success: false,
				Message: "link not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "successfully delete links",
	})
}
