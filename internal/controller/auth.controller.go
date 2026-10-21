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

type AuthController struct {
	as *service.AuthService
}

func NewController(as *service.AuthService) *AuthController {
	return &AuthController{
		as: as,
	}
}

// Register
//
// @Summary			Create new user
// @Description		Create new user
// @Tags			auth
// @Produce			json
// @Router			/auth/register		[post]
// @Param			data				body		dto.Register	true	"body to register"
// @Success			201					{object}	dto.Response
// @Failure			406					{object}	dto.ErrorResponse
// @Failure			500					{object}	dto.ErrorResponse
func (a *AuthController) Register(ctx *gin.Context) {
	var body dto.Register
	if err := ctx.ShouldBindWith(&body, binding.JSON); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "internal server error",
		})

		return
	}

	if err := a.as.Register(ctx, body); err != nil {
		log.Println(err.Error())
		if errors.Is(err, custom_error.UserAlreadyExist) {
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
		Message: "account created successfully",
	})
}
